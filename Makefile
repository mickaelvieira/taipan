OS    := $(shell uname -s)
SHELL := /bin/bash

build-app: build build-ui

build: build-web build-feeds build-documents

build-web:
	cd cmd/web && go build

build-feeds:
	cd cmd/feeds && go build

build-documents:
	cd cmd/documents && go build

build-migration:
	cd cmd/migration && go build

run:
	cd cmd/web && ./web

run-feeds:
	cd cmd/feeds && ./feeds

run-documents:
	cd cmd/documents && ./documents

gen-proto:
	protoc --proto_path=web/proto --go_out=internal/domain/document web/proto/document.proto

analyse:
	staticcheck cmd/web/main.go
	staticcheck cmd/feeds/main.go
	staticcheck cmd/migration/main.go

# run-migration:
# 	cd cmd/migration && ./migration

fmt:
	gofmt -s -w -l internal/**/*.go
	gofmt -s -w -l cmd/**/*.go

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
