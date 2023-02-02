$(VERBOSE).SILENT:
############################# Main targets #############################
# Install everything, update submodule, and compile proto files.
install: grpc-install mockgen-install goimports-install update-proto

# Compile proto files.
proto: grpc goimports proxy grpc-mock copyright

# Update submodule and compile proto files.
update-proto: update-proto-submodule proto update-dependencies gomodtidy
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
PROTO_FILES = $(shell find $(PROTO_ROOT) -name "*.proto")
PROTO_DIRS = $(sort $(dir $(PROTO_FILES)))
PROTO_OUT := .
PROTO_IMPORTS = -I=$(PROTO_ROOT) -I=$(shell go list -modfile build/go.mod -m -f '{{.Dir}}' github.com/temporalio/gogo-protobuf)/protobuf

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

##### git submodule for proto files #####
update-proto-submodule:
	printf $(COLOR) "Update proto-submodule..."
	git submodule update --init --force --remote $(PROTO_ROOT)

##### Compile proto files for go #####
grpc: gogo-grpc fix-path

go-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for go-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --fatal_warnings $(PROTO_IMPORTS) --go_out=plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

gogo-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for gogo-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --fatal_warnings $(PROTO_IMPORTS) --gogoslick_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

fix-path:
	mv -f $(PROTO_OUT)/temporal/api/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal

# All generated service files pathes relative to PROTO_OUT.
PROTO_GRPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service.pb.go"))
service_name = $(firstword $(subst /, ,$(1)))
mock_file_name = $(call service_name,$(1))mock/$(subst $(call service_name,$(1))/,,$(1:go=mock.go))

grpc-mock:
	printf $(COLOR) "Generate gRPC mocks..."
	$(foreach PROTO_GRPC_SERVICE,$(PROTO_GRPC_SERVICES),cd $(PROTO_OUT) && mockgen -package $(call service_name,$(PROTO_GRPC_SERVICE))mock -source $(PROTO_GRPC_SERVICE) -destination $(call mock_file_name,$(PROTO_GRPC_SERVICE))$(NEWLINE) )

.PHONY: proxy
proxy:
	printf $(COLOR) "Generate proxy code..."
	(cd ./cmd/proxygenerator && go mod tidy && go run ./)

goimports:
	printf $(COLOR) "Run goimports..."
	goimports -w $(PROTO_OUT)

##### Plugins & tools #####
grpc-install: gogo-protobuf-install
	printf $(COLOR) "Install/update gRPC plugins..."
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gogo-protobuf-install: go-protobuf-install
	go install github.com/temporalio/gogo-protobuf/protoc-gen-gogoslick@latest
	go install -modfile build/go.mod github.com/temporalio/gogo-protobuf/protoc-gen-gogoslick

go-protobuf-install:
	go install github.com/golang/protobuf/protoc-gen-go@latest

mockgen-install:
	printf $(COLOR) "Install/update mockgen..."
	go install github.com/golang/mock/mockgen@v1.6.0

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

test:
	go test ./proxy ./serviceerror

##### Check #####

generatorcheck:
	printf $(COLOR) "Check generated code is not stale..."
	(cd ./cmd/proxygenerator && go mod tidy && go run ./ -verifyOnly)

check: generatorcheck

##### Clean #####
clean:
	printf $(COLOR) "Deleting generated go files..."
	# Delete all directories with *.pb.go and *.mock.go files from $(PROTO_OUT)
	find $(PROTO_OUT) \( -name "*.pb.go" -o -name "*.mock.go" \) | xargs -I{} dirname {} | sort -u | xargs rm -rf