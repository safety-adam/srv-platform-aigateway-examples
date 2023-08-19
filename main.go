package main

import (
	"context"
	"fmt"

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

	example(outCtx, aiClient)
}

// This is a simple example of how to generate text using only a prompt.
func example(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	fmt.Println("See the tutorial in the README for more information.")
}
