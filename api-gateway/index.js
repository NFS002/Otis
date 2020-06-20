const express = require("express")
const bodyParser = require("body-parser")
const { getValue } = require("./utils")
const helmet = require("helmet")
const { logger } = require("./middleware/logger")
const morganBody = require("./middleware/morgan-body")
const requestId = require("./middleware/request-id")
const responseTime = require("./middleware/response-time")
const { logHandler, errorHandler, setRawBody, notFoundHandler, invalidContentTypeHandler } = require("./middleware/handlers")

const app = express()

// Security middleware
app.use(helmet())

if (getValue("morgan") === true) {
	/* Each request is assigned a unique ID */
	app.use(requestId())
	morganBody(app, { theme: getValue("morgan_theme") })
}

/* Log response time and other request/response metrics */
app.use(responseTime(logHandler()))

/*
 * Try parse as application/json request content type,
 * and set the raw request body on the request object for debugging
 */
app.use(invalidContentTypeHandler())
app.use(bodyParser.json({ extended: true, type: "*/json", verify: setRawBody }))

/* Add global authentication middlewares */
const auth = getValue("global_auth") || []
for (var func of auth) {
	const message = `Adding global auth middleware: ${func.name}`
	logger.info(message)
	app.use(func)
}

/* Add API router modules */
const apis = getValue("apis")
for (var api in apis) {
	const a = apis[api]
	const module = require(a.path)
	const message = `[api-gateway:${process.env.OTIS_ENV}] Adding API: ${api}`
	const prefix = a.prefix
	logger.info(message)
	app.use(prefix, module)
}

var port = getValue("port")
var address = getValue("address")

app.use(notFoundHandler())

app.use(errorHandler())

app.listen(port, address, () => {
	var msg = `[api-gateway:${process.env.OTIS_ENV}] Listening at http://${address}:${port}`
	logger.info(msg)
})
