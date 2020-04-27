# backend


[![pipeline status](https://gitlab.com/otis_team/backend/badges/master/pipeline.svg)](https://gitlab.com/otis_team/backend/-/commits/master)
[![coverage report](https://gitlab.com/otis_team/backend/badges/master/coverage.svg)](https://gitlab.com/otis_team/backend/-/commits/master)

Otis backend

## Directory structure

- api (contains any Micro API services that handle requests from API Gateway)
- service (contains all Micro services)
- client (contains some clients for respective services)

## Local Development

To clone:

`git clone git@gitlab.com:otis_team/backend.git` 

IMPORTANT! Make sure you ammend your git config file to include your name and OTIS EMAIL! 

To run:

`docker-compose build`

`docker-compose run`

## Namespaces

The namespace for Go Modules should follow the directory structure of this repository.

i.e. the module name for the merchant service is *gitlab.com/otis_team/backend/service/merchant* 

This helps with future importing.

The namespace for services within the Micro runtime is:

go.micro.[api/service/client].[name]

This is to provide integration with the default Micro namespace. This will likely be changed in the future.

## Golang local modules
All go.mod files that require another local module use
the replace directive as follows replace gitlab.com/.. => /path/to/local/directory/...
However, for deployment, we have to use the remote modules rather than the local ones, so we need to remove these directives.
The shell script scripts/mod-deploy.sh searches all go.mod files and comments out any lines that end with '//_LOCAL'. This script should be run after
pushing to a remote repo, or before deployment. Conversely, the file scripts/mod-local.sh does the opposite, and uncomments
any lines in any go.mod files that are commented out, and that end with '//_LOCAL'. 

### <a href='https://thewebivore.com/using-replace-in-go-mod-to-point-to-your-local-module/'> See this article for details </a>

The shell script scripts/mod-require-delete.sh should then be used
after scripts/mod-deploy.sh and before 'git push'/deployment which deletes every 'require block' in every go.mod file and hence clears any dependencies on local module versions.

## Dev environment

Source the shell script otis.env to set up your environment variables for local development.
The file env.example.sh serves as a template for env.sh, but does not include any sensitive information
For example, after pulling, or before building a docker image:
```shell script
$ source .env
```
Before pushing:
```shell script
$ cp .env .env.example # Then edit .env to remove sensitive information such as passwords
```
