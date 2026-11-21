# Temporal gRPC API and proto files compiled for Go

Generated Go files from Temporal [api](https://github.com/temporalio/api) repository.

_Note that any changes merged to [api](https://github.com/temporalio/api) will automatically trigger a [GitHub workflow](https://github.com/temporalio/api-go/blob/main/.github/workflows/update-proto.yml) that recompiles the proto files and commits the results to main._

## How to use

To install in your project run:
```
go get go.temporal.io/api
```

## Draft API

`api_next/` contains the bleeding-edge Go API. `make api-next` generates it from the
`temporal/api_next` tree in the `proto/api` submodule. The `OmitDraftFields` methods
in `*.draft.pb.go` clear draft fields at runtime when their features are disabled.

The draft and stable modules both declare `go.temporal.io/api`. The draft module is a
drop-in replacement with identical import paths, and each program links exactly one version.

To use the draft API, add a replacement in your project's module using a fetchable
`api-go` commit. Pin a commit from the target branch after the PR merges, since a
squash merge can leave the original PR commits unreachable. Substitute the pinned
commit SHA for `YOUR_COMMIT_SHA` in the command below.

```sh
go mod edit -replace=go.temporal.io/api=github.com/temporalio/api-go/api_next@YOUR_COMMIT_SHA
go mod tidy
```

The resulting directive in `go.mod` uses the commit's pseudo-version:

```go.mod
replace go.temporal.io/api => github.com/temporalio/api-go/api_next v0.0.0-<timestamp>-<short-commit>
```

Tests that need both trees use separate modules. The `internal/draftwire` module links
`api_next` and compiles the stable tree's protos with `protoc`.

## Rebuild

Run `make` once to install all plugins and tools (`protoc` and `go` must be installed manually).

Run `make update-proto` to update the `proto/api` submodule and recompile proto files. The `proto/api-cloud` submodule
must be updated manually.

## License

MIT License, please see [LICENSE](LICENSE) for details.
