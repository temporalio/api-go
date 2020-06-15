$(VERBOSE).SILENT:

# default target
default: all-install update-proto

ifndef GOPATH
GOPATH := $(shell go env GOPATH)
endif

COLOR := "\e[1;36m%s\e[0m\n"

PROTO_ROOT := temporal-proto
PROTO_DIRS = $(sort $(dir $(shell find $(PROTO_ROOT) -name "*.proto")))
PROTO_OUT := .
PROTO_IMPORT := $(PROTO_ROOT):$(GOPATH)/src/github.com/temporalio/gogo-protobuf/protobuf

update-proto: update-proto-submodule all

all: grpc grpc-mock copyright update-dependencies gomodtidy

all-install: grpc-install mockgen-install

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

# git submodule for proto files

update-proto-submodule:
	printf $(COLOR) "Update proto-submodule..."
	git submodule update --force --remote $(PROTO_ROOT)

# Compile proto files to go

grpc: gogo-grpc fix-path

gogo-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for gogo-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogoslick_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

gogo-protobuf: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for gogo-protobuf..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogofaster_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

go-protobuf: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for go-protobuf..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

go-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for go-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

fix-path:
	mv -f $(PROTO_OUT)/temporal/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal

# Generate mocks

# All generated service files pathes relative to PROTO_OUT.
PROTO_GRPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service.pb.go"))
service_name = $(firstword $(subst /, ,$(1)))
mock_file_name = $(call service_name,$(1))mock/$(subst $(call service_name,$(1))/,,$(1:go=mock.go))

grpc-mock:
	printf $(COLOR) "Generate gRPC mocks..."
	$(foreach PROTO_GRPC_SERVICE,$(PROTO_GRPC_SERVICES),cd $(PROTO_OUT) && mockgen -package $(call service_name,$(PROTO_GRPC_SERVICE))mock -source $(PROTO_GRPC_SERVICE) -destination $(call mock_file_name,$(PROTO_GRPC_SERVICE))$(NEWLINE) )

# Plugins & tools

grpc-install: gogo-protobuf-install
	printf $(COLOR) "Installing/updating gRPC plugin..."
	GO111MODULE=off go get -u google.golang.org/grpc

gogo-protobuf-install: go-protobuf-install
	printf $(COLOR) "Installing/updating gogo protobuf plugin..."
	GO111MODULE=off go get -u github.com/temporalio/gogo-protobuf/protoc-gen-gogoslick

go-protobuf-install:
	printf $(COLOR) "Installing/updating go protobuf plugin..."
	GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go

mockgen-install:
	printf $(COLOR) "Installing/updating mockgen..."
	GO111MODULE=off go get -u github.com/golang/mock/mockgen

# Add licence header to generated files

copyright:
	printf $(COLOR) "Update license headers..."
	go run ./cmd/copyright/licensegen.go

# Keep go.mod updated

update-dependencies:
	printf $(COLOR) "Update dependencies..."
	go get -u -t ./...

gomodtidy:
	printf $(COLOR) "Run 'go mod tidy'..."
	go mod tidy

# Clean

clean:
	printf $(COLOR) "Deleting generated go files..."
# Delete all directories with *.pb.go and *.mock.go files from $(PROTO_OUT)
	$(foreach PROTO_OUT_DIR,$(shell find $(PROTO_OUT) \( -name "*.pb.go" -o -name "*.mock.go" \) -printf "%h\n" | sort -u),rm -rf $(dir $(PROTO_OUT_DIR)) )
