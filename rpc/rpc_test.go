package rpc_test

import (
	"educationalsp/rpc"
	"testing"
)

type EncodingExample struct {
  Testing bool
}

func TestEncode(t *testing.T) {
  expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
  actual := rpc.EncodeMessage(EncodingExample{Testing:true})
  if expected != actual {
    t.Fatalf("Expected: %s, Actual: %s", expected, actual)
  }
}
