# backend


[![pipeline status](https://gitlab.com/otis-team/backend/badges/master/pipeline.svg)](https://gitlab.com/otis-team/backend/-/commits/master)
[![coverage report](https://gitlab.com/otis-team/backend/badges/master/coverage.svg)](https://gitlab.com/otis-team/backend/-/commits/master)

Otis backend

## Directory structure

- api (contains any Micro API services that handle requests from API Gateway)
- service (contains all Micro services)
- client (contains some clients for respective services)

## Local Development

To clone:

`git clone git@gitlab.com:otis-team/backend.git` 

IMPORTANT! Make sure you ammend your git config file to include your name and OTIS EMAIL! 

To run:

`docker-compose build`

`docker-compose run`

## Namespaces

The namespace for Go Modules should follow the directory structure of this repository.

i.e. the module name for the merchant service is *gitlab.com/otis-team/backend/service/merchant* 

This helps with future importing.

The namespace for services within the Micro runtime is:

go.micro.[api/service/client].[name]

This is to provide integration with the default Micro namespace. This will likely be changed in the future.