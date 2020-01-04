.PHONY: grpc yarpc clean yarpc-install grpc-install
$(VERBOSE).SILENT:

# default target
default: all

# List only subdirectories with *.proto files.
# sort to remove duplicates.
PROTO_ROOT := temporal-proto
PROTO_DIRS = $(sort $(dir $(wildcard $(PROTO_ROOT)/*/*.proto)))
PROTO_SERVICES = $(wildcard $(PROTO_ROOT)/*/service.proto)
PROTO_OUT := .
PROTO_IMPORT := $(PROTO_ROOT)

all: update-proto-submodule yarpc grpc grpc-mock yarpc-mock copyright gomodtidy

all-install: grpc-install yarpc-install mockgen-install

# git submodule for proto files

update-proto-submodule:
	git submodule update --init --remote $(PROTO_ROOT)

# Compile proto files to go

yarpc: gogo-protobuf
	echo "Compiling for YARPC..."
	$(foreach PROTO_SERVICE,$(PROTO_SERVICES),protoc --proto_path=$(PROTO_IMPORT) --yarpc-go_out=$(PROTO_OUT) $(PROTO_SERVICE);)

grpc: gogo-grpc

gogo-grpc: clean $(PROTO_OUT)
	echo "Compiling for gogo-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogoslick_out=plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

gogo-protobuf: clean $(PROTO_OUT)
	echo "Compiling for gogo-protobuf..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --gogoslick_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

go-protobuf: clean $(PROTO_OUT)
	echo "Compiling for go-protobuf..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

go-grpc: clean $(PROTO_OUT)
	echo "Compiling for go-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),protoc --proto_path=$(PROTO_IMPORT) --go_out=plugins=grpc,paths=source_relative:$(PROTO_OUT) $(PROTO_DIR)*.proto;)

# Generate mocks

# All generated service files pathes relative to PROTO_OUT
PROTO_GRPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service.pb.go"))
PROTO_YARPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service.pb.yarpc.go"))
dir_no_slash = $(patsubst %/,%,$(dir $(1)))
dirname = $(notdir $(call dir_no_slash,$(1)))

grpc-mock: gobin-install
	@echo "Generate gRPC mocks..."
	@$(foreach PROTO_GRPC_SERVICE,$(PROTO_GRPC_SERVICES),cd $(PROTO_OUT) && mockgen -package $(call dirname,$(PROTO_GRPC_SERVICE))mock -source $(PROTO_GRPC_SERVICE) -destination $(call dir_no_slash,$(PROTO_GRPC_SERVICE))mock/$(notdir $(PROTO_GRPC_SERVICE:go=mock.go)) )

yarpc-mock: gobin-install
	@echo "Generate YARPC mocks..."
	@$(foreach PROTO_YARPC_SERVICE,$(PROTO_YARPC_SERVICES),cd $(PROTO_OUT) && mockgen -package $(call dirname,$(PROTO_YARPC_SERVICE))mock -source $(PROTO_YARPC_SERVICE) -destination $(call dir_no_slash,$(PROTO_YARPC_SERVICE))mock/$(notdir $(PROTO_YARPC_SERVICE:go=mock.go)) )

# Plugins & tools

yarpc-install: gogo-protobuf-install
	echo "Installing/updaing YARPC plugins..."
	go get -u go.uber.org/yarpc/encoding/protobuf/protoc-gen-yarpc-go

grpc-install: gogo-protobuf-install
	echo "Installing/updaing gRPC plugins..."
	go get -u google.golang.org/grpc

gogo-protobuf-install: go-protobuf-install
	go get -u github.com/gogo/protobuf/protoc-gen-gogoslick

go-protobuf-install:
	go get -u github.com/golang/protobuf/protoc-gen-go

gobin-install:
	GO111MODULE=off go get -u github.com/myitcv/gobin

mockgen-install: gobin-install
	gobin -mod=readonly github.com/golang/mock/mockgen

# Add licence header to generated files

copyright:
	go run ./cmd/copyright/licensegen.go

# Keep go.mod updated

gomodtidy:
	go mod tidy

# clean

clean:
	echo "Deleting generated go files..."
	rm -rf $(PROTO_OUT)/*/*.go