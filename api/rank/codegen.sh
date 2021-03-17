#!/usr/bin/env bash

protoc --version

# grpc
protoc -I . --go_out=paths=source_relative,plugins=grpc:. *.proto

# micro rpc
#protoc -I . --go_out=paths=source_relative:./go/ --micro_out=paths=source_relative:./go/ *.proto

