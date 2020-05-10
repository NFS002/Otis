const express = require('express')
const api = require('./api')
const router = express.Router()
const utils = require('../utils')

/* Add API methods */
for (endpoint in api) {
    prefix = api[endpoint].prefix
    handler = utils.wrapFuncInMiddleware(api[endpoint].handler)
    verb = api[endpoint].verb
    router[verb](prefix, handler)
}

/* Return router */
module.exports = router
