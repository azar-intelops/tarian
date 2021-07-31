# Image URL to use all building/pushing image targets
IMG ?= localhost:5000/tarian-cluster-agent

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif


# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

default: help


##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

build: generate fmt vet lint # TODO: include proto, install proto in ci.yaml
	CGO_ENABLED=0 go build -o ./bin/tarian-server ./cmd/tarian-server/
	CGO_ENABLED=0 go build -o ./bin/tarian-cluster-agent ./cmd/tarian-cluster-agent/
	CGO_ENABLED=0 go build -o ./bin/tarian-pod-agent ./cmd/tarian-pod-agent/
	CGO_ENABLED=0 go build -o ./bin/tarianctl ./cmd/tarianctl/

proto:
	protoc --experimental_allow_proto3_optional -I=./pkg --go_out=./pkg --go-grpc_out=./pkg --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/tarianpb/types.proto
	protoc --experimental_allow_proto3_optional -I=./pkg --go_out=./pkg --go-grpc_out=./pkg --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./pkg/tarianpb/api.proto

lint:
	revive -formatter stylish -config .revive.toml ./pkg/...

local-images: build
	docker build -f Dockerfile-server -t localhost:5000/tarian-server . && docker push localhost:5000/tarian-server
	docker build -f Dockerfile-cluster-agent -t localhost:5000/tarian-cluster-agent . && docker push localhost:5000/tarian-cluster-agent
	docker build -f Dockerfile-pod-agent -t localhost:5000/tarian-pod-agent . && docker push localhost:5000/tarian-pod-agent

unit-test:
	go test -v ./pkg/...

e2e-test:
	go test -v ./test/e2e/...

manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) webhook paths="./..."

create-kind-cluster:
	kind create cluster --config=dev/cluster-config.yaml

delete-kind-cluster:
	kind delete cluster

ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
controller-test: manifests generate fmt vet ## Run tests.
	mkdir -p ${ENVTEST_ASSETS_DIR}
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/v0.8.3/hack/setup-envtest.sh
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); go test ./... -coverprofile cover.out


##@ Deployment

deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | kubectl delete -f -


CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1)

KUSTOMIZE = $(shell pwd)/bin/kustomize
kustomize: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v3@v3.8.7)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go get $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
