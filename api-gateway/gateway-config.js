const { join } = require("path")
/* const auth = require("./auth") */

module.exports = {
	development: {

		address: "localhost",
		port: 3000,

		apis: {
			merchant: {
			    name: "merchant",
				path: "./merchant",
				prefix: "/merchant",
				overrides: {
				    auth: false,
				    validators: false
				}
			}
		},

		global_auth: [/* auth.basicJwtCheck */],

		morgan: true,
		morgan_theme: "inverted",

		logs: {
			rollbar: true,
			console: true,
			pretty_print: true,
			meta: false,
			files: [
				{
					filename: join(process.env.OTIS_HOME, "api-gateway/logs/combined.log"),
					maxsize: 1024 * 1024 * 20
				},
				{
					level: "error",
					maxsize: 1024 * 1024 * 20,
					filename: join(process.env.OTIS_HOME, "api-gateway/logs/error.log")
				}
			]
		},

		services: {
			merchant: {
				addresses: ["localhost:3005"],
				name: "merchant",
				proto_dirs: [process.env.OTIS_HOME],
				proto_file: join(process.env.OTIS_HOME, "lib/proto/merchant/merchant.proto"),
				package: "merchant",
				service: "MerchantService"
			}
		},

		tls: {
		    use_tls: true,
		    root_dir: join(process.env.OTIS_HOME, "certs"),
			domain: "localhost",
			domain_override: "merchant.service.slide.com",
			root_ca: "Slide-local.crt",
			private_key: "gateway.slide.com.key",
			cert_chain: "gateway.slide.com.crt"
		}
	},
	production: {
		address: "0.0.0.0",
		port: 3000,

		apis: {
			merchant: {
				name: "merchant",
				path: "./merchant",
				prefix: "/merchant",
				overrides: {
					auth: false,
					validators: false
				}
			}
		},

		global_auth: [/* auth.basicJwtCheck */],

		morgan: true,
		morgan_theme: "inverted",

		logs: {
			rollbar: true,
			console: false,
			pretty_print: false,
			meta: true,
			files: [
				{
					filename: join(process.env.OTIS_HOME, "api-gateway/logs/combined.log"),
					maxsize: 1024 * 1024 * 20
				},
				{
					level: "error",
					maxsize: 1024 * 1024 * 20,
					filename: join(process.env.OTIS_HOME, "api-gateway/logs/error.log")
				}
			]
		},

		services: {
			merchant: {
				name: "merchant",
				addresses: ["localhost:3005"],
				proto_dirs: [process.env.OTIS_HOME],
				proto_file: join(process.env.OTIS_HOME, "lib/proto/merchant/merchant.proto"),
				package: "merchant",
				service: "MerchantService"
			}
		},

		tls: {
		    use_tls: true,
			root_dir: join(process.env.OTIS_HOME, "certs"),
			domain_override: "localhost",
			domain: "gateway.slide.com",
			root_ca: "Slide-local.crt",
			private_key: "gateway.slide.com.key",
			cert_chain: "gateway.slide.com.crt"
		}
	}
}
