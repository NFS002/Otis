module.exports = {
	development: {

		address: "127.0.0.1",
		port: 3000,

		apis: {
			merchant: {
				prefix: "/merchant",
				name: "merchant",
				path: "./merchant"
			}
		},

		logs: {
			rollbar: true,
			console: false,
			pretty_print: true,
			files: [
				{
					filename: process.env.OTIS_HOME + "./api-gateway/logs/combined.log",
					maxsize: 1024 * 1024 * 20
				},
				{
					level: "error",
					maxsize: 1024 * 1024 * 20,
					filename: process.env.OTIS_HOME + "./api-gateway/logs/error.log"
				}
			]
		},

		services: {
			merchant: {
				addresses: ["localhost:3005"],
				name: "merchant",
				address: process.env.OTIS_SERVICE_MERCHANT_ADDRESS,
				proto_dirs: [process.env.OTIS_HOME],
				proto_file: "./service/merchant/proto/merchant/merchant.proto",
				package: "merchant",
				service: "MerchantService"
			}
		}
	},
	production: {
		address: "0.0.0.0",
		port: 3000,

		apis: {
			merchant: {
				name: "merchant",
				path: "./merchant",
				prefix: "/merchant"
			}
		},

		logs: {
			rollbar: true,
			console: false,
			pretty_print: false,
			files: [
				{
					filename: process.env.OTIS_HOME + "./api-gateway/logs/combined.log",
					maxsize: 1024 * 1024 * 20
				},
				{
					level: "error",
					maxsize: 1024 * 1024 * 20,
					filename: process.env.OTIS_HOME + "./api-gateway/logs/error.log"
				}
			]
		},

		services: {
			merchant: {
				name: "merchant",
				addresses: ["localhost:3005"],
				proto_dirs: [process.env.OTIS_HOME],
				proto_file: "./service/merchant/proto/merchant/merchant.proto",
				package: "merchant",
				service: "MerchantService"
			}
		}
	}
}
