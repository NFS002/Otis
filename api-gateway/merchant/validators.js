/* Compiled JSON schemas for validating HTTP requests */
const { join } = require("path")
const Ajv = require("ajv")

const schemasHomePath = join(process.env.OTIS_HOME, "dtypes")

/* General Merchant */
const generalMerchantSchema = require(join(schemasHomePath, "generalmerchant/schema/generalmerchant.schema.json"))
/* Partner Merchant */
const partnerMerchantSchema = require(join(schemasHomePath, "partnermerchant/schema/partnermerchant.schema.json"))

/* Merchant Query */
const merchantQuerySchema = require(join(schemasHomePath, "merchant/schema/merchantquery.schema.json"))

/* Merchant expense Bands */
const expenseBandSchema = require(join(schemasHomePath, "merchant/schema/expenseband.schema.json"))


var ajv = new Ajv({
	allErrors: true,
	schemas: [
		generalMerchantSchema,
		partnerMerchantSchema,
		merchantQuerySchema,
		expenseBandSchema
	]
})

module.exports = {
	partnerMerchantValidator: ajv.compile(partnerMerchantSchema),
	merchantQueryValidator: ajv.compile(merchantQuerySchema)
}
