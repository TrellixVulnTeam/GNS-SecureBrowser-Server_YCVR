#!/bin/bash
#
# Generates protocol buffer for golang server and client
protoc --go_out=proto --go-grpc_out=proto proto/gnsservice.proto

# Generate protocol buffer for macOS swift client
protoc -I proto gnsservice.proto \
    --swift_out=macos_client \
    --swiftgrpc_out=Client=true,Server=false:macos_client

# Generate protocol buffer for macOS objective-C client
#protoc -I proto  --objc_out=macos_client --objcgrpc_out=macos_client gnsservice.proto