$(VERBOSE).SILENT:

# default target
default: all-install update-proto

ifndef GOPATH
GOPATH := $(shell go env GOPATH)
endif

PROTO_ROOT := temporal-proto
# List only subdirectories with *.proto files. Sort to remove duplicates.
PROTO_DIRS = $(sort $(dir $(wildcard $(PROTO_ROOT)/*/*.proto)))
PROTO_SERVICES = $(wildcard $(PROTO_ROOT)/*/service.proto)
PROTO_OUT := .
PROTO_IMPORT := $(PROTO_ROOT):$(GOPATH)/src/github.com/gogo/protobuf/protobuf

update-proto: update-proto-submodule all

all: grpc grpc-mock copyright gomodtidy

all-install: grpc-install mockgen-install

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

# git submodule for proto files

update-proto-submodule:
#	git submodule update --init --remote $(PROTO_ROOT)

# Compile proto files to go

grpc: gogo-grpc

gogo-grpc: clean $(PROTO_OUT)
	echo "Compiling for gogo-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogofaster_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

gogo-protobuf: clean $(PROTO_OUT)
	echo "Compiling for gogo-protobuf..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogofaster_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

go-protobuf: clean $(PROTO_OUT)
	echo "Compiling for go-protobuf..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

go-grpc: clean $(PROTO_OUT)
	echo "Compiling for go-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

# Generate mocks

# All generated service files pathes relative to PROTO_OUT
PROTO_GRPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service.pb.go"))
dir_no_slash = $(patsubst %/,%,$(dir $(1)))
dirname = $(notdir $(call dir_no_slash,$(1)))

grpc-mock:
	@echo "Generate gRPC mocks..."
	@$(foreach PROTO_GRPC_SERVICE,$(PROTO_GRPC_SERVICES),cd $(PROTO_OUT) && mockgen -package $(call dirname,$(PROTO_GRPC_SERVICE))mock -source $(PROTO_GRPC_SERVICE) -destination $(call dir_no_slash,$(PROTO_GRPC_SERVICE))mock/$(notdir $(PROTO_GRPC_SERVICE:go=mock.go)) )

# Plugins & tools

grpc-install: gogo-protobuf-install
	echo "Installing/updaing gRPC plugins..."
	GO111MODULE=off go get -u google.golang.org/grpc

gogo-protobuf-install: go-protobuf-install
	GO111MODULE=off go get -u github.com/gogo/protobuf/protoc-gen-gogofaster

go-protobuf-install:
	GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go

mockgen-install:
	GO111MODULE=off go get -u github.com/golang/mock/mockgen

# Add licence header to generated files

copyright:
	go run ./cmd/copyright/licensegen.go

# Keep go.mod updated

gomodtidy:
	go mod tidy

# clean

clean:
	echo "Deleting generated go files..."
	rm -rf $(PROTO_OUT)/*/*.pb.*go
