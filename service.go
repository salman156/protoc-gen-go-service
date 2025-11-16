package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func genService(g *protogen.GeneratedFile, service *protogen.Service) error {
	if len(service.Methods) == 0 {
		return nil
	}

	g.P("type _", service.GoName, "Impl struct {")
	g.P("  impl ", service.GoName, "Server")
	g.P("  interceptor ", grpcImport.Ident("UnaryServerInterceptor"))
	g.P("}")

	g.P("func Register", service.GoName, "Impl(s ", grpcImport.Ident("ServiceRegistrar"), ", impl ", service.GoName, "Server, interceptor ", grpcImport.Ident("UnaryServerInterceptor"), ") {")
	g.P("  service := &_", service.GoName, "Impl{impl: impl, interceptor: interceptor}")
	g.P("")
	g.P("  s.RegisterService(&", service.GoName, "_ServiceDesc, service)")
	g.P("}")
	g.P("")

	for _, method := range service.Methods {
		err := genMethod(g, method)
		if err != nil {
			return err
		}
	}

	return nil
}
