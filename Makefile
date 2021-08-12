SHELL = /bin/bash

PKGMAP=Mblockchain/v1alpha/common.proto=github.com/xiaobaiskill/blockchain-release-monitor/api/protos/blockchain/v1alpha
REPO_NAME ?= blockchain-release-monitor
REPO_PREFIX ?= jinmz

.PHONY: all
all: proto build

.PHONY: fmt
fmt:
	@CGO_ENABLED=0 go fmt ./...

.PHONY: vet
vet: fmt
	@CGO_ENABLED=0 go vet ./...

.PHONE: test
test: vet
	CGO_ENABLE=0 go test ./...

.PHONY: build
build: vet
	@CGO_ENABLED=0 go build  -ldflags="-s -w" -o build/app .

.PHONY: run
run:
	export LOG_MODE="server"
	./build/app server ./.config.json

.PHONY: proto-clean
proto-clean:
	rm -f api/protos/*/*/*.go api/protos/*/*/*.json

.PHONY: proto
proto: proto-clean
	docker run --rm -i -v $$PWD/api:/go/api -v $$PWD/proto.sh:/go/proto.sh ankrnetwork/ankr-grpc /bin/bash proto.sh


.PHONY: docker-build
docker-build:
	@echo "build docker image"
	@SHA1_SHORT=$(shell git rev-parse --short HEAD); \
	docker build -t $(REPO_PREFIX)/$(REPO_NAME):$$SHA1_SHORT .

.PHONY: docker-push
docker-push: docker-build
	@echo "tag & push image"
	@BRANCH_NAME=$(shell git rev-parse --abbrev-ref HEAD); SHA1_SHORT=$(shell git rev-parse --short HEAD); \
	if [[ $$BRANCH_NAME == "develop" ]]; then \
		ENV="stage"; \
	elif [[ $$BRANCH_NAME == "main" ]]; then \
		ENV="prod"; \
	else \
		ENV="feat"; \
	fi;  \
	docker tag $(REPO_PREFIX)/$(REPO_NAME):$$SHA1_SHORT  $(REPO_PREFIX)/$(REPO_NAME):$$ENV; \
	docker push $(REPO_PREFIX)/$(REPO_NAME):$$SHA1_SHORT; \
	docker push $(REPO_PREFIX)/$(REPO_NAME):$$ENV;
