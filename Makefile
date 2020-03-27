PROJECT_NAME := "backend"
PATH := $(./...)

.PHONY: all lint

all: lint

lint: ## Lint the files
	@golint -set_exit_status ./...
