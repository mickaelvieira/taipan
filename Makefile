OS    := $(shell uname -s)
SHELL := /bin/bash

build: build-app build-ui

test: test-app test-ui

run:
	./taipan web

run-feeds:
	./taipan feeds

run-documents:
	./taipan documents

build-app:
	go build

test-app:
	go test ./...

build-ui:
	cd web/app && yarn && yarn build

test-ui:
	cd web/app && yarn test

watch-ui:
	cd web/app && yarn watch

watch-test-ui:
	cd web/app && yarn watch-test

fmt:
	gofmt -s -w -l internal/**/*.go
	gofmt -s -w -l cmd/**/*.go
	cd web/app && yarn lint:fix

clean:
	go mod tidy
	go clean
	rm -rf web/app/node_modules
	rm -f web/app/schema.json
	rm -rf web/static/js
	rm -f web/static/hashes.json

analyse:
	staticcheck taipan.go
	cd web/app && yarn lint

gen-proto:
	protoc --proto_path=web/proto --go_out=internal/domain/document web/proto/document.proto

gen-schema:
	cd web/app && yarn gen:graphql:schema


# build-migration:
# 	cd cmd/migration && go build

# run-migration:
# 	cd cmd/migration && ./migration

