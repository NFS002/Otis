/* Merchant API definition */

const validators = require("./validators.js")
const impl = require("./impl.js")
const { basicJwtCheck } = require("../auth")

module.exports = {

	auth: [basicJwtCheck],

	routes: {
		gGet: {
			prefix: "/general/get",
			handler: impl.gGet,
			verb: "get",
			validators: [validators.merchantQueryValidator],
			scopes: ["read:merchant"]
		},

		gCreate: {
			prefix: "/general/create",
			handler: impl.gCreate,
			verb: "post",
			validators: [validators.partnerMerchantValidator],
			scopes: ["write:merchant"]
		},

		gDelete: {
			prefix: "/general/delete",
			handler: impl.gDelete,
			verb: "delete",
			validators: [validators.merchantQueryValidator],
			scopes: ["delete:merchant"]
		},

		/* Partner merchant handlers */

		pGet: {
			prefix: "/partner/get",
			handler: impl.pGet,
			verb: "get",
			validators: [validators.merchantQueryValidator],
			scopes: ["read:merchant"]
		},

		pCreate: {
			prefix: "/partner/create",
			handler: impl.pCreate,
			verb: "post",
			validators: [validators.partnerMerchantValidator],
			scopes: ["write:merchant"]
		},

		pDelete: {
			prefix: "/partner/delete",
			handler: impl.pDelete,
			verb: "delete",
			validators: [validators.merchantQueryValidator],
			scopes: ["delete:merchant"]
		},

		test: {
		    prefix: "/test",
		    handler: impl.test,
		    verb: "get"
		}
	}
}
