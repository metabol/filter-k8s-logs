BINARY_NAME = filter

DOCKER_IMG ?= somebody/filter-kubernetes-logs
TAG ?= edge

OUTPUT_DIR = bin

.PHONY: build
build:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME)

.PHONY: linux
linux:
	$(MAKE) filter-linux

.PHONY: filter-linux
filter-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)

.PHONY: docker-build
docker-build: linux
docker-build:
	docker build -f Dockerfile -t $(DOCKER_IMG):$(TAG) .

.PHONY: test
test:
	go test ./...

HAS_GLIDE := $(shell command -v glide;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
	glide install