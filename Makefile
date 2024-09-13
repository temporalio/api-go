$(VERBOSE).SILENT:
all: install test
############################# Main targets #############################
# Install everything, update submodule, and compile proto files.
install: grpc-install mockgen-install goimports-install update-proto

# Compile proto files.
proto: http-api-docs grpc goimports proxy grpc-mock copyright

# Update submodule and compile proto files.
update-proto: update-proto-submodule proto gomodtidy
########################################################################

##### Variables ######
ifndef GOPATH
GOPATH := $(shell go env GOPATH)
endif

GOBIN := $(if $(shell go env GOBIN),$(shell go env GOBIN),$(GOPATH)/bin)
PATH := $(GOBIN):$(PATH)

COLOR := "\e[1;36m%s\e[0m\n"

PINNED_DEPENDENCIES := \

PROTO_ROOT := proto/api
PROTO_CLOUD_ROOT := proto/api-cloud
PROTO_OUT := .
PROTO_IMPORTS = \
	-I=$(PROTO_ROOT)
PROTO_PATHS = paths=source_relative:$(PROTO_OUT)

OAPI_ROOT := $(PROTO_ROOT)/openapi
OAPI_OUT := temporalproto/openapi

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

##### git submodule for proto files #####
update-proto-submodule:
	printf $(COLOR) "Update proto-submodule..."
	git -c protocol.file.allow=always submodule update --init --force --remote $(PROTO_ROOT)

##### Compile proto files for go #####
grpc: http-api-docs go-grpc copy-helpers

# Only install helper when its source has changed
HELPER_FILES = $(shell find ./cmd/protoc-gen-go-helpers)
.go-helpers-installed: $(HELPER_FILES)
	printf $(COLOR) "Installing protoc plugin"
	@go install ./cmd/protoc-gen-go-helpers
	@touch $@

go-grpc: clean .go-helpers-installed $(PROTO_OUT)
	printf $(COLOR) "Compile for go-gRPC..."
	go run ./cmd/protogen \
		--root=$(PROTO_ROOT) \
		--root=$(PROTO_CLOUD_ROOT) \
		--output=$(PROTO_OUT) \
		--exclude=internal \
		--exclude=proto/api/google \
		-p go-grpc_out=$(PROTO_PATHS) \
		-p grpc-gateway_out=allow_patch_feature=false,$(PROTO_PATHS) \
		-p go-helpers_out=$(PROTO_PATHS)

	mv -f $(PROTO_OUT)/temporal/api/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal

http-api-docs: go-grpc
	go run cmd/encode-openapi-spec/main.go \
		-v2=$(OAPI_ROOT)/openapiv2.json \
		-v2-out=$(OAPI_OUT)/openapiv2.go \
		-v3=$(OAPI_ROOT)/openapiv3.yaml \
		-v3-out=$(OAPI_OUT)/openapiv3.go

# Copy the payload helpers
copy-helpers:
	chmod +w $(PROTO_OUT)/common/v1/payload_json.go 2>/dev/null || true
	cp $(PROTO_OUT)/internal/temporalcommonv1/payload_json.go $(PROTO_OUT)/common/v1/
	chmod -w $(PROTO_OUT)/common/v1/payload_json.go

grpc-mock:
	printf $(COLOR) "Generate gRPC mocks..."
	mockgen -package operatorservicemock -source operatorservice/v1/service_grpc.pb.go -destination operatorservicemock/v1/service_grpc.pb.mock.go
	mockgen -package workflowservicemock -source workflowservice/v1/service_grpc.pb.go -destination workflowservicemock/v1/service_grpc.pb.mock.go
	mockgen -package cloudservicemock -source cloud/cloudservice/v1/service_grpc.pb.go -destination cloud/cloudservicemock/v1/service_grpc.pb.mock.go
	go run ./cmd/mockgen-fix OperatorService operatorservicemock/v1/service_grpc.pb.mock.go
	go run ./cmd/mockgen-fix WorkflowService workflowservicemock/v1/service_grpc.pb.mock.go
	go run ./cmd/mockgen-fix CloudService cloud/cloudservicemock/v1/service_grpc.pb.mock.go

.PHONY: proxy
proxy:
	printf $(COLOR) "Generate proxy code..."
	(cd ./cmd/proxygenerator && go mod tidy && go run ./)

goimports:
	printf $(COLOR) "Run goimports..."
	goimports -w $(PROTO_OUT)

##### Plugins & tools #####
grpc-install:
	@printf $(COLOR) "Install/update grpc and plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

mockgen-install:
	printf $(COLOR) "Install/update mockgen..."
	go install -modfile=build/go.mod github.com/golang/mock/mockgen

goimports-install:
	printf $(COLOR) "Install/update goimports..."
	go install golang.org/x/tools/cmd/goimports@latest

##### License header #####
copyright:
	printf $(COLOR) "Update license headers..."
	go run ./cmd/copyright/licensegen.go

##### go.mod #####
update-dependencies:
	printf $(COLOR) "Update go dependencies..."
	go get -u -t $(PINNED_DEPENDENCIES) ./...

gomodtidy:
	printf $(COLOR) "go mod tidy..."
	go mod tidy

##### Test #####

# We need to ensure protos are up to date to test our UTF-8 post-processing
test: copy-helpers
	go test -tags protolegacy ./...

##### Check #####

generatorcheck:
	printf $(COLOR) "Check generated code is not stale..."
	#(cd ./cmd/proxygenerator && go mod tidy && go run ./ -verifyOnly)

check: generatorcheck

##### Clean #####
clean:
	printf $(COLOR) "Deleting generated go files..."
	# Delete all directories with *.pb.go and *.mock.go files from $(PROTO_OUT)
	find $(PROTO_OUT) \( -name "*.pb.go" -o -name "*.mock.go" -o -name "*.go-helpers.go" \) | xargs -I{} dirname {} | egrep -v 'testprotos' | sort -u | xargs rm -rf
	# Delete entire cloud dir
	rm -rf cloud
