package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
}
