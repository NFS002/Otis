const express = require("express")
const api = require("./api")
const router = express.Router()
const utils = require("../utils")

/* Add API methods */
for (var endpoint in api) {
	var prefix = api[endpoint].prefix
	var schemas = api[endpoint].schemas
	var handler = utils.wrapFuncInMiddleware(api[endpoint].handler, schemas)
	var verb = api[endpoint].verb
	router[verb](prefix, handler)
}

/* Return router */
module.exports = router
