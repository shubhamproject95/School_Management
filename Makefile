VERBOSE_ORIGINS := "command line" "environment"
ifdef V
  ifeq ($(filter $(VERBOSE_ORIGINS),$(origin V)),)
    BUILD_VERBOSE := $(V)
  endif
endif

ifndef BUILD_VERBOSE
  BUILD_VERBOSE := 0
endif

ifeq ($(BUILD_VERBOSE),1)
  Q :=
else
  Q := @
endif

ifneq ($(GIT_TAG),)
  IMAGE_TAG = $(GIT_TAG)
else ifeq ($(GIT_BRANCH),main)
  IMAGE_TAG = "latest"
else ifneq ($(GIT_BRANCH),)
  IMAGE_TAG = $(GIT_BRANCH)
endif

PHONY += all
CURDIR := $(shell pwd)
GOCMD=go
GOBUILD=$(GOCMD) build
BUILD_DIR ?= $(CURDIR)/build
GOBIN_DIR := $(BUILD_DIR)/bin
GOBIN_EXAMPLE_DIR := $(BUILD_DIR)/bin/examples
HOST_DIR := $(BUILD_DIR)/host
HOSTBIN_DIR := $(HOST_DIR)/bin
GOTOOLSBIN_DIR := $(HOSTBIN_DIR)
GOTOOLS_DIR := $(CURDIR)/tools
TMP_DIR := $(BUILD_DIR)/tmp
DIRS := \
	$(GOBIN_DIR) \
	$(HOST_DIR) \
	$(HOSTBIN_DIR) \
	$(TMP_DIR) \
	$(GOBIN_EXAMPLE_DIR)

HOST_OS := $(shell uname -s)

REV ?= $(shell git rev-parse --short HEAD 2> /dev/null)
GOPATH ?= $(shell go env GOPATH)

export PATH:=$(GOTOOLS_DIR):$(HOSTBIN_DIR):$(PATH)
export REV

define find-subdir
$(shell find $(1) -maxdepth 1 -mindepth 1 -type d -o -type l)
endef

APPS := $(sort $(notdir $(call find-subdir,cmd)))
PHONY += $(APPS)

#
# Project-specific variables
#
# Base service name. Used for binary name as well
SERVICE=mojito-api

# Service name used in Docker-compose
DOCKER_COMPOSE_SERVICE=$(SERVICE)-service

# Change it to $(SERVICE) after migration to the new deployment process (from TF to Ansible)
DOCKER_IMAGE=api
K8S_DEPLOYMENT_NAME=$(DOCKER_IMAGE)

# Enable Go Modules.
GO111MODULE=on
# Enable Go Proxy.
GOPROXY=https://proxy.golang.org
# Private Go dependency.
GOPRIVATE=github.com/mojitoinc

#
# General variables
#
# Path to Docker file
PATH_DOCKER_FILE=$(realpath ./docker/Dockerfile)
# Path to Docker file for CI
PATH_DOCKER_FILE_CI=$(realpath ./Dockerfile)
# Path to docker-compose file
PATH_DOCKER_COMPOSE_FILE=$(realpath ./docker/docker-compose.yml)
# Docker compose starting options.
DOCKER_COMPOSE_OPTIONS= -f $(PATH_DOCKER_COMPOSE_FILE)
# Docker repository path
DOCKER_REGISTRY_ENTRY=$(GCR_NAME)/mojitoinc/mojito/$(DOCKER_IMAGE)
# K8S namespace
K8S_NAMESPACE=mojito-backend-dev

$(DIRS) :
	$(Q)mkdir -p $@

.SECONDEXPANSION:
$(EXAMPLES): $(addprefix $(GOBIN_EXAMPLE_DIR)/,$$@)


.SECONDEXPANSION:
$(APPS): $(addprefix $(GOBIN_DIR)/,$$@)
$(GOBIN_DIR)/%: $(GOBIN_DIR) FORCE
	$(Q)go build -o $@ ./cmd/$(notdir $@)
	@echo "Done building"
	@echo "Run \"$(subst $(CURDIR),.,$@)\" to launch $(notdir $@)."

CODEGEN_DEPS := \
	# $(GOTOOLS_DIR)/abigen

.PHONY: build
build: ## Build service
	$(GOBUILD) -o ./build/server cmd/server.go
	$(GOBUILD) -o ./build/nats_consumers cmd/nats_consumers.go

