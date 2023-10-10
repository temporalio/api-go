$(VERBOSE).SILENT:
all: install test
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
HELPER_FILES = $(shell find protoc-gen-go-helpers)
PROTO_FILES = $(shell find $(PROTO_ROOT) -name "*.proto" -not -path "$(PROTO_ROOT)/google/*")
PROTO_DIRS = $(sort $(dir $(PROTO_FILES)))
PROTO_ENUMS := $(shell grep -R '^enum ' $(PROTO_ROOT) | cut -d ' ' -f2)
PROTO_OUT := .
PROTO_IMPORTS = \
	-I=$(PROTO_ROOT)
PROTO_PATHS = paths=source_relative:$(PROTO_OUT)

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

##### git submodule for proto files #####
update-proto-submodule:
	printf $(COLOR) "Update proto-submodule..."
	git submodule update --init --force --remote $(PROTO_ROOT)


##### Compile proto files for go #####
grpc: go-grpc fix-path fix-enums copy-helpers

# Only install helper when its source has changed
.go-helpers-installed: $(HELPER_FILES)
	@go install ./protoc-gen-go-helpers

go-grpc: clean .go-helpers-installed $(PROTO_OUT)
	printf $(COLOR) "Compile for go-gRPC..."
	$(foreach PROTO_DIR,$(PROTO_DIRS),\
		protoc --fatal_warnings $(PROTO_IMPORTS) \
		 	--go_out=$(PROTO_PATHS) \
			--go-grpc_out=$(PROTO_PATHS) \
            --grpc-gateway_out=allow_patch_feature=false,$(PROTO_PATHS) \
			--go-helpers_out=$(PROTO_PATHS) \
			$(PROTO_DIR)*.proto;)

fix-path: go-grpc
	mv -f $(PROTO_OUT)/temporal/api/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal

copy-helpers:
	# Also copy the payload and history JSON helpers
	cp $(PROTO_OUT)/internal/temporalcommonv1/payload_json.go $(PROTO_OUT)/common/v1/

# The generated enums are go are just plain terrible, so we fix them
# by removing the typename prefixes. We already made good choices with our enum
# names, so this shouldn't be an issue
fix-enums: fix-path
	printf $(COLOR) "Fixing enum naming..."
	$(foreach PROTO_ENUM,$(PROTO_ENUMS),\
      $(shell grep -Rl "$(PROTO_ENUM)" $(PROTO_OUT) | grep -E "\.go" | xargs -P 8 sed -i "" -e "s/$(PROTO_ENUM)_\(.*\) $(PROTO_ENUM)/\1 $(PROTO_ENUM)/g;s/\.$(PROTO_ENUM)_\(.*\)/.\1/g"))

# All generated service files pathes relative to PROTO_OUT.
PROTO_GRPC_SERVICES = $(patsubst $(PROTO_OUT)/%,%,$(shell find $(PROTO_OUT) -name "service_grpc.pb.go"))
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

test: copy-helpers
	go test ./...

##### Check #####

generatorcheck:
	printf $(COLOR) "Check generated code is not stale..."
	#(cd ./cmd/proxygenerator && go mod tidy && go run ./ -verifyOnly)

check: generatorcheck

##### Clean #####
clean:
	printf $(COLOR) "Deleting generated go files..."
	# Delete all directories with *.pb.go and *.mock.go files from $(PROTO_OUT)
	find . \( -name "*.pb.go" -o -name "*.mock.go" -o -name "*.go-helpers.go" \) | xargs -I{} dirname {} | sort -u | xargs rm -rf
