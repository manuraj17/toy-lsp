package lsp

type TextDocumentItem struct {
	URI        string `json:"uri,omitempty"`
	LanguageID string `json:"languageId,omitempty"`
	Version    int    `json:"version,omitempty"`
	Text       string `json:"text,omitempty"`
}

type TextDocumentIdentifier struct {
	URI string `json"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	version int `json:"version"`
}
