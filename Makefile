OS    := $(shell uname -s)
SHELL := /bin/bash

build:
	cd cmd/web && go build

run:
	cd cmd/web && ./web

build-migration:
	cd cmd/migration && go build

# run-migration:
# 	cd cmd/migration && ./migration

fmt:
	gofmt -s -w ./**/*.go

clean:
	cd cmd/web && go clean && go mod tidy

clean-ui:
	rm -rf web/static/css && rm -rf web/static/js && rm -f web/static/hashes.json

watch-ui:
	cd web/app && yarn watch:client

gen-schema:
	cd web/app && yarn schema:json
