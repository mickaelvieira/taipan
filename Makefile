OS    := $(shell uname -s)
SHELL := /bin/bash
GOFMT := gofmt -s -w -l
CDWEB := cd web/app

build: build-app build-ui

test: test-app test-ui

run:
	./taipan web

run-feeds:
	./taipan syndication

run-documents:
	./taipan documents

build-app:
	go build

test-app:
	go test ./...

build-ui:
	$(CDWEB) && yarn && yarn build

test-ui:
	$(CDWEB) && yarn test

watch-ui:
	$(CDWEB) && yarn watch

watch-test-ui:
	$(CDWEB) && yarn watch-test

fmt:
	$(GOFMT) taipan.go
	$(GOFMT) internal/**/*.go
	$(GOFMT) cmd/*.go
	$(CDWEB) && yarn lint:fix

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

