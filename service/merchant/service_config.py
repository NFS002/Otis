""" Configuration file for the merchant service, indexed by the current environment"""
import os
from service.lib.psql.connection import dump_psql

SERVICE_HOME = os.path.join(os.path.join(os.environ["OTIS_HOME"], "service/merchant"))

SERVICE_CONFIG = \
    {
        "development": {

            "network": {
                "address": "localhost",
                "port": 3005
            },

            "rdb_engine": {
                "strategy": "mock",
                "executor": dump_psql
            },

            "logs": {
                "meta": False,
                "rollbar": True,
                "console": True,
                "files": [
                    {
                        "filename": os.path.join(SERVICE_HOME, "logs/combined.log"),
                        "maxsize": 20000000,
                        "backups": 5
                    },
                    {
                        "filename": os.path.join(SERVICE_HOME, "logs/error.log"),
                        "maxsize": 20000000,
                        "backups": 5,
                        "level": "ERROR"
                    }
                ]
            },

            "tls": {
                "use_tls": True,
                "root_dir": os.path.join(os.environ["OTIS_HOME"], "certs"),
                "domain": "merchant.service.slide",
                "root_ca": "Slide-local.crt",
                "private_key": "merchant.service.slide.com.key",
                "cert_chain": "merchant.service.slide.com.crt",
                "verify_client": True
            }
        },
        "production": {

            "network": {
                "address": "0.0.0.0",
                "port": 3005
            },

            "logs": {
                "meta": True,
                "rollbar": True,
                "console": True,
                "files": [
                    {
                        "filename": os.path.join(SERVICE_HOME, "logs/combined.log"),
                        "maxsize": 20000000,
                        "backups": 5
                    },
                    {
                        "filename": os.path.join(SERVICE_HOME, "logs/error.log"),
                        "maxsize": 20000000,
                        "backups": 5,
                        "level": "ERROR"
                    }
                ]
            },

            "tls": {
                "use_tls": True,
                "root_dir": os.path.join(os.environ["OTIS_HOME"], "certs"),
                "domain": "merchant.service.slide.com",
                "root_ca": "Slide-local.crt",
                "private_key": "merchant.service.slide.com.key",
                "cert_chain": "merchant.service.slide.com.crt",
                "verify_client": True
            }
        }
    }[os.environ.get("OTIS_ENV", "development")]
