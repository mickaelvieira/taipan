OS     := $(shell uname -s)
SHELL  := /bin/bash
GOFMT  := gofmt -s -w -l
GOLINT := golint
GOVET  := go vet
GOSHDW := go vet -vettool=$$(which shadow)
GOSEC  := gosec --quiet
CDWEB  := cd web/app
RMSCRIPTS := rm -rf web/static/js/
RMSTYLES := rm -rf web/static/css/

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
	$(RMSCRIPTS) && $(RMSTYLES) && $(CDWEB) && yarn && yarn build

test-ui:
	$(CDWEB) && yarn test

watch-ui:
	$(RMSCRIPTS) && $(RMSTYLES) && $(CDWEB) && yarn watch

watch-test-ui:
	$(CDWEB) && yarn watch-test

watch-e2e:
	$(CDWEB) && yarn watch-e2e

test-e2e:
	$(CDWEB) && yarn test-e2e

fmt:
	$(GOFMT) *.go
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

# https://github.com/actions/setup-go/issues/27
lint:
	$(GOLINT) ./...
	$(GOVET) ./...
	$(GOSEC) ./...
	$(GOSHDW) ./...
	# staticcheck taipan.go
	cd web/app && yarn lint

gen-proto:
	protoc --proto_path=web/proto --go_out=internal/domain/messages web/proto/document.proto

gen-schema:
	cd web/app && yarn gen:graphql:schema


# build-migration:
# 	cd cmd/migration && go build

# run-migration:
# 	cd cmd/migration && ./migration
