const express = require('express')
const api = require('./api')

const router = express.Router()

router.get('/transactions', api.transactions)

module.exports = router
