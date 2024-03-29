package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params,omitempty"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientinfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result,omitempty"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
	ServerInfo   ServerInfo         `json:"serverInfo,omitempty"`
}

type ServerCapabilities struct {
	TextDocumentSync int `json:"textDocumentSync"`
}

type ServerInfo struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumentSync: 1,
			},
			ServerInfo: ServerInfo{
				Name:    "toy-lsp",
				Version: "0.0.0-beta",
			},
		},
	}
}
