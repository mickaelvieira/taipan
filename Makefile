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
