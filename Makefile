APP_NAME=rest-starter

.PHONY: help build

help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

build: ## Build the container(s)
	docker build -t $(APP_NAME) -f ./build/Dockerfile .

run: stop ## Run the container
	docker run -d -p 8080:8080 --rm --name="$(APP_NAME)" $(APP_NAME)

stop: ## Stop running container(s)
	docker stop $(APP_NAME) || true

up: ## Compose the container(s)
	docker-compose -f ./deployments/docker-compose.yml up -d

down: ## Take down the composed container(s)
	docker-compose -f ./deployments/docker-compose.yml down
