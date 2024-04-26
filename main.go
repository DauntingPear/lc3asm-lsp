package main

import (
	"bufio"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	logger := getLogger("/Users/adrianbrady/Workspace/Projects/educationalsp/log.txt")
	logger.Println("Hey, I started")
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}

		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not parse: %s", err)
		}

		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)

		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Print("Sent the reply")
	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not parse: %s", err)
		}

		logger.Printf("Connected to: %s %s", request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, you didnt give me a good file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
