all: gogen nodegen

.PHONY: gogen nodegen

gogofaster:
		go get github.com/gogo/protobuf/protoc-gen-gogofaster

src=*.proto

gogen: gogofaster
		protoc --proto_path=$(GOPATH)/src/github.com/gogo/protobuf/:. --gogofaster_out=plugins=grpc:. $(src)
		mv *.go ../server/meta

nodegen: 
	    protoc --proto_path=$(GOPATH)/src/github.com/gogo/protobuf/:. --js_out=import_style=commonjs,binary:./ --plugin=protoc-gen-grpc=../app/v0.0.1/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./ $(src)
		mv *.js ../app/v0.0.1/node_modules/meta/
