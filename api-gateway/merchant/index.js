const express = require("express")
const api = require("./api")
const router = express.Router()
const utils = require("../utils")
const jwtAuthz = require("express-jwt-authz")
const { logger } = require("../middleware/logger")

/* Add API Auth */
var jwtAuthFuncs = api.auth
for (var fn of jwtAuthFuncs) {
	const message = `Adding auth func to Merchant API: ${fn.name}`
	logger.info(message)
	router.use(fn)
}

/* Add API methods */
for (var endpoint in api.routes) {
	var route = api.routes[endpoint]
	var prefix = route.prefix
	var validators = route.validators
	var handler = utils.wrapFuncInMiddleware(route.handler, validators)
	var verb = route.verb
	var jwtAuthScopes = route.scopes || []
	if (jwtAuthScopes.length) {
		var message = `Adding route to Merchant API: ${verb} ${prefix}. Using handler: '${route.handler.name}', with scopes '${jwtAuthScopes}'`
		logger.info(message)
		jwtAuthScopes = jwtAuthz(jwtAuthScopes, {
			failWithError: true,
			checkAllScopes: true
		})
	}
	router[verb](prefix, jwtAuthScopes, handler)
}

/* Return router */
module.exports = router
