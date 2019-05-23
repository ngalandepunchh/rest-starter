APP_NAME=starter

.PHONY: help build

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build the container
	docker build -t $(APP_NAME) -f ./build/Dockerfile .

run: ## Run the container
	docker run -i -t --rm --name="$(APP_NAME)" $(APP_NAME)
