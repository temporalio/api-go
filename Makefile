$(VERBOSE).SILENT:
############################# Main targets #############################
# Install everything, update submodule, run all linters, and compile proto files.
install: grpc-install mockgen-install api-linter-install buf-install update-proto

# Run all linters and compile proto files.
proto: grpc grpc-mock copyright

# Update submodule, run all linters, and compile proto files.
update-proto: update-proto-submodule proto update-dependencies gomodtidy
########################################################################

##### Variables ######
ifndef GOPATH
GOPATH := $(shell go env GOPATH)
endif

GOBIN := $(if $(shell go env GOBIN),$(shell go env GOBIN),$(GOPATH)/bin)
export PATH := $(GOBIN):$(PATH)

COLOR := "\e[1;36m%s\e[0m\n"

PROTO_ROOT := proto/api
PROTO_FILES = $(shell find $(PROTO_ROOT) -name "*.proto")
PROTO_DIRS = $(sort $(dir $(PROTO_FILES)))
PROTO_OUT := .
PROTO_IMPORT := $(PROTO_ROOT):$(GOPATH)/src/github.com/temporalio/gogo-protobuf/protobuf

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

# git submodule for proto files

update-proto-submodule:
	printf $(COLOR) "Update proto-submodule..."
	git submodule update --force --remote $(PROTO_ROOT)

##### Compile proto files for go #####
grpc: buf api-linter gogo-grpc fix-path

go-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for go-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

gogo-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for gogo-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogoslick_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

fix-path:
	mv -f $(PROTO_OUT)/temporal/api/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal

# All generated service files pathes relative to PROTO_OUT.
PROTO_GRPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service.pb.go"))
service_name = $(firstword $(subst /, ,$(1)))
mock_file_name = $(call service_name,$(1))mock/$(subst $(call service_name,$(1))/,,$(1:go=mock.go))

grpc-mock:
	printf $(COLOR) "Generate gRPC mocks..."
	$(foreach PROTO_GRPC_SERVICE,$(PROTO_GRPC_SERVICES),cd $(PROTO_OUT) && mockgen -package $(call service_name,$(PROTO_GRPC_SERVICE))mock -source $(PROTO_GRPC_SERVICE) -destination $(call mock_file_name,$(PROTO_GRPC_SERVICE))$(NEWLINE) )

##### Plugins & tools #####
grpc-install: gogo-protobuf-install
	printf $(COLOR) "Installing/updating gRPC plugins..."
	GO111MODULE=off go get -u google.golang.org/grpc

gogo-protobuf-install: go-protobuf-install
	GO111MODULE=off go get -u github.com/temporalio/gogo-protobuf/protoc-gen-gogoslick

go-protobuf-install:
	GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go

mockgen-install:
	printf $(COLOR) "Installing/updating mockgen..."
	GO111MODULE=off go get -u github.com/golang/mock/mockgen

api-linter-install:
	printf $(COLOR) "Installing/updating api-linter..."
	GO111MODULE=off go get -u github.com/googleapis/api-linter/cmd/api-linter

buf-install:
	printf $(COLOR) "Installing/updating buf..."
	GO111MODULE=off go get -u github.com/bufbuild/buf/cmd/buf

##### License header #####
copyright:
	printf $(COLOR) "Update license headers..."
	go run ./cmd/copyright/licensegen.go

##### go.mod #####
update-dependencies:
	printf $(COLOR) "Update go dependencies..."
	go get -u -t ./...

gomodtidy:
	printf $(COLOR) "go mod tidy..."
	go mod tidy

##### Linters #####
api-linter:
	printf $(COLOR) "Running api-linter..."
	api-linter --set-exit-status --output-format summary --config api-linter.yaml $(PROTO_FILES)

buf:
	printf $(COLOR) "Running buf linter..."
	buf check lint

##### Clean #####
clean:
	printf $(COLOR) "Deleting generated go files..."
# Delete all directories with *.pb.go and *.mock.go files from $(PROTO_OUT)
	$(foreach PROTO_OUT_DIR,$(shell find $(PROTO_OUT) \( -name "*.pb.go" -o -name "*.mock.go" \) -printf "%h\n" | sort -u),rm -rf $(dir $(PROTO_OUT_DIR)) )
