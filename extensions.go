package main

import (
	"github.com/envoyproxy/protoc-gen-validate/validate"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func genExtensions(g *protogen.GeneratedFile, method *protogen.Method) {
	genValidation(g, method)
}

func genValidation(g *protogen.GeneratedFile, method *protogen.Method) {
	hasValidation := false
	for _, field := range method.Input.Fields {
		if proto.HasExtension(field.Desc.Options(), validate.E_Rules) {
			hasValidation = true
			break
		}
	}

	if hasValidation {
		g.P("  if err := req.ValidateAll(); err != nil {")
		g.P("    return nil, err")
		g.P("  }")
		g.P("")
	}
}
