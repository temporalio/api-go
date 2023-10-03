$(VERBOSE).SILENT:
############################# Main targets #############################
# Install everything, update submodule, and compile proto files.
install: mockgen-install goimports-install update-proto

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
PROTO_ENUMS := $(shell grep -R '^enum ' $(PROTO_ROOT) | cut -d ' ' -f2)
PROTO_OUT := .

$(PROTO_OUT):
	mkdir $(PROTO_OUT)

##### git submodule for proto files #####
update-proto-submodule:
	printf $(COLOR) "Update proto-submodule..."
	git submodule update --init --force --remote $(PROTO_ROOT)


##### Compile proto files for go #####
grpc: go-grpc fix-path fix-enums

go-grpc: clean $(PROTO_OUT)
	printf $(COLOR) "Compiling for go-gRPC..."
	# Relative to PROTO_ROOT
	PROTO_OUT=../.. make -C $(PROTO_ROOT) go-grpc

fix-path: go-grpc
	mv -f $(PROTO_OUT)/temporal/api/* $(PROTO_OUT) && rm -rf $(PROTO_OUT)/temporal
	# Also copy the payload and history JSON helpers
	cp $(PROTO_OUT)/internal/temporalcommonv1/payload_json.go $(PROTO_OUT)/common/v1/
	cp $(PROTO_OUT)/internal/temporalhistoryv1/load*.go $(PROTO_OUT)/history/v1/

# The generated enums are go are just plain terrible, so we fix them
# by removing the typename prefixes. We already made good choices with our enum
# names, so this shouldn't be an issue
fix-enums: fix-path
	printf $(COLOR) "Fixing enum naming..."
    # NOTE: we can speed this up by doing this in parallel or using multiple sed expressions
	$(foreach PROTO_ENUM,$(PROTO_ENUMS),\
      $(shell grep -Rl "$(PROTO_ENUM)" | xargs sed -i "" -e "s/$(PROTO_ENUM)_\(.*\) $(PROTO_ENUM)/\1 $(PROTO_ENUM)/g"))

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

test:
	go test ./serviceerror/... ./proxy/... ./history/...

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
