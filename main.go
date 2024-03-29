package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"toy-lsp/rpc"

	"toy-lsp/lsp"
)

func main() {
	logger := getLogger("/Users/manu/Code/github/manuraj17/toy-lsp/log.txt")
	logger.Println("TOY LSP Started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s\n", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, couldn't parse this: %s", err)
		}

		logger.Printf(
			"Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)

		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		_, _ = writer.Write([]byte(reply))

		logger.Println("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, we couldn't parse this: %s", err)
		}
		logger.Printf(
			"Opened: %s %s\n",
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text,
		)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("No good log file")
	}

	return log.New(logfile, "[toy-lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
