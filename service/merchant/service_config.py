""" Configuration file for the merchant service, indexed by the current environment"""
import os

SERVICE_HOME = os.path.join(os.path.join(os.environ["OTIS_HOME"], "service/merchant"))

SERVICE_CONFIG = \
    {
        "development": {

            "network": {
                "address": "localhost",
                "port": 3005
            },

            "rdb_engine": {
                "echo": False,
                "strategy": "plain"
            },

            "logs": {
                "rdb": {
                    "level": "INFO",
                    "format": "%(asctime)s - %(message)s;",
                    "logger": "sqlalchemy.engine",
                    "file": os.path.join(SERVICE_HOME, "logs/rdb.log")
                },
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

            "rdb_engine": {
                "echo": False
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
