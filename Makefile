ifneq (,$(wildcard ./.env))
    include .env
    export
endif

GIT_VERSION := $(shell git describe --always --tags)
IMAGE_VERSION := $(GIT_VERSION:v%=%)

.PHONY: all
all:
	@echo "make serve"

.PHONY: serve
serve:
	go run . -- serve