const ClientPool = require('../connection-pool')
const { getProtoPackage, getService, getValue } = require('../utils.js')

const include_dirs = ['proto', '/Users/noah/Otis/backend/' ]
const merchant_service_proto_file = 'example.proto'
const merchant_package_name = 'example'

const merchant_service = getValue('services::merchant')
const merchant_proto_pkg = getProtoPackage(merchant_service.proto_file, merchant_service.package, merchant_service.proto_dirs )
const merchant_pool_opts = { grpcPkg: merchant_proto_pkg, serviceName: merchant_service.service, url: merchant_service.address }

const merchant_pool = new ClientPool( merchant_pool_opts )

async function transactions(req, res) {
   const { RPC_Bite } = merchant_pool
   const value = await RPC_Bite({ "id" : "12hj47" })
   res.end( JSON.stringify( value ) )
}

module.exports = {
    transactions
}