# Pokemon API gRPC

This repo contains the source code for the pokemon API gRPC server, found
in the [Dreams of Code YouTube video](https://youtube.com/@dreamsofcode)

You can deploy this code on acorn.io, for free, by clicking the button below

[![Run in Acorn](https://acorn.io/v1-ui/run/badge?image=docker.io+dreamsofcode+pokemon-api:acorn&ref=dreamsofcode&count=false&style=for-the-badge)](https://acorn.io/run/docker.io/dreamsofcode/pokemon-api:acorn?ref=dreamsofcode)

## Requirements

To run this code, you'll need the following dependencies:

- Go 1.20
- PostgreSQL 16
- Make

To make modifications to this code, make sure you have protobuf tools installed
as well.

### macOS

```
$ brew install protobuf protoc-gen-go protoc-gen-go-grpc
```

alternatively you can install the protoc plugins using golang

```
$ go install google.golang.org/grpc/cmd/protoc-gen-go@v1.1
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## Building

To build this code, run `go build .` or `make build`.

If you make any modifications to the protobuf, run `make generate`
