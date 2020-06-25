# Run from ${OTIS_HOME} directory
# cat -e -t -v Makefile

PROJECT_NAME := "backend"


.PHONY: all lint

all: lint-all

lint-all: eslint-api-gateway merchant-service-lint-all lib-lint-all

merchant-service-lint-all: pylint-merchant-service pycodestyle-merchant-service

lib-lint-all: pylint-lib pycodestyle-lib

pylint-merchant-service:
	source "service/merchant/merchant-service-venv/bin/activate"
	python3 -m pip install -r "service/merchant/requirements.txt"
	pylint "service/merchant" --rcfile="service/pylintrc"

pycodestyle-merchant-service:
	source "service/merchant/merchant-service-venv/bin/activate"
	python3 -m pip install -r "service/merchant/requirements.txt"
	pycodestyle "service/merchant" --config="service/tox.ini"

pylint-lib:
	source "service/merchant/merchant-service-venv/bin/activate"
	python3 -m pip install -r "service/merchant/requirements.txt"
	pylint "lib/service" --rcfile="service/pylintrc"

pycodestyle-lib:
	source "service/merchant/merchant-service-venv/bin/activate"
	python3 -m pip install -r "service/merchant/requirements.txt"
	pycodestyle "lib/service" --config="service/tox.ini"

eslint-api-gateway:
	cd api-gateway
	npm install -g eslint
	eslint .

docker-login:
	@echo "docker login..."
	@docker login -u ${OTIS_CI_REGISTRY_USER} -p ${OTIS_CI_REGISTRY_PASSWORD} ${OTIS_CI_REGISTRY}

docker-build:
	docker-compose build

docker-push: docker-build docker-login
	docker-compose push




