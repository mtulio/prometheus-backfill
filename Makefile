
CONTAINER_CMD ?= sudo podman
BIN_DIR ?= ./bin
PROJECT ?=  github.com/mtulio/prometheus-backfill
APP ?= prometheus-backfill

VER_TAG ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER_COMMIT := $(shell git rev-parse --short HEAD)

REGISTRY ?= docker.pkg.github.com
REGISTRY_USER ?= mtulio
GH_REPO ?= $(APP)
BASE_IMAGE ?= $(REGISTRY)/$(REGISTRY_USER)/$(GH_REPO)/$(APP)
IMAGE ?= $(BASE_IMAGE):$(VER_COMMIT)

build:
	test -d $(BIN_DIR) || mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP) $(PROJECT)/cmd/$(APP)

container-build:
	$(CONTAINER_CMD) build -t $(IMAGE) -f Dockerfile .
	$(CONTAINER_CMD) tag $(IMAGE) $(BASE_IMAGE):latest

container-push:
	$(CONTAINER_CMD) push $(IMAGE)
	$(CONTAINER_CMD) push $(BASE_IMAGE):latest