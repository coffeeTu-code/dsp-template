#!/usr/bin/env bash

protoc --version

# grpc
function gen() {
    protoc -I . --go_out=paths=source_relative,plugins=grpc:. $1/*.proto

    sed -i "" 's/"base"/"dsp-template\/api\/base"/g' $1/*.go
}

gen base
gen juno
gen polaris
gen rank
