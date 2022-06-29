# Variable
IMAGE_REGISTRY := aanogueira
PROJECT_VERSION := 0.0.1

# Static
PROJECT_NAME := prefect-exporter

.PHONY: help
help: ## Display all available options
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Options

.PHONY: build
build: ## Builds image
	@docker build -t $(IMAGE_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERSION) -f Dockerfile .

.PHONY: push
push: ## Pushes image to registry
	@docker push $(IMAGE_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERSION)

##@ Automation

.PHONY: all
all: build push ## Builds/pushes image to registry
