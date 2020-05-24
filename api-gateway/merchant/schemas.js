const Joi = require("@hapi/joi")

/* Schemas for validating HTTP requests */

const Exported = {
	merchantQuerySchema: Joi.object({
		merchantID: Joi.string().alphanum()
	})
}

module.exports = Exported
