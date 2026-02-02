# kiosk-webservice - Build & Push ke Docker Hub
# Usage:
#   make build                    # build image
#   make push                     # build lalu push (set DOCKER_USER)
#   make push DOCKER_USER=myuser   # push ke docker.io/myuser/kiosk-webservice:latest

DOCKER_USER   ?= $(shell echo $$DOCKERHUB_USERNAME)
IMAGE_NAME    := kiosk-webservice
TAG           ?= latest
IMAGE         := $(DOCKER_USER)/$(IMAGE_NAME):$(TAG)
DOCKERFILE    := Dockerfile

.PHONY: build push run test help

default: help

help:
	@echo "Targets:"
	@echo "  make build              Build image: $(IMAGE_NAME):$(TAG)"
	@echo "  make push               Build lalu push ke Docker Hub"
	@echo "  make run                 Jalankan container lokal (port 8099)"
	@echo "  make test                Jalankan go test di container"
	@echo ""
	@echo "Variabel:"
	@echo "  DOCKER_USER  = $(DOCKER_USER)  (wajib untuk push; atau set DOCKERHUB_USERNAME)"
	@echo "  IMAGE_NAME   = $(IMAGE_NAME)"
	@echo "  TAG         = $(TAG)"
	@echo "  IMAGE       = $(IMAGE)"
	@echo ""
	@echo "Contoh push ke Docker Hub:"
	@echo "  docker login"
	@echo "  make push DOCKER_USER=your-username"
	@echo "  make push DOCKER_USER=your-username TAG=v1.0.0"

build:
	docker build -t $(IMAGE_NAME):$(TAG) -f $(DOCKERFILE) .

push: check-user build
	docker tag $(IMAGE_NAME):$(TAG) $(IMAGE)
	docker push $(IMAGE)
	@echo "Pushed: $(IMAGE)"

check-user:
	@if [ -z "$(DOCKER_USER)" ]; then \
		echo "Error: DOCKER_USER kosong. Set env atau: make push DOCKER_USER=your-dockerhub-username"; \
		exit 1; \
	fi

run:
	docker run --rm -p 8099:8099 $(IMAGE_NAME):$(TAG)

test:
	docker build --target run-test-stage -t $(IMAGE_NAME)-test -f $(DOCKERFILE) .
