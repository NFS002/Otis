const express = require('express')
const bodyParser = require('body-parser')
const { getService, getValue } = require('./utils')
const helmet = require('helmet')
const { log_handler, error_handler } = require('./middleware/handlers')
const { logger } = require('./middleware/logger')
const responseTime = require('response-time')
const morganBody = require('morgan-body');

const app = express()

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }))

// parse application/json
app.use(bodyParser.json())

// Security middleware
app.use(helmet())

morganBody(app);

/* Log response time and other request/response metrics */
app.use( responseTime( log_handler() ) )

const apis = getValue('apis')
for (api in apis) {
  let a = apis[api]
  let module = require(a.path)
  let name = a.name
  let prefix = a.prefix
  app.use(prefix, module)
}

port = getValue('port')
address = getValue('address')

app.use( error_handler() )

app.listen(port, address, () => {
    var msg = `[api-gateway:${process.env['OTIS_ENV']}] Listening at http://${address}:${port}`
    var info = {
        level: 'info',
        message: msg
    }
    console.log(msg)
    logger.log(info)
  }
)