#!/bin/sh

apt-get update

apt install -y protobuf-compiler

export PATH=$PATH:/usr/local/go/bin/go
export GO111MODULE=on

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

export PATH="$PATH:$(go env GOPATH)/bin"
