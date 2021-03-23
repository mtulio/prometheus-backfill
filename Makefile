
CONTAINER_CMD ?= sudo podman
BIN ?= ./bin
PROJECT ?=  github.com/mtulio/prometheus-backfill
APP ?= prometheus-backfill

VER_TAG ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER_COMMIT := $(shell git rev-parse --short HEAD)

REGISTRY ?= docker.pkg.github.com
REGISTRY_USER ?= mtulio
GH_REPO ?= $(APP)
IMAGE ?= $(REGISTRY)/$(REGISTRY_USER)/$(GH_REPO)/$(APP):$(VER_COMMIT)

build:
	go build -o $(BIN)/$(APP) $(PROJECT)/cmd/$(APP)

container-build:
	$(CONTAINER_CMD) build -t $(IMAGE) -f Dockerfile .

container-push:
	$(CONTAINER_CMD) push $(IMAGE)