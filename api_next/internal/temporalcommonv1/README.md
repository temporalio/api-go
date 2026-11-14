# temporalcommonv1

This package contains [payload_json.go](payload_json.go) which is copied to [../../common/v1](../../common/v1) after
proto generation. The code is kept here separately to keep the delete-entire-directory cleanup before code generation
simple.

This code adds "shorthand" formatting support to payloads by implementing `MaybeMarshalProtoJSON` and
`MaybeUnmarshalProtoJSON` on both `Payloads` and `Payload`. Tests for this code are in
[../../proxy/marshal_test.go](../../proxy/marshal_test.go).
