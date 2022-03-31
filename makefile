THIS_FILE := $(lastword $(MAKEFILE_LIST))
.PHONY: help build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

GIT_SHA1 		    = $(shell git rev-parse --verify HEAD)
IMAGES_TAG 		    = ${shell git describe --exact-match --tags 2> /dev/null || echo 'latest'}
IMAGE_PREFIX 		= sb-

IMAGE_DIRS = $(wildcard db api ui)

# All targets are `.PHONY` ie always need to be rebuilt
.PHONY: build ${IMAGE_DIRS}

# Build all images
build: ${IMAGE_DIRS} ## build images and push to gitlab container registry

# Build and tag a single image
${IMAGE_DIRS}:
	$(eval IMAGE_NAME := $@)

ifeq ($@, 'ui')
	docker buildx b -t ${DOCKERHUB_OWNER}/${IMAGE_PREFIX}${IMAGE_NAME}:${IMAGES_TAG} \
	-t ${DOCKERHUB_OWNER}/${IMAGE_PREFIX}${IMAGE_NAME}:latest \
	--platform linux/arm64 --target prod_arm \
    --load --build-arg API_BASE_URL=${API_BASE_URL} --build-arg NODE_ENV=${NODE_ENV} \
    --build-arg TAG=${IMAGE_PREFIX}${IMAGE_NAME} --build-arg GIT_SHA1=${GIT_SHA1} $@
else
	docker buildx b -t ${DOCKERHUB_OWNER}/${IMAGE_PREFIX}${IMAGE_NAME}:${IMAGES_TAG} \
	-t ${DOCKERHUB_OWNER}/${IMAGE_PREFIX}${IMAGE_NAME}:latest \
	--platform linux/arm64 --target prod \
    --load --build-arg TAG=${IMAGE_PREFIX}${IMAGE_NAME} --build-arg GIT_SHA1=${GIT_SHA1} $@
endif
	docker push ${DOCKERHUB_OWNER}/${IMAGE_PREFIX}${IMAGE_NAME}:${IMAGES_TAG}
    docker push ${DOCKERHUB_OWNER}/${IMAGE_PREFIX}${IMAGE_NAME}:latest
