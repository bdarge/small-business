THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: help build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dev_build: ## to build images with docker-compose in development. Pass ARM=false for none ARM os
ifeq ($(ARM), false)
	docker compose -f docker-compose.yml build $(c)
else
	docker compose -f docker-compose.yml -f docker-compose-arm.yml build $(c)
endif

build: ## build images with docker-compose. Pass ARM=false for none ARM os
ifeq ($(ARM), false)
	docker compose -f docker-compose.yml -f docker-compose.prod.yml build $(c)
else
	docker compose -f docker-compose.yml -f docker-compose-arm.yml -f docker-compose.prod.yml build $(c)
endif

dev_up: ## run containers in development. Pass ARM=false for none ARM os
ifeq ($(ARM), false)
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
else
	docker compose -f docker-compose.yml -f docker-compose-arm.yml -f docker-compose.prod.yml up -d
endif

up: ## run containers. Pass ARM=false for none ARM os
ifeq ($(ARM), false)
	docker compose -f docker-compose.yml up -d
else
	docker compose -f docker-compose.yml -f docker-compose-arm.yml up -d
endif

down: ## remove running containers (docker compose)
	docker compose down

destroy: ## destroy running containers and volumes (docker compose)
	docker compose down -v
