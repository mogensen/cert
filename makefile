# Vars
DOCKER_REPO=mogensen
APP_NAME=gofish-bots
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

run:  ## run
	go run ./cmd/cert github.com

install:  ## install
	go install ./cmd/cert

test-coverage: ## Generate test coverage report
	mkdir -p tmp
	go test ./... --coverprofile tmp/outfile
	go tool cover -html=tmp/outfile

lint: ## Run golint on all go files
	golint