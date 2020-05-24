/* Compiled JSON schemas for validating HTTP requests */
const { join } = require("path")
const Ajv = require("ajv")

var ajv = new Ajv({ allErrors: true })
const schemasHome = join(process.env.OTIS_HOME, "dtypes")


const Exported = {
	merchantQueryValidator: ajv.compile(
	    require(
	        join(schemasHome, "merchant/schema/merchantquery.schema.json")
	    )
	)
}

module.exports = Exported
