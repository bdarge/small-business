THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: help build

TAG = v0.0.1
ARCH = true

DB_NAME       := bdarge/sb-db
DB_IMG        := ${DB_NAME}:$(TAG)
DB_LATEST     := nfs.my.home:5000/${DB_NAME}:latest

UI_NAME       := bdarge/sb-ui
UI_IMG        := ${UI_NAME}:${TAG}
UI_LATEST     := nfs.my.home:5000/${UI_NAME}:latest
API_BASE_URL   := http://sb-info.my.home/api/v1

API_NAME       := bdarge/sb-api
API_IMG        := ${API_NAME}:${TAG}
API_LATEST     := nfs.my.home:5000/${API_NAME}:latest

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dev_dc_build: ## to build containers in development. Pass ARM=false for none ARM os
ifeq ($(ARM), false)
	docker compose -f docker-compose.yml build $(c)
else
	docker compose -f docker-compose.yml -f docker-compose-arm.yml build $(c)
endif

dc_build: ## build containers. Pass ARM=false for none ARM os
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

build: k_ui_arm k_api_arm k_db_arm ## build images for local k8s

k_ui_arm: ## build ui image for for local k8s
	cd ui; docker buildx b -f docker/Dockerfile-arm64 --platform linux/arm64 --target prod \
    --load --build-arg API_BASE_URL=${API_BASE_URL} --build-arg NODE_ENV=PROD -t ${UI_IMG} .; cd ..
	docker image tag ${UI_IMG} ${UI_LATEST}

api_doc: ## create api doc
	cd api/cmd/sm; swag init cd -

k_api_arm: api_doc ## build api image for for local k8s
	cd api; docker buildx b --platform linux/arm64 --target prod \
	--build-arg NODE_ENV=PROD --load -t ${API_IMG} .; cd ..
	docker image tag ${API_IMG} ${API_LATEST}

k_db_arm: ## build db image for for local k8s
	cd db; docker buildx b -f Dockerfile-mariadb --platform linux/arm64 -t ${DB_IMG} .; cd ..
	docker image tag ${DB_IMG} ${DB_LATEST}

push: ## push images to local registry
	docker push ${UI_LATEST}
	docker push ${DB_LATEST}
	docker push ${API_LATEST}

