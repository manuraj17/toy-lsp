package lsp

type Request struct {
	RPC    string `json:"jsonrpc,omitempty"`
	ID     int    `json:"id,omitempty"`
	Method string `json:"method,omitempty"`
}

type Response struct {
	RPC string `json:"jsonrpc,omitempty"`
	ID  *int   `json:"id,omitempty"`
}

type Notification struct {
	RPC    string `json:"jsonrpc,omitempty"`
	Method string `json:"method,omitempty"`
}
