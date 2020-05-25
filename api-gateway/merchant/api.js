const ClientPool = require("../connection-pool")
const { getProtoPackage, getValue } = require("../utils.js")
const validators = require("./validators.js")

const merchantService = getValue("services::merchant")
const merchantProtoPkg = getProtoPackage(merchantService.proto_file, merchantService.package, merchantService.proto_dirs)
const merchantPoolOpts = { grpcPkg: merchantProtoPkg, serviceName: merchantService.service, urls: merchantService.addresses, tlsConf: getValue("tls") }

const merchantPool = new ClientPool(merchantPoolOpts)

/* General merchant handlers */

async function gGet (req, res) {
	const { _GetGeneralMerchant } = merchantPool
	var value = await _GetGeneralMerchant({ merchantID: "23034" })
	res.locals.body = value
	res.end(JSON.stringify(value))
}

async function gCreate (req, res) {
	const { _CreateGeneralMerchant } = merchantPool
	const value = await _CreateGeneralMerchant({ partnerMerchant: {}, generalMerchant: {} })
	res.end(JSON.stringify(value))
}

async function gDelete (req, res) {
	const { _DeleteGeneralMerchant } = merchantPool
	const value = await _DeleteGeneralMerchant({ merchantID: "230-34" })
	res.end(JSON.stringify(value))
}

/* Partner merchant handlers */

async function pGet (req, res) {
	const { _GetPartnerMerchant } = merchantPool
	const value = await _GetPartnerMerchant(req.query)
	res.locals.body = value
	res.end(JSON.stringify(value))
}

async function pCreate (req, res) {
	const { _CreatePartnerMerchant } = merchantPool
	const value = await _CreatePartnerMerchant(req.body)
	res.locals.body = value
	res.end(JSON.stringify(value))
}

async function pDelete (req, res) {
	const { _DeletePartnerMerchant } = merchantPool
	const value = await _DeletePartnerMerchant({ merchantID: "230-34" })
	res.end(JSON.stringify(value))
}

module.exports = {

	/* General merchant handlers */

	gGet: {
		prefix: "/general/get",
		handler: gGet,
		verb: "get"
	},

	gCreate: {
		prefix: "/general/create",
		handler: gCreate,
		verb: "post"
	},

	gDelete: {
		prefix: "/general/delete",
		handler: gDelete,
		verb: "get"
	},

	/* Partner merchant handlers */

	pGet: {
		prefix: "/partner/get",
		handler: pGet,
		verb: "get",
		validators: [validators.merchantQueryValidator]
	},

	pCreate: {
		prefix: "/partner/create",
		handler: pCreate,
		verb: "post",
		validators: [validators.partnerMerchantValidator]
	},

	pDelete: {
		prefix: "/partner/delete",
		handler: pDelete,
		verb: "get"
	}

}
