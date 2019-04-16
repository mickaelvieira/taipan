OS    := $(shell uname -s)
SHELL := /bin/bash

build:
	go build

run:
	./taipain

fmt:
	gofmt -s -w ./**/*.go

clean:
	go clean && go mod tidy
