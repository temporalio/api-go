# Temporal gRPC API and proto files compiled for Go

Generated Go files from Temporal [api](https://github.com/temporalio/api) repository.

_Note that any changes merged to [api](https://github.com/temporalio/api) will automatically trigger a [GitHub workflow](https://github.com/temporalio/api-go/blob/master/.github/workflows/update-proto.yml) that recompiles the proto files and commits them to master._

## How to use

To install in your project run:
```
go get -u go.temporal.io/api
```

## Rebuild

Run `make` once to install all plugins and tools (`protoc` and `go` must be installed manually).

Run `make update-proto` to update submodule and recompile proto files.

## License

MIT License, please see [LICENSE](LICENSE) for details.
