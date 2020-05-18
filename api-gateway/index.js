const express = require("express")
const bodyParser = require("body-parser")
const { getValue } = require("./utils")
const helmet = require("helmet")
const { logHandler, errorHandler } = require("./middleware/handlers")
const { logger } = require("./middleware/logger")
const responseTime = require("response-time")
const morganBody = require("morgan-body")

const app = express()

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }))

// parse application/json
app.use(bodyParser.json())

// Security middleware
app.use(helmet())

morganBody(app)

/* Log response time and other request/response metrics */
app.use(responseTime(logHandler()))

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

app.use(errorHandler())

app.listen(port, address, () => {
	var msg = `[api-gateway:${process.env.OTIS_ENV}] Listening at http://${address}:${port}`
	var info = {
		level: "info",
		message: msg
	}
	console.log(msg)
	logger.log(info)
}
)