.PHONY: build-ci
build-ci: ## Build service for CI
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -ldflags="-w -s" -a -installsuffix cgo -o ./build/server cmd/server.go
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -ldflags="-w -s" -a -installsuffix cgo -o ./build/nats_consumers cmd/nats_consumers.go

.PHONY: generate
generate: $(CODEGEN_DEPS) ## Run code generation
	$(Q)go generate -x ./...

.PHONY: graphql
graphql: $(CODEGEN_DEPS)  ## Generate GraphQL resolver from schema
	go get github.com/99designs/gqlgen/cmd@v0.13.0
	$(Q)go generate -x ./pkg/graphql/resolver.go

.PHONY: go-get
go-get: ## Get go modules locally
	@echo '>>> Getting go modules'
	@env GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) go mod download

.PHONY: go-build
go-build: ## Build service binary locally
	@echo '>>> Building go binary'
	@env GOOS=linux GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) go build -ldflags="-s -w" -o $(SERVICE) cmd/server.go

.PHONY: go-test-integration
go-test-integration: ## Run integration tests locally
	@echo ">>> Running integration tests"
	@env GOPROXY=$(GOPROXY) GOPRIVATE=$(GOPRIVATE) go test -v -p 1 -tags="integration" ./tests/integration/...

.PHONY: service-build
service-build: ## Build service and all it's dependencies in Docker
	@env PRIVATE_REPO_KEY="$$(cat ~/.ssh/id_rsa)" docker-compose $(DOCKER_COMPOSE_OPTIONS) build --no-cache

.PHONY: start-service-dependencies
service-start-dependencies: ## Start service dependencies in Docker
	@echo ">>> Start all Service dependencies"
	@env PRIVATE_REPO_KEY="$$(cat ~/.ssh/id_rsa)" docker-compose $(DOCKER_COMPOSE_OPTIONS) up \
	-d \
	mojito-api-nats mojito-api-postgres mojito-mock-server

.PHONY: service-start
service-start: service-build service-start-dependencies ## Start service in Docker
	@echo ">>> Sleeping 15 seconds until dependencies start"
	@sleep 15
	@echo ">>> Starting service"
	@echo ">>> Starting up service container in Docker"
	@env PRIVATE_REPO_KEY="$$(cat ~/.ssh/id_rsa)" docker-compose $(DOCKER_COMPOSE_OPTIONS) up -d $(DOCKER_COMPOSE_SERVICE)

.PHONY: service-stop
service-stop: ## Stop service in Docker
	@echo ">>> Stopping service"
	@docker-compose $(DOCKER_COMPOSE_OPTIONS) down -v --remove-orphans

.PHONY: service-restart
service-restart: service-stop service-start ## Restart service in Docker

.PHONY: service-integration-test
service-integration-test: service-stop service-start ## Run integration tests in Docker
	@echo ">>> Running integration tests in Docker"
	@env PRIVATE_REPO_KEY="$$(cat ~/.ssh/id_rsa)" docker-compose $(DOCKER_COMPOSE_OPTIONS) \
		run mojito-api-integration-tests

.PHONY: build-and-run-integration-tests
build-and-run-integration-tests: ## Run integration tests locally, check out tests/README.md for more info.
	@echo ">>> Rebuilding integration test image"
	@env PRIVATE_REPO_KEY="$$(cat ~/.ssh/id_rsa)" docker-compose $(DOCKER_COMPOSE_OPTIONS) \
		build mojito-api-integration-tests
	@echo ">>> Running integration test in Docker"
	@env PRIVATE_REPO_KEY="$$(cat ~/.ssh/id_rsa)" docker-compose $(DOCKER_COMPOSE_OPTIONS) \
		run mojito-api-integration-tests

.PHONY: test
test: ## Run unit tests locally
	go test -v ./pkg/...


.PHONY: lint
lint: ## Run linting
	golangci-lint run -v


.PHONY: mocks
mocks: clean-mocks ## For generating mock based on all project interfaces
	mockery --all --dir "./pkg" --inpackage --case underscore


.PHONY: clean-mocks
clean-mocks: ## Cleans old mocks
	find . -name "mock_*.go" -type f -print0 | xargs -0 /bin/rm -f


