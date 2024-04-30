package lsp

type ShowMessageRequest struct {
	Request
	Params ShowMessageParams `json:"params"`
}

type ShowMessageParams struct {
	MessageType MessageType `json:"type"`
	Message     string      `json:"message"`
}

type MessageType int

const (
	Error   MessageType = 1
	Warning MessageType = 2
	Info    MessageType = 3
	Log     MessageType = 4
	Debug   MessageType = 5
)

func ShowMessage(messageType MessageType, message string) ShowMessageRequest {
	// in a real LSP, this would look up the type in our type analysis code

	return ShowMessageRequest{
		Request{
			RPC:    "2.0",
			Method: "window/showMessage",
		},
		ShowMessageParams{
			MessageType: messageType,
			Message:     message,
		},
	}
}
