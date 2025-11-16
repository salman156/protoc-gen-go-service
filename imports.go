package main

import "google.golang.org/protobuf/compiler/protogen"

var (
	grpcImport    = protogen.GoImportPath("google.golang.org/grpc")
	contextImport = protogen.GoImportPath("context")
)
