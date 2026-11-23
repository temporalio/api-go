module go.temporal.io/api

go 1.25.4

require (
	github.com/golang/mock v1.6.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0
	github.com/nexus-rpc/nexus-proto-annotations v0.1.0
	github.com/nexus-rpc/sdk-go v0.6.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260526163538-3dc84a4a5aaa
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa
	google.golang.org/grpc v1.83.1
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
	golang.org/x/text v0.37.0 // indirect
)

retract (
	v1.26.2 // Contains retractions only.
	v1.26.1 // Published prematurely
)
