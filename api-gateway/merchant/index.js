const express = require("express")
const api = require("./api")
const router = express.Router()
const utils = require("../utils")
const jwtAuthz = require("express-jwt-authz")
const { logger } = require("../middleware/logger")

var has_override = function(key) {
    var cfg_overrides = utils.getValue("apis::merchant::overrides") || {}
    return (Object.hasOwnProperty.call(cfg_overrides, key) && cfg_overrides[key] === false)
}

/* Add API Auth */
var jwtAuthFuncs = api.auth
for (var fn of jwtAuthFuncs) {
	const message = `Adding auth func to Merchant API: ${fn.name}`
	logger.info(message)
	router.use(fn)
}

/* Add API methods */
for (var endpoint in api.routes) {
    var jwtAuthScopes = has_override("auth") ? [] : route.scopes || []
    var validators = has_override("validators") ? [] : route.validators || []
	var route = api.routes[endpoint]
	var prefix = route.prefix
	var handler = utils.wrapFuncInMiddleware(route.handler, validators)
	var verb = route.verb
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
