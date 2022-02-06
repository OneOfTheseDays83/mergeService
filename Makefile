#!/usr/bin/make

.PHONY: help
.SECONDEXPANSION:


################################################################################
# Set common variables
PROJECT                             := merge-service
BUILD_OUTPUT_DIR                    ?= dist
SERVICE_PORT						:= 8000

################################################################################
help: ## Print this help message.
	@echo "Usable make targets:"
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\1:\2/' | column -c2 -t -s : | sort)"

################################################################################
# Build, Package, Test and Code Quality Make Targets

download-deps:
	go mod download -xl
	docker pull mongo

build:
	docker build \
		--network=host \
		-f docker/build.Dockerfile \
		-t "$(PROJECT)" \
		.
build-local:
	cd ./cmd; go build -o ../dist/mergeService

test-local:
	cd ./cmd; go test ./... -coverprofile ../dist/coverage.out

test-input:
	cd ./test; go run input.go 1000

start:
	docker run -p $(SERVICE_PORT):$(SERVICE_PORT) --env SERVICE_PORT=$(SERVICE_PORT) $(PROJECT)

start-with-metrics:
	docker run -p $(SERVICE_PORT):$(SERVICE_PORT) --env SERVICE_PORT=$(SERVICE_PORT) --env PRINT_METRICS="true" $(PROJECT)

start-local:
	SERVICE_PORT=$(SERVICE_PORT) ./dist/mergeService

start-local-with-metrics:
	SERVICE_PORT=$(SERVICE_PORT) PRINT_METRICS="true" ./dist/mergeService

gen-mocks:
