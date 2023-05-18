package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/SafetyCulture/s12-apis-go/aigateway/v1"
)

var identity string = "ai-gateway-examples"
var namespace string = "actions"

func main() {
	// Get an instance of an AIGateway client
	aiClient := getClient()

	// Get a Soter admin token
	adminToken, err := getSoterAdminToken()
	if err != nil {
		panic(err)
	}
	// Put the token in an outgoing context
	outCtx := getOutgoingContext(adminToken)

	// Complete text example
	completeText(outCtx, aiClient)

	// Extract text from image example
	extractTextFromImage(outCtx, aiClient)
}

func completeText(ctx context.Context, c aigateway.AIGatewayClient) {
	req := &aigateway.CompleteTextRequest{
		Prompt: "Write a short poem about a panda.",
	}
	resp, err := c.CompleteText(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Raw)
}

func extractTextFromImage(ctx context.Context, c aigateway.AIGatewayClient) {
	b, _ := ioutil.ReadFile("sample.jpg")

	req := &aigateway.ExtractTextFromImageRequest{
		Document: b,
	}
	resp, err := c.ExtractTextFromImage(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.TextLines)
}
