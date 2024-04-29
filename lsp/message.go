package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// Params - defined later
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// Result
	// Err
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
