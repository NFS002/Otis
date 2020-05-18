const ClientPool = require("../connection-pool")
const { getProtoPackage, getValue } = require("../utils.js")

const merchantService = getValue("services::merchant")
const merchantProtoPkg = getProtoPackage(merchantService.proto_file, merchantService.package, merchantService.proto_dirs)
const merchantPoolOpts = { grpcPkg: merchantProtoPkg, serviceName: merchantService.service, urls: merchantService.addresses }

const merchantPool = new ClientPool(merchantPoolOpts)

/* General merchant handlers */

async function gGet (req, res) {
	var value
	const { _GetGeneralMerchant } = merchantPool
	value = await _GetGeneralMerchant({ merchantID: "23034" })
	res.locals.body = { somekey: "somevalue" }
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
	const value = await _GetPartnerMerchant({ merchantID: "230-34" })
	res.end(JSON.stringify(value))
}

async function pCreate (req, res) {
	const { _CreatePartnerMerchant } = merchantPool
	const value = await _CreatePartnerMerchant({ partnerMerchant: {}, generalMerchant: {} })
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
		verb: "get"
	},

	pCreate: {
		prefix: "/partner/create",
		handler: pCreate,
		verb: "post"
	},

	pDelete: {
		prefix: "/partner/delete",
		handler: pDelete,
		verb: "get"
	}

}
