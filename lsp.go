package main

// import (
// 	"github.com/tliron/glsp"
// 	protocol "github.com/tliron/glsp/protocol_3_16"
// 	glspserv "github.com/tliron/glsp/server"
// )
//
// type Lsp struct {
// 	Name    string
// 	Version string
// 	Handler *protocol.Handler
// }
//
// func CreateLSP(name string, version string, handler *protocol.Handler) Lsp {
// 	return Lsp{
// 		Name:    name,
// 		Version: version,
// 		Handler: handler,
// 	}
// }
//
// func CreateServer() {
// 	handler := protocol.Handler{}
// 	name := "mu-lsp"
// 	server := glspserv.NewServer(&handler, name, true)
//
// 	var clientCapabilities protocol.ClientCapabilities
//
// 	capabilities := handler.CreateServerCapabilities()
// 	capabilities.HoverProvider = true
//
// 	handler.TextDocumentDidChange = func(context *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
// 		doc, ok := server.documents.Get(params.TextDocument.URI)
// 		if !ok {
// 			return nil
// 		}
//
// 		doc.ApplyChanges(params.ContentChanges)
// 		server.refreshDiagnosticsOfDocument(doc, context.Notify, true)
// 		return nil
// 	}
//
// }
