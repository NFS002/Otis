/* Merchant API handlers */
const ClientPool = require("../connection-pool")
const { getProtoPackage, getValue, getReqData, throwForProperty } = require("../utils.js")

const merchantService = getValue("services::merchant")
const merchantProtoPkg = getProtoPackage(merchantService.proto_file, merchantService.package, merchantService.proto_dirs)
const merchantPoolOpts = { grpcPkg: merchantProtoPkg, serviceName: merchantService.service, urls: merchantService.addresses, tlsConf: getValue("tls") }

const merchantPool = new ClientPool(merchantPoolOpts)

module.exports = {

	/* General merchant handlers */
	gGet: async function (req, res) {
		const { _GetGeneralMerchant } = merchantPool
		var value = await _GetGeneralMerchant(getReqData(req))
		throwForProperty(value, "executed")
		res.locals.body = value
		res.end(JSON.stringify(value))
	},

	gCreate: async function (req, res) {
		const { _CreateGeneralMerchant } = merchantPool
		var value = await _CreateGeneralMerchant({ partnerMerchant: req.body })
		throwForProperty(value, "executed")
		res.locals.body = value
		res.end(JSON.stringify(value))
	},

	gDelete: async function (req, res) {
		const { _DeleteGeneralMerchant } = merchantPool
		var value = await _DeleteGeneralMerchant(getReqData(req))
		throwForProperty(value, "executed")
		res.locals.body = value
		res.end(JSON.stringify(value))
	},

	/* Partner merchant handlers */
	pGet: async function (req, res) {
		const { _GetPartnerMerchant } = merchantPool
		var value = await _GetPartnerMerchant(getReqData(req))
		throwForProperty(value, "executed")
		res.locals.body = value
		res.end(JSON.stringify(value))
	},

	pCreate: async function (req, res) {
		const { _CreatePartnerMerchant } = merchantPool
		var value = await _CreatePartnerMerchant({ partnerMerchant: req.body })
		throwForProperty(value, "executed")
		res.locals.body = value
		res.end(JSON.stringify(value))
	},

	pDelete: async function (req, res) {
		const { _DeletePartnerMerchant } = merchantPool
		var value = await _DeletePartnerMerchant(getReqData(req))
		throwForProperty(value, "executed")
		res.locals.body = value
		res.end(JSON.stringify(value))
	}
}
