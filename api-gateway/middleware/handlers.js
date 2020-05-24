
/**
 * Module dependencies.
 */

var { logger, getLogBody } = require("./logger")
var Ajv = require("ajv")
var HttpStatus = require("http-status-codes")

/**
 * Error handler:
 *
 * Development error handler, providing stack traces
 * and error message responses for requests accepting text, html,
 * or json.
 * @return {Function}
 * @api public
 */
const errorHandler = () => function (err, req, res, next) {
	// Set local err
	res.locals.error = err
	res.locals.logLevel = "error"

	// Set status code
	res.statusCode = getStatusCode(err)

	// cannot actually respond
	if (res._header)
		return req.socket.destroy()

	var errorBody = {
	    message: HttpStatus.getStatusText(res.statusCode),
		details: err.message || "unknown"
	}

	var jsonErrorBody = { }
	if ((process.env.OTIS_ENV || "development") === "development") {
	    errorBody.error = err
		jsonErrorBody = JSON.stringify(errorBody, null, "\t")
	} else {
	    jsonErrorBody = JSON.stringify(errorBody)
	}


	if (req.accepts("application/json") || req.accepts("json"))
		res.setHeader("Content-Type", "application/json; charset=utf-8")
	else
		res.setHeader("Content-Type", "text/plain; charset=utf-8")

	res.end(jsonErrorBody)
	next()
}

/* Returns the equivalent HTTP status code of a JS error */
function getStatusCode (err) {
	if (err.status && err.status !== 200)
		return err.status
	if (err.statusCode && err.statusCode !== 200)
		return err.statusCode
	if (err instanceof Ajv.ValidationError)
		return 400

	return 500
}

const logHandler = () => function (req, res, time) {
	const logBody = getLogBody(time, req, res, res.locals.error)
	logger.log(logBody)
}

module.exports = {
	logHandler,
	errorHandler
}
