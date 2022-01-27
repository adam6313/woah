#!/bin/bash

dir=$1

cd $dir

protoc \
    -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf \
    --gogofaster_out=Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,plugins=grpc:. \
    *.proto