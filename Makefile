# cat -e -t -v Makefile

PROJECT_NAME := "backend"

.PHONY: all lint

all: lint-all

lint-all: eslint-api-gateway merchant-service-lint-all lib-lint-all

merchant-service-lint-all: pylint-merchant-service pycodestyle-merchant-service

lib-lint-all: pylint-lib pycodestyle-lib

pylint-merchant-service:
	cd ${OTIS_HOME}
	source "service/merchant/merchant-service-venv/bin/activate"
	python3 -m pip install -r "service/merchant/requirements.txt"
	pylint "service/merchant" --rcfile="service/pylintrc"

pycodestyle-merchant-service:
	cd ${OTIS_HOME}
	source "service/merchant/merchant-service-venv/bin/activate"
	python3 -m pip install -r "service/merchant/requirements.txt"
	pycodestyle service/merchant --config=service/tox.ini

pylint-lib:
	echo "pylint-lib"

pycodestyle-lib:
	echo "pycodestyle-lib"

eslint-api-gateway:
	echo "eslint-api-gateway"

