const ClientPool = require('../connection-pool')
const { getProtoPackage, getService, getValue } = require('../utils.js')

const merchant_service = getValue('services::merchant')
const merchant_proto_pkg = getProtoPackage(merchant_service.proto_file, merchant_service.package, merchant_service.proto_dirs )
const merchant_pool_opts = { grpcPkg: merchant_proto_pkg, serviceName: merchant_service.service, url: merchant_service.address }

const merchant_pool = new ClientPool( merchant_pool_opts )

/* General merchant handlers */

async function g_get(req, res) {
   var value;
   const { _GetGeneralMerchant } = merchant_pool
   value = await _GetGeneralMerchant({ "merchantID" : "23034"  })
   res.locals.body = { "somekey": "somevalue" }
   res.end( JSON.stringify( value ) )

}

async function g_create(req, res) {
   const { _CreateGeneralMerchant } = merchant_pool
   const value = await _CreateGeneralMerchant({ "partnerMerchant" : {}, "generalMerchant": {}  })
   res.end( JSON.stringify( value ) )
}

async function g_delete(req, res) {
   const { _DeleteGeneralMerchant } = merchant_pool
   const value = await _DeleteGeneralMerchant( { "merchantID" : "230-34"  } )
   res.end( JSON.stringify( value ) )
}


/* Partner merchant handlers */

async function p_get(req, res) {
   const { _GetPartnerMerchant } = merchant_pool
   const value = await _GetPartnerMerchant({ "merchantID" : "230-34"  })
   res.end( JSON.stringify( value ) )
}

async function p_create(req, res) {
   const { _CreatePartnerMerchant } = merchant_pool
   const value = await _CreatePartnerMerchant({ "partnerMerchant" : {}, "generalMerchant": {}  })
   res.end( JSON.stringify( value ) )
}

async function p_delete(req, res) {
   const { _DeletePartnerMerchant } = merchant_pool
   const value = await _DeletePartnerMerchant( { "merchantID" : "230-34"  } )
   res.end( JSON.stringify( value ) )
}

module.exports = {

    /* General merchant handlers */

    g_get: {
        prefix: '/general/get',
        handler: g_get,
        verb: 'get'
    },

    g_create: {
        prefix: '/general/create',
        handler: g_create,
        verb: 'post'
    },

    g_delete: {
        prefix: '/general/delete',
        handler: g_delete,
        verb: 'get'
    },

    /* Partner merchant handlers */

    p_get: {
        prefix: '/partner/get',
        handler: p_get,
        verb: 'get'
    },

    p_create: {
        prefix: '/partner/create',
        handler: p_create,
        verb: 'post'
    },

    p_delete: {
        prefix: '/partner/delete',
        handler: p_delete,
        verb: 'get'
    }

}