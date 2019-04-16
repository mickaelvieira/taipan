OS    := $(shell uname -s)
SHELL := /bin/bash

build:
	cd cmd/web && go build

run:
	cd cmd/web && ./web

fmt:
	gofmt -s -w ./**/*.go

clean:
	cd cmd/web && go clean && go mod tidy

clean-ui:
	rm -rf web/static/css && rm -rf web/static/js && rm -f web/static/hashes.json
