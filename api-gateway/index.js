const express = require('express')
const bodyParser = require('body-parser')
const { getService, getValue } = require('./utils.js')
const helmet = require('helmet')
const errorhandler = require('./middleware/errorhandler')

const app = express()

// parse application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: false }))

// parse application/json
app.use(bodyParser.json())

// Security middleware
app.use(helmet())

/* Add API modules */
//app.get('/merchant/general/get', function(req, res) {
  //  throw new Error("Testing")
//})


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

app.use( errorhandler() )

app.listen(port, address, () => console.log(`API gateway listening at http://${address}:${port}`))