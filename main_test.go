package main

//
// import (
// 	"testing"
//
// 	protocol "github.com/tliron/glsp/protocol_3_16"
// )
//
// func TestHandleCompletion(t *testing.T) {
// 	// Mock parameters for the completion request
// 	params := &protocol.CompletionParams{
// 		TextDocument: protocol.TextDocumentIdentifier{
// 			URI: "file://path/to/testfile",
// 		},
// 		Position: protocol.Position{
// 			Line:      0,
// 			Character: 0,
// 		},
// 	}
//
// 	// Call the function to be tested
// 	completionList, err := handleCompletion(params)
// 	if err != nil {
// 		t.Errorf("handleCompletion returned an error: %v", err)
// 	}
//
// 	// Check if the returned value is as expected
// 	if completionList == nil || len(completionList.Items) == 0 {
// 		t.Errorf("Expected non-empty completion list, got nil or empty")
// 	}
//
// 	// Further detailed tests can be added here, such as checking the content of the completion items
// 	expectedLabel := "ExampleCompletion1"
// 	if completionList.Items[0].Label != expectedLabel {
// 		t.Errorf("Expected label %v, got %v", expectedLabel, completionList.Items[0].Label)
// 	}
// }
