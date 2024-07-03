module go.temporal.io/api

go 1.21

toolchain go1.22.4

require (
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/stretchr/testify v1.9.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v1.26.2 // Contains retractions only.
	v1.26.1 // Published prematurely
)