.PHONY: proto
proto: ## Generates protobuf files
	git submodule update --remote --init grpc-contracts && protoc --go-grpc_opt=require_unimplemented_servers=false --go_out=$(CURDIR)/grpc-contracts/proto/. --go_opt=paths=source_relative \
        --go-grpc_out=$(CURDIR)/grpc-contracts/proto/. --go-grpc_opt=paths=source_relative -I $(CURDIR)/grpc-contracts/proto/. \
        $(CURDIR)/grpc-contracts/proto/*.proto


.PHONY: docker_image_build
docker_image_build:
	@echo ">>> Building docker image"
	docker build \
		-t $(DOCKER_IMAGE) \
		-t $(DOCKER_COMPOSE_SERVICE) \
		--build-arg GIT_REPO="$(GIT_REPO)" \
		--build-arg GIT_TAG="$(GIT_TAG)" \
		--build-arg GIT_BRANCH="$(GIT_BRANCH)" \
		--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
		-f $(PATH_DOCKER_FILE_CI) \
		.


.PHONY: docker_image_inspect
docker_image_inspect:
	@echo ">>> Inspecting docker container"
	docker inspect \
		-f '{{index .ContainerConfig.Labels "repo"}}' \
		-f '{{index .ContainerConfig.Labels "tag"}}' \
		-f '{{index .ContainerConfig.Labels "branch"}}' \
		-f '{{index .ContainerConfig.Labels "commit"}}' \
		$(DOCKER_IMAGE)


.PHONY: docker_image_registry_push
docker_image_registry_push:
	@echo ">>> Tag and push docker image"
	@docker tag $(DOCKER_IMAGE) $(DOCKER_REGISTRY_ENTRY):$(IMAGE_TAG)
	@docker tag $(DOCKER_IMAGE) $(DOCKER_REGISTRY_ENTRY):$(GIT_COMMIT)
	@docker push --all-tags $(DOCKER_REGISTRY_ENTRY)


.PHONY: docker_image_deploy
docker_image_deploy:
	@echo ">>> List public hosts"
	@kubectl get ing | awk {'print $3'}
	@echo ">>> Get current running K8S-DEV \"$(K8S_DEPLOYMENT_NAME)\" deployment version"
	@CURRENT_IMAGE=$$(kubectl -n $(K8S_NAMESPACE) get deploy $(K8S_DEPLOYMENT_NAME) -ojsonpath={..image} 2>/dev/null) && \
	CURRENT_IMAGE_TAG=$$(echo $${CURRENT_IMAGE} | cut -d ':' -f2 2>/dev/null) && \
	if [ "$${CURRENT_IMAGE}" != "" ]; then \
	  echo "Current \"$(K8S_DEPLOYMENT_NAME)\" image within K8S-DEV cluster is: \"$${CURRENT_IMAGE}\""; \
		echo "New image TAG: $(IMAGE_TAG)"; \
		if [ "$${CURRENT_IMAGE_TAG}" != "$(IMAGE_TAG)" ]; then \
	  	echo ">>> Update K8S-DEV \"$(K8S_DEPLOYMENT_NAME)\" deployment image tag to \"$(IMAGE_TAG)\""; \
			kubectl -n $(K8S_NAMESPACE) set image deployment/$(K8S_DEPLOYMENT_NAME) $(K8S_DEPLOYMENT_NAME)=$(DOCKER_REGISTRY_ENTRY):$(IMAGE_TAG); \
			echo ">>> Get updated K8S-DEV \"$(K8S_DEPLOYMENT_NAME)\" deployment image"; \
	    kubectl -n $(K8S_NAMESPACE) get deploy $(K8S_DEPLOYMENT_NAME) -ojsonpath={..image}; \
	  else \
			echo ">>> Update \"$(K8S_DEPLOYMENT_NAME)\" deployment by setting \"VERSION\" environment variable, since the new image has the same tag as the running image"; \
			kubectl -n $(K8S_NAMESPACE) set env deployment/$(K8S_DEPLOYMENT_NAME) VERSION=$(GIT_COMMIT); \
			echo ">>> Get updated \"$(K8S_DEPLOYMENT_NAME)\" deployment version"; \
	    kubectl -n $(K8S_NAMESPACE) get deploy $(K8S_DEPLOYMENT_NAME) -ojsonpath={..env}; \
	  fi; \
	else \
		echo ">>> Deployment for \"$(K8S_DEPLOYMENT_NAME)\" not found"; \
	fi


.PHONY: help
help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-10s\033[0m - %s\n", $$1, $$2}'
	@ echo

