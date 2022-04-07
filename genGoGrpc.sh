#!/bin/bash

export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=WikiLinkParser/proto --go_opt=paths=source_relative --go-grpc_out=WikiLinkParser/proto --go-grpc_opt=paths=source_relative WikiLinkParserService.proto
