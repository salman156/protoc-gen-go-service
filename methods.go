package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func genMethod(g *protogen.GeneratedFile, method *protogen.Method) error {

	genUnaryMethod(g, method)
	genClientStreamMethod(g, method)
	genServerStreamMethod(g, method)
	genBiStreamMethod(g, method)

	return nil
}

func genUnaryMethod(g *protogen.GeneratedFile, method *protogen.Method) {
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		return
	}

	g.P("func (s *_", method.Parent.GoName, "Impl) ", method.GoName, "(ctx ", contextImport.Ident("Context"), ", req *", method.Input.GoIdent, ") (*", method.Output.GoIdent, ", error) {")

	genExtensions(g, method)

	g.P("	if s.interceptor != nil {")
	g.P("		handler := func(ctx context.Context, req any) (any, error) {")
	g.P("			return s.impl.", method.GoName, "(ctx, req.(*", method.Input.GoIdent, "))")
	g.P("		}")
	g.P("		info := &", grpcImport.Ident("UnaryServerInfo"), "{")
	g.P("			Server: s,")
	g.P("			FullMethod: ", method.Parent.GoName, "_", method.GoName, "_FullMethodName,")
	g.P("		}")
	g.P("")
	g.P("		resp, err := s.interceptor(ctx, &req, info, handler)")
	g.P("")
	g.P("		return resp.(*", method.Output.GoIdent, "), err")
	g.P("	}")
	g.P("")

	g.P("	return s.impl.", method.GoName, "(ctx, req)")
	g.P("}")
	g.P("")
}

func genClientStreamMethod(g *protogen.GeneratedFile, method *protogen.Method) {
	if !method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		return
	}

	g.P("func (s *_", method.Parent.GoName, "Impl) ", method.GoName, "(stream ",
		grpcImport.Ident("ClientStreamingServer"), "[", method.Input.GoIdent, ",", method.Output.GoIdent, "]) error {")
	g.P("  panic(\"not implemented\")")
	g.P("}")
}

func genServerStreamMethod(g *protogen.GeneratedFile, method *protogen.Method) {
	if method.Desc.IsStreamingClient() || !method.Desc.IsStreamingServer() {
		return
	}

	g.P("func (s *_", method.Parent.GoName, "Impl) ", method.GoName, "(req *", method.Input.GoIdent, ", stream ",
		grpcImport.Ident("ServerStreamingServer"), "[", method.Output.GoIdent, "]) error {")
	g.P("  panic(\"not implemented\")")
	g.P("}")
}

func genBiStreamMethod(g *protogen.GeneratedFile, method *protogen.Method) {
	if !method.Desc.IsStreamingClient() || !method.Desc.IsStreamingServer() {
		return
	}

	g.P("func (s *_", method.Parent.GoName, "Impl) ", method.GoName, "(stream ",
		grpcImport.Ident("BidiStreamingServer"), "[", method.Input.GoIdent, ",", method.Output.GoIdent, "]) error {")
	g.P("  panic(\"not implemented\")")
	g.P("}")

}
