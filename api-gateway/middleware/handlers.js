
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
	var env = process.env.OTIS_ENV || "development"
	if (env === "development") {
	    errorBody.error = err
		jsonErrorBody = JSON.stringify(errorBody, null, "\t")
	} else {
	    errorBody.error = "true"
	    jsonErrorBody = JSON.stringify(errorBody)
	}


	if (req.accepts("application/json") || req.accepts("json"))
		res.setHeader("Content-Type", "application/json; charset=utf-8")
	else
		res.setHeader("Content-Type", "text/plain; charset=utf-8")

	res.locals.body = jsonErrorBody
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
	if (err instanceof PathNotFoundError)
		return 404
	if (err instanceof InvalidContentTypeError)
	    return 415

	return 500
}

const logHandler = () => function (req, res, time) {
	const logBody = getLogBody(time, req, res, res.locals.error)
	logger.log(logBody)
}

/* Set the req.rawBody property to the raw request body */
var setRawBody = function (req, res, buf, encoding) {
	if (buf && buf.length) {
		req.rawBody = buf.toString(encoding || "utf8")
		req.hasRawBody = true
	}
}

class PathNotFoundError extends Error {
	constructor (message, ...args) {
		super(message, ...args)
	}
}

const notFoundHandler = () => function (req, res, next) {
	throw new PathNotFoundError(`Request path (${req.path}) could not be resolved to a valid route`)
}

class InvalidContentTypeError extends Error {
	constructor (message, ...args) {
		super(message, ...args)
	}
}

const invalidContentTypeHandler = () => function (req, res, next) {
	var ctype = req.headers["content-type"]
	var validCType = "*/json"
	if (!req.is(validCType))
	    throw new InvalidContentTypeError(`Request content type (${ctype}) must be set to ${validCType}`)

	next()
}

module.exports = {
	logHandler,
	errorHandler,
	notFoundHandler,
	setRawBody,
	invalidContentTypeHandler
}
