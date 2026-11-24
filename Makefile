$(VERBOSE).SILENT:
all: install test
############################# Main targets #############################
# Install everything, update submodule, and compile proto files.
install: grpc-install mockgen-install goimports-install update-proto

# Compile proto files.
proto: http-api-docs grpc goimports proxy grpc-mock

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

NEXUS_SCHEMA_ROOT := $(PROTO_ROOT)/nexus
NEXUS_PROTO_OUT := workflowservice/v1/workflowservicenexus
NEXUS_WIT_OUT := systemnexus

##### Compile proto files for go #####
grpc: http-api-docs go-grpc copy-helpers nexus-gen system-nexus nexus-annotations

.PHONY: nexus-annotations
nexus-annotations:
	go run ./cmd/gen-nexus-annotations nexusannotations/v1/options.go

nexus-gen:
	printf $(COLOR) "Generate nexus service definitions..."
	mkdir -p $(NEXUS_PROTO_OUT)
	cd $(NEXUS_SCHEMA_ROOT) && nexus-rpc-gen \
		--lang go \
		--package workflowservicenexus \
		--out-file $(CURDIR)/$(NEXUS_PROTO_OUT)/service_nexus.pb.go \
		temporal-proto-models-nexusrpc.yaml

.PHONY: system-nexus
system-nexus:
	printf $(COLOR) "Update system-nexus WIT files..."
	rm -rf $(NEXUS_WIT_OUT)
	mkdir -p $(NEXUS_WIT_OUT)/workflowservice/v1
	cp $(NEXUS_SCHEMA_ROOT)/workflow-service.wit $(NEXUS_WIT_OUT)/workflowservice/v1/
	cp -R $(NEXUS_SCHEMA_ROOT)/deps $(NEXUS_WIT_OUT)/

# Only install helpers when their source has changed
HELPER_FILES = $(shell find ./cmd/protoc-gen-go-helpers)
.go-helpers-installed: $(HELPER_FILES)
	printf $(COLOR) "Installing protoc-gen-go-helpers plugin"
	@go install ./cmd/protoc-gen-go-helpers
	@touch $@

go-grpc: clean .go-helpers-installed $(PROTO_OUT)
	printf $(COLOR) "Compile for go-gRPC..."
	# Post-processing walks every .pb.go under $(PROTO_OUT). Exclude the
	# separately generated api_next module and the helper test fixtures.
	go run ./cmd/protogen \
		--root=$(PROTO_ROOT)/temporal/api \
		-I=$(PROTO_ROOT) \
		--output=$(PROTO_OUT) \
		--exclude=internal \
		--exclude=proto/api/google \
		--exclude=$(PROTO_ROOT)/temporal/api_next \
		--exclude=$(API_NEXT_OUT) \
		--exclude=cmd/protoc-gen-go-helpers/testprotos \
		-p go-grpc_out=$(PROTO_PATHS) \
		-p grpc-gateway_out=allow_patch_feature=false,$(PROTO_PATHS) \
		-p go-helpers_out=$(PROTO_PATHS)

	# cp is safer than mv because mv cannot merge into existing directories
	cp -rf $(PROTO_OUT)/temporal/api/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal

API_NEXT_OUT := api_next
DRAFT_MANIFEST := $(PROTO_ROOT)/temporal/api/draft-manifest.json
API_NEXT_PROTO_PATHS := paths=source_relative:$(API_NEXT_OUT)

# workflowservicenexus is produced by nexus-gen, outside protogen.
API_NEXT_INPUTS := LICENSE go.mod go.sum \
	cmd/protoc-gen-go-helpers cmd/protogen \
	internal/protojson internal/strcase nexusannotations/v1/options.go \
	proxy/service_util.go proxy/system_payload.go serviceerror systemnexus temporalnexus temporalproto \
	$(NEXUS_PROTO_OUT)/service_nexus.pb.go

