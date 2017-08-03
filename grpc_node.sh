#!/bin/bash

./app/v0.0.1/node_modules/grpc-tools/bin/protoc --js_out=import_style=commonjs,binary:./ --plugin=protoc-gen-grpc=./app/v0.0.1/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./app/v0.0.1/assest/js/ ./protoc/HelloWorldService.proto
