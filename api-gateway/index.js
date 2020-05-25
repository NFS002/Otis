const express = require("express")
const bodyParser = require("body-parser")
const { getValue } = require("./utils")
const helmet = require("helmet")
const { logHandler, errorHandler, notFoundHandler, invalidContentTypeHandler } = require("./middleware/handlers")
const { logger } = require("./middleware/logger")
const morganBody = require("./middleware/morgan-body")
const requestId = require("./middleware/request-id")
const responseTime = require("./middleware/response-time")

const app = express()

// Security middleware
app.use(helmet())


var useMorgan = getValue("morgan")
if (useMorgan === true) {
	/* Each request is assigned a unique ID */
	app.use(requestId())
	morganBody(app)
}

/* Log response time and other request/response metrics */
app.use(responseTime(logHandler()))

/*
 * Try parse application/json,
 * and then parse all other request content types as application/x-www-form-urlencoded
 */
app.use(invalidContentTypeHandler())
app.use(bodyParser.json({ extended: true, type: "*/*" }))

const apis = getValue("apis")
for (var api in apis) {
	const a = apis[api]
	const module = require(a.path)
	const name = a.name
	const message = `[api-gateway:${process.env.OTIS_ENV}] Adding API: ${name}`
	const prefix = a.prefix
	console.log(message)
	logger.info(message)
	app.use(prefix, module)
}

var port = getValue("port")
var address = getValue("address")

app.use(notFoundHandler())

app.use(errorHandler())

app.listen(port, address, () => {
	var msg = `[api-gateway:${process.env.OTIS_ENV}] Listening at http://${address}:${port}`
	var info = {
		level: "info",
		message: msg
	}
	console.log(msg)
	logger.log(info)
})