.PHONY: api-next
api-next: .go-helpers-installed mockgen-install nexus-annotations
	printf $(COLOR) "Compile api_next (draft) proto tree..."
	rm -rf $(API_NEXT_OUT)
	mkdir -p $(API_NEXT_OUT)/common/v1
	rsync -aR --exclude='*_test.go' --exclude=testprotos/ $(API_NEXT_INPUTS) $(API_NEXT_OUT)/
	cp internal/temporalcommonv1/payload_json.go $(API_NEXT_OUT)/common/v1/

	go run ./cmd/protogen \
		--root=$(PROTO_ROOT)/temporal/api_next \
		-I=$(PROTO_ROOT) \
		--output=$(API_NEXT_OUT) \
		-p go-grpc_out=$(API_NEXT_PROTO_PATHS) \
		-p grpc-gateway_out=allow_patch_feature=false,$(API_NEXT_PROTO_PATHS) \
		-p go-helpers_out=draft_omission=true,draft_manifest=$(DRAFT_MANIFEST),$(API_NEXT_PROTO_PATHS)
	cp -rf $(API_NEXT_OUT)/temporal/api_next/* $(API_NEXT_OUT) && rm -rf $(API_NEXT_OUT)/temporal

	go run ./cmd/protogen \
		--root=$(PROTO_ROOT)/temporal/api_next \
		-I=$(PROTO_ROOT) \
		--output=$(API_NEXT_OUT) \
		--no-rewrite-enum-const \
		--no-rewrite-enum-string \
		--no-strip-version \
		--output-descriptor=$(API_NEXT_OUT)/descriptor_set.pb
	go -C cmd/proxygenerator run ./ -apiRoot "$(abspath $(API_NEXT_OUT))"

	# The service mocks embed the concrete request/response types generated
	# above, so they have to be regenerated rather than copied from the snapshot.
	# api_next/go.mod (copied by the snapshot) already gives these the right
	# module identity, so mockgen infers the same "operatorservice"/"workflowservice"
	# import aliases it would at the module root.
	mockgen -package operatorservicemock -source $(API_NEXT_OUT)/operatorservice/v1/service_grpc.pb.go -destination $(API_NEXT_OUT)/operatorservicemock/v1/service_grpc.pb.mock.go
	mockgen -package workflowservicemock -source $(API_NEXT_OUT)/workflowservice/v1/service_grpc.pb.go -destination $(API_NEXT_OUT)/workflowservicemock/v1/service_grpc.pb.mock.go
	go run ./cmd/mockgen-fix OperatorService $(API_NEXT_OUT)/operatorservicemock/v1/service_grpc.pb.mock.go
	go run ./cmd/mockgen-fix WorkflowService $(API_NEXT_OUT)/workflowservicemock/v1/service_grpc.pb.mock.go

	goimports -w $(API_NEXT_OUT)
	(cd $(API_NEXT_OUT) && go mod tidy)

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
	go run ./cmd/mockgen-fix OperatorService operatorservicemock/v1/service_grpc.pb.mock.go
	go run ./cmd/mockgen-fix WorkflowService workflowservicemock/v1/service_grpc.pb.mock.go

.PHONY: proxy
proxy: gen-proto-desc
	printf $(COLOR) "Generate proxy code..."
	(cd ./cmd/proxygenerator && go mod tidy && go run ./)

goimports:
	printf $(COLOR) "Run goimports..."
	goimports -w $(PROTO_OUT)

gen-proto-desc:
	printf $(COLOR) "Generating proto descriptors..."
	go run ./cmd/protogen \
		--root=$(PROTO_ROOT)/temporal/api \
		-I=$(PROTO_ROOT) \
		--output=$(PROTO_OUT) \
		--exclude=internal \
		--exclude=proto/api/google \
		--exclude=$(PROTO_ROOT)/temporal/api_next \
		--no-rewrite-enum-const \
		--no-rewrite-enum-string \
		--no-strip-version \
		--output-descriptor=$(PROTO_OUT)/descriptor_set.pb

##### Plugins & tools #####
.PHONY: api-next-tools
api-next-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.2
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.30.0
	go install "golang.org/x/tools/cmd/goimports@$$(go list -modfile=build/go.mod -m -f '{{.Version}}' golang.org/x/tools)"

grpc-install: nexus-rpc-gen-install
	@printf $(COLOR) "Install/update grpc and plugins..."
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

NEXUS_RPC_GEN_CACHE := $(HOME)/.cache/nexus-rpc-gen

.PHONY: nexus-rpc-gen-install
nexus-rpc-gen-install:
	@printf $(COLOR) "Install nexus-rpc-gen from main..."
	@rm -rf $(NEXUS_RPC_GEN_CACHE)
	@git clone --depth 1 https://github.com/nexus-rpc/nexus-rpc-gen.git $(NEXUS_RPC_GEN_CACHE)
	@cd $(NEXUS_RPC_GEN_CACHE)/src && pnpm install --ignore-scripts && pnpm run build
	@chmod +x $(NEXUS_RPC_GEN_CACHE)/src/packages/nexus-rpc-gen/dist/index.js
	@cd $(NEXUS_RPC_GEN_CACHE)/src/packages/nexus-rpc-gen && pnpm add --global .

mockgen-install:
	printf $(COLOR) "Install/update mockgen..."
	go install -modfile=build/go.mod github.com/golang/mock/mockgen

goimports-install:
	printf $(COLOR) "Install/update goimports..."
	go install golang.org/x/tools/cmd/goimports@latest

##### go.mod #####
update-dependencies:
	printf $(COLOR) "Update go dependencies..."
	go get -u -t $(PINNED_DEPENDENCIES) ./...

gomodtidy:
	printf $(COLOR) "go mod tidy..."
	go mod tidy

##### Test #####

.PHONY: draft-inputs
draft-inputs:
	@if [ -d "$(PROTO_ROOT)/temporal/api_next" ] && \
	    [ -f "$(PROTO_ROOT)/temporal/api_next/protometa/v1/draft.proto" ]; then \
		echo true; \
	else \
		echo false; \
	fi

.PHONY: helper-test-protos
helper-test-protos:
	@set -eu; tools_dir=$$(mktemp -d); trap 'rm -rf "$$tools_dir"' EXIT; \
	GOBIN="$$tools_dir" go install google.golang.org/protobuf/cmd/protoc-gen-go; \
	go build -o "$$tools_dir/protoc-gen-go-helpers" ./cmd/protoc-gen-go-helpers; \
	find cmd/protoc-gen-go-helpers/testprotos/draft -name '*.draft.pb.go' -delete; \
	PATH="$$tools_dir:$(PATH)" go run ./cmd/protogen \
		--root=cmd/protoc-gen-go-helpers/testprotos/draft \
		--output=cmd/protoc-gen-go-helpers/testprotos/draft \
		--no-rewrite-enum-const \
		--no-rewrite-enum-string \
		-p go-helpers_out=paths=source_relative,draft_omission=true,draft_manifest=cmd/protoc-gen-go-helpers/testprotos/draft/draft-manifest.json:cmd/protoc-gen-go-helpers/testprotos/draft

# We need to ensure protos are up to date to test our UTF-8 post-processing
.PHONY: test test-stable
test: copy-helpers
	go test -race -tags protolegacy ./...
	go -C cmd/proxygenerator test ./...
	if [ -f internal/draftcompat/go.mod ]; then go -C internal/draftcompat test ./...; fi

test-stable: copy-helpers
	go test -race -tags protolegacy ./...
	go -C cmd/proxygenerator test -skip '^TestAPINextProxy$$' ./...

##### Check #####

generatorcheck: gen-proto-desc
	printf $(COLOR) "Check generated code is not stale..."
	(cd ./cmd/proxygenerator && go mod tidy && go run ./ -verifyOnly)

check: generatorcheck

##### Clean #####
clean:
	printf $(COLOR) "Deleting generated go files..."
	# Delete all directories with *.pb.go and *.mock.go files from $(PROTO_OUT)
	find $(PROTO_OUT) \( -name "*.pb.go" -o -name "*.mock.go" -o -name "*.go-helpers.go" \) | xargs -I{} dirname {} | egrep -v 'testprotos|api_next' | sort -u | xargs rm -rf
