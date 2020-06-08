# backend


[![pipeline status](https://gitlab.com/otis_team/backend/badges/master/pipeline.svg)](https://gitlab.com/otis_team/backend/-/commits/master)
[![coverage report](https://gitlab.com/otis_team/backend/badges/master/coverage.svg)](https://gitlab.com/otis_team/backend/-/commits/master)

Otis backend

## Requirements
- gRPC & protobuf compiler (only required for development, rather than running)
- Python >= 3.5
- Node.js >= 10
- Npm >= 6
- Tested on Mac OS Mojave (10.14.1)


## Directory structure:

- api-gateway (Node.js/Express API Gateway)
- service (Contains all Python/gRPC microservices)
- lib/service (Shared library for each microservice )
- lib/types (Implementations of common data types)
- lib/proto (gRPC Mictoservice protobuf definitions and compoled *.proto files)
- lib/db (Database utilities and migrations)

## Setting up a local development environment


### Environment variables
Source the shell script .env to set up your environment variables for local development.
The file env.example.sh serves as a template for env.sh, but does not include any sensitive information

```shell script
$ cp .env.example .env 
# Then manually edit .env to include values for your own set of credentials
$ source .env
```

### Python virtual environement
Creating a virtual environment for each microservice is not necessary but recommended.
```shell script
$ cd service/<service_name>
$ python3 -m venv <virtual_env_name>
```
Then, to activate the virtual environment and install dependencies before running the service:
```shell script
$ source <virtual_env_name>/bin/activate
$ pip install -r requirements.txt
$ deactivate # Deactivates the current virtual environment
```

### Node.js dependencies
To install all node.js dependencies required for the api-gateway, we use
npm:

```shell script
$ cd api-gateway
$ npm install
```



## Configuration
Configuration specifically for the api-gateway can be made through
a control file at *api-gateway/gateway-config.js*. Similarly, there is a configuration file
for each microservice at service/<service_name>/service-config.py.
You will probably need to manually edit these files to meet your own requirements
before running the api-gateway or any of individual microservices

## Running the api-gateway

```shell script
$ node .
```

## Running a microservice

```shell script
$ cd service/<service_name>
# Activate the virtual environment, if necessary
$ python3 .
```

## TLS/SSL
By default, we use mutual TLS, to encrypt all gRPC connections, where both the client and server (in our case the api-gateway and a microservice)
provide certificates and private RSA keys, and each certificate is verified by the certificate authority of the other. Of course, the private keys or the certificates 
are not included in this repository, so you if you want to use tls you will need to generate your own, and place them in the configured directory (See 'Configuration' section above).
To do that, you can follow the tutorial here (https://jsherz.com/grpc/node/nodejs/mutual/authentication/ssl/2017/10/27/grpc-node-with-mutual-auth.html)
or alternatively use the open-ssl CLI tool. Using TLS is not strictly necessary for a local development environment, so if you want to use unencrypted and unauthenticated 
connections instead, you must set the "use_tls" variable in the api-gateway configuration file, and the configuration file of *every* microservice to false.
For example:
```
{
...
    "use_tls": false,
...
}
```
See the 'Configuration' section for more details.

## Code linting
To help maintain code quality standars and consistency, we use pylint and pycodestyle (formerly pep8) for each microservice and the shared library,
as well as eslint for the api-gateway. Comfiguration files for code linting can be found in *~/service/pylintrc*, 
*~/service/tox.ini*, and *~/api-gateway/.eslintrc.json* respectively.

To run the code linting programs:

```shell script
$ eslint --debug api-gateway
$ pylint service/merchant lib/service --rcfile=service/pylintrc
$ pycodestyle service/merchant lib/service --config=service/tox.ini
```

## System architecture diagrams, data policies, and further documentation are available in a seperate repository at
### https://gitlab.com/otis_team/docs