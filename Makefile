OS    := $(shell uname -s)
SHELL := /bin/bash

build-app: build build-ui

build: build-web build-feeds build-migration

build-web:
	cd cmd/web && go build

build-feeds:
	cd cmd/feeds && go build

build-migration:
	cd cmd/migration && go build

run:
	cd cmd/web && ./web

run-feeds:
	cd cmd/feeds && ./feeds

analyse:
	staticcheck cmd/web/main.go
	staticcheck cmd/feeds/main.go
	staticcheck cmd/migration/main.go

# run-migration:
# 	cd cmd/migration && ./migration

fmt:
	gofmt -s -w ./**/*.go

clean:
	go mod tidy
	cd cmd/web && go clean
	cd cmd/feeds && go clean
	cd cmd/migration && go clean

clean-ui:
	rm -rf web/app/node_modules
	rm -f web/app/schema.json
	rm -rf web/static/js
	rm -f web/static/hashes.json

watch-ui:
	cd web/app && yarn watch:client

build-ui:
	cd web/app && yarn && yarn build:client

gen-schema:
	cd web/app && yarn schema:json
