/* Merchant API definition */

const validators = require("./validators.js")
const impl = require("./impl.js")

module.exports = {

	gGet: {
		prefix: "/general/get",
		handler: impl.gGet,
		verb: "get",
		validators: [validators.merchantQueryValidator]
	},

	gCreate: {
		prefix: "/general/create",
		handler: impl.gCreate,
		verb: "post",
		validators: [validators.partnerMerchantValidator]
	},

	gDelete: {
		prefix: "/general/delete",
		handler: impl.gDelete,
		verb: "get",
		validators: [validators.merchantQueryValidator]
	},

	/* Partner merchant handlers */

	pGet: {
		prefix: "/partner/get",
		handler: impl.pGet,
		verb: "get",
		validators: [validators.merchantQueryValidator]
	},

	pCreate: {
		prefix: "/partner/create",
		handler: impl.pCreate,
		verb: "post",
		validators: [validators.partnerMerchantValidator]
	},

	pDelete: {
		prefix: "/partner/delete",
		handler: impl.pDelete,
		verb: "get",
		validators: [validators.merchantQueryValidator]
	}

}
