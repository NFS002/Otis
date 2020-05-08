module.exports = {
  "development": {

    "address":"127.0.0.1",
    "port": 3000,

    "apis": {
        "merchant": {
            "prefix": "/merchant",
            "name": "merchant",
            "path": "./merchant"
        }
    },

    "services": {
      "merchant": {
        "name": "merchant",
        "address": "127.0.0.1:3005",
        "proto_dirs": [ process.env["OTIS_HOME"], "proto", '.'],
        "proto_file": "example.proto",
        "package": "example",
        "service": "ExampleService"
      }
    }
  },
  "production": {
    "address":"0.0.0.0",
    "port": 80,

    "apis": {
        "merchant": {
            "name": "merchant",
            "path": "./merchant",
            "prefix": "/merchant"
        }
    },

    "services": {
      "merchant": {
          "name": "merchant",
          "address": "127.0.0.1:3005",
          "proto_dirs": [ process.env["OTIS_HOME"], "proto", '.' ],
          "proto_file": "example.proto",
          "package": "example",
          "service": "ExampleService"
        }
    }
  }
}