#!/bin/bash
# build_protoc
## build protocol buffer for golang
echo "Build Protoc File Start..."

protoc --go_out=plugins=grpc:. songs.proto

echo "Build Protoc File End..."
