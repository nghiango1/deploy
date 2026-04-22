.PHONY: all
 all: help

GIT_TAG := $(shell git rev-parse --short HEAD)
PACKAGE_NAME := webhook-handler
SERVICE_IMAGE := docker.io/nghiango1/$(PACKAGE_NAME):$(GIT_TAG)

### Build command

.PHONY: build
build: # Build the webhook-handler
	mkdir -p dist
	go build -o dist/${PACKAGE_NAME} cmd/${PACKAGE_NAME}/main.go

.PHONY: help
help: # Show this help
	@cat Makefile | \
		grep -E '^[^.:[:space:]]+:|[#]##' | \
		sed -E 's/:[^#]*#([^:]+)$$/: #:\1/' | \
		sed -E 's/([^.:[:space:]]+):([^#]*#(.+))?.*/    make \1\3/' | \
		sed -E 's/[#][#]# *(.+)/# \1/' | \
		column -ts: -L -W2
