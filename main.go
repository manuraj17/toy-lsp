package main

import (
	"bufio"
	"log"
	"os"
	"toy-lsp/rpc"
)

func main() {
	logger := getLogger("/Users/manu/Code/github/manuraj17/toy-lsp/log.txt")
	logger.Println("TOY LSP Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("No good log file")
	}

	return log.New(logfile, "[toy-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

// package main
//
// import (
// 	"toy-lsp/handlers"
//
// 	"github.com/tliron/commonlog"
// 	"github.com/tliron/glsp"
// 	protocol "github.com/tliron/glsp/protocol_3_16"
// 	"github.com/tliron/glsp/server"
//
// 	_ "github.com/tliron/commonlog/simple"
// )
//
// const lsName = "Emoji Autocomplete Language Server"
//
// var version string = "0.0.1"
// var handler protocol.Handler
//
// func main() {
// 	commonlog.Configure(2, nil)
//
// 	handler = protocol.Handler{
// 		Initialize:             initialize,
// 		Shutdown:               shutdown,
// 		TextDocumentCompletion: handlers.TextDocumentCompletion,
// 	}
//
// 	server := server.NewServer(&handler, lsName, true)
//
// 	server.RunStdio()
// }
//
// func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
// 	commonlog.NewInfoMessage(0, "Initializing server...")
//
// 	capabilities := handler.CreateServerCapabilities()
//
// 	capabilities.CompletionProvider = &protocol.CompletionOptions{}
//
// 	return protocol.InitializeResult{
// 		Capabilities: capabilities,
// 		ServerInfo: &protocol.InitializeResultServerInfo{
// 			Name:    lsName,
// 			Version: &version,
// 		},
// 	}, nil
// }
//
// func shutdown(context *glsp.Context) error {
// 	return nil
// }
//
// // package main
// //
// // import (
// // 	"github.com/tliron/commonlog"
// // 	"github.com/tliron/glsp"
// // 	protocol "github.com/tliron/glsp/protocol_3_16"
// // 	"github.com/tliron/glsp/server"
// //
// // 	// Must include a backend implementation
// // 	// See CommonLog for other options: https://github.com/tliron/commonlog
// // 	_ "github.com/tliron/commonlog/simple"
// // )
// //
// // const lsName = "my language"
// //
// // var version string = "0.0.1"
// // var handler protocol.Handler
// //
// // func main() {
// // 	// This increases logging verbosity (optional)
// // 	commonlog.Configure(1, nil)
// //
// // 	handler = protocol.Handler{
// // 		Initialize:  initialize,
// // 		Initialized: initialized,
// // 		Shutdown:    shutdown,
// // 		SetTrace:    setTrace,
// // 	}
// //
// // 	server := server.NewServer(&handler, lsName, false)
// //
// // 	server.RunStdio()
// // }
// //
// // func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
// // 	capabilities := handler.CreateServerCapabilities()
// //
// // 	return protocol.InitializeResult{
// // 		Capabilities: capabilities,
// // 		ServerInfo: &protocol.InitializeResultServerInfo{
// // 			Name:    lsName,
// // 			Version: &version,
// // 		},
// // 	}, nil
// // }
// //
// // func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
// // 	return nil
// // }
// //
// // func shutdown(context *glsp.Context) error {
// // 	protocol.SetTraceValue(protocol.TraceValueOff)
// // 	return nil
// // }
// //
// // func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
// // 	protocol.SetTraceValue(params.Value)
// // 	return nil
// // }
// //
// // func handleCompletion(params *protocol.CompletionParams) (*protocol.CompletionList, error) {
// // 	// Example logic for completion
// // 	completionItems := []protocol.CompletionItem{
// // 		{Label: "ExampleCompletion1"},
// // 		{Label: "ExampleCompletion2"},
// // 	}
// //
// // 	return &protocol.CompletionList{
// // 		Items: completionItems,
// // 	}, nil
// // }
