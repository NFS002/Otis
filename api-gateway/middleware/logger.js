/* Dependencies */
const os = require("os")
const winston = require("winston")
const { getValue } = require("../utils")
const { Rollbar } = require("./winston-rollbar")

const { format } = winston

const logConfig = getValue("logs")
const logFiles = logConfig.files

const winstonTransports = logFiles.map(log => new winston.transports.File(log))

/* Print to console */
if (logConfig.console === true)
	winstonTransports.push(new winston.transports.Console())


/* Log to rollbar  */
if (logConfig.rollbar === true) {
	const r = new Rollbar({
		rollbarConfig: {
			accessToken: process.env.OTIS_ROLLBAR_ACCESS_TOKEN,
			environment: process.env.OTIS_ENV,
			reportLevel: "error"
		}
	})
	winstonTransports.push(r)
}

const getWinstonFormat = function () {
	if (logConfig.pretty_print === true)
		return format.combine(
			format.timestamp(),
			format.json(),
			format.prettyPrint()
		)
	 else
		return format.combine(
			format.timestamp(),
			format.json()
		)
}

const winstonConf = {
	level: "info",
	format: getWinstonFormat(),
	meta: true,
	defaultMeta: { service: "api-gateway" },
	transports: winstonTransports
}

const getLogMessage = function (req, res) {
	return "[api-gateway:" + req.hostname + "] " + req.method + "/" + req.httpVersion +
    " " + req.originalUrl + " (" + res.statusCode + ")"
}

const getMetaLogBody = function (req, res) {
	const meta = {}
	meta.memory_usage = process.memoryUsage()
	meta.total_mem = os.totalmem()
	meta.pid = process.pid ? process.pid : "-"
	meta.gid = process.gid ? process.gid : "-"
	meta.argv = process.argv ? process.argv : "-"
	meta.platform = process.platform ? process.platform : "-"
	meta.total_unused_mem = os.freemem()
	return meta
}


const getLogBody = function (time, req, res, err) {
	const reqLogBody = {}
	const resLogBody = {}
	const logBody = {}

	reqLogBody.type = "http request"
	resLogBody.type = "http response"

	reqLogBody.headers = req.headers
	resLogBody.headers = res.getHeaders()

	reqLogBody.url = req.originalUrl

	reqLogBody.host = req.hostname
	reqLogBody.ip = req.ip

	const reqBody = req.body
	const resBody = res.locals.body

	if (reqBody) reqLogBody.body = reqBody
	if (resBody) resLogBody.body = resBody

	const reqCookies = req.cookies
	const resCookies = res.cookies

	if (reqCookies) reqLogBody.cookies = reqCookies
	if (resCookies) resLogBody.cookies = resCookies

	logBody.request = reqLogBody
	logBody.response = resLogBody
	logBody.message = getLogMessage(req, res)
	logBody.meta = getMetaLogBody(req, res)


	if (time) {
	    var roundedStrTime = String(time.toFixed(3))
		logBody.response_time = roundedStrTime
		if (time > 3500) {
		    logBody.message += (" - Warning: Slow Response Time: " + roundedStrTime)
			logBody.level = "warning"
		}
	}

	if (err) {
	    logBody.level = "error"
		logBody.error = err
		logBody.message += (" - Error: " + err.message)
	}

	/* Default log level */
	logBody.level = logBody.level || "info"

	return logBody
}

const logger = winston.createLogger(winstonConf)

module.exports = {
	logger,
	getLogBody
}
