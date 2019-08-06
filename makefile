NAME := mrbotchi
VERSION := v0.1.0

SRCS:=$(shell find . -type f -name '*.go')
LDFLAGS:= -ldflags="-s -w -X \"main.version=$(VERSION)\" -extldflags \"-static\""

bin/$(NAME): $(SRCS)
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)

.PHONY: deps
deps:
	dep ensure -v -vendor-only=true

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*
