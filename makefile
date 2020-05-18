NAME     := mrbotchi
VERSION  := v0.1.0
REVISION := $(shell git rev-parse --short HEAD)
CODENAME := ASAYAKE

SRCS:=$(shell find . -type f -name '*.go')
LDFLAGS:= -ldflags="-s -w -X \"main.version=$(VERSION)\" -X \"main.revision=$(REVISION)\" -X \"main.codeName=$(CODENAME)\" -extldflags \"-static\""

bin/$(NAME): $(SRCS)
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME).out

.PHONY: debug
debug:
	dlv debug main.go -- -c ./etc/mrbotchi/config.toml

.PHONY: clean
clean:
	rm -rf bin/*
