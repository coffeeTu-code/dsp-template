#!/usr/bin/env bash

# 执行上面命令之后会在 $GOPATH/bin目录生成protoc-gen-go,protoc-gen-go-grpc两个文件
#go get google.golang.org/protobuf/cmd/protoc-gen-go
#go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc --version

# grpc
function gen() {
  protoc -I . --go_out=paths=source_relative,plugins=grpc:. $1/*.proto

  sed -i "" 's/"base"/"dsp-template\/api\/base"/g' $1/*.go

  echo $1

  #  protoc -I ./example/api \
  #  --go_out ./example/api --go_opt=paths=source_relative \
  #  --go-gin_out ./example/api --go-gin_opt=paths=source_relative \
  #  example/api/product/app/v1/v1.proto
}

gen base
gen juno
gen polaris
gen rank
