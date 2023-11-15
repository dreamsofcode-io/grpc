# Dreams of Code gRPC

This project provides the source code for the [gRPC video](https://youtu.be/_4TPM6clQjM) by [Dreams of Code](https://youtube.com/@dreamsofcode).

## Video instructions

In order to install protobuf on your own machine, you can use the following
commands

### macOS

```
$ brew install protobuf protoc-gen-go protoc-gen-go-grpc
```

alternatively you can install the protoc plugins using golang

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

### Arch Linux

```
$ pacman -S protobuf
```

#### Plugins

Using yay

```
$ yay -S protoc-gen-go protoc-gen-go-grpc
```

Using golang

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

## GRPC Clients

In this video, I use [grpcui](https://github.com/fullstorydev/grpcui) and [grpcurl](https://github.com/fullstorydev/grpcurl)

## Calculator

The calculator code can be found in the [./calulator](./calulator) directory.

You can also deploy an instance of it yourself on acorn by clicking the following button:

[![Run in Acorn](https://acorn.io/v1-ui/run/badge?image=docker.io+dreamsofcode+grpcalculator:acorn&ref=dreamsofcode&count=false&style=for-the-badge)](https://acorn.io/run/docker.io/dreamsofcode/grpcalculator:acorn?ref=dreamsofcode)

## Pokedex

The Pokedex API can be found in the [./pokemon-api](./pokemon-api) directory.

You can also deploy an instance of it yourself on acorn by clicking the following button:

[![Run in Acorn](https://acorn.io/v1-ui/run/badge?image=docker.io+dreamsofcode+pokemon-api:acorn&ref=dreamsofcode&count=false&style=for-the-badge)](https://acorn.io/run/docker.io/dreamsofcode/pokemon-api:acorn?ref=dreamsofcode)

## Disabling TLS

When writing gRPC code, you can disable tls by using the following lines

```go
opts := []grpc.DialOption{
    grpc.WithTransportCredentials(insecure.NewCredentials())
}
```

If you're wanting to connect locally, you'll need to use the `-plaintext`
flag with both grpcui and grpcurl
