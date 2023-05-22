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

	// Complete text example with a strucutred response
	completeTextStructured(outCtx, aiClient)

	// Extract text from image example
	extractTextFromImage(outCtx, aiClient)

	// Generate an image from a prompt
	generateImage(outCtx, aiClient)

	// Detect objects in image that would fail moderation
	objectDetection(outCtx, aiClient)

	// Detect objects in image that would fail moderation
	objectDetectionWithModeration(outCtx, aiClient)

	// Detect PPE equipment on people
	ppeDetectionWithModeration(outCtx, aiClient)
}

// This is a simple example of how to generate text using only a prompt.
func completeText(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	req := &aigateway.CompleteTextRequest{
		Prompt: "Write a short funny poem about a panda.",
	}
	resp, err := c.CompleteText(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("CompleteText example")
	fmt.Println(resp.Raw)
	fmt.Println()
}

// This example demonstrates how to request responses in a structured form by specifying a format for the response.
// The format is given in the form of a JSON example.
// The service will attempt to generate an ARRAY of responses in this format.
// The example should represent ONLY A SINGLE ELEMENT of the array.
func completeTextStructured(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	req := &aigateway.CompleteTextRequest{
		Prompt:          "You are an expert health and safety template engine. Create a template to clean a kitchen which asks 10 questions.",
		ResponseExample: `{"question": "Has the fridge been cleaned?"}`,
	}
	resp, err := c.CompleteText(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("CompleteTextStructured example")
	for i := 0; i < len(resp.Structured); i++ {
		fmt.Printf("%d. %s\n", i+1, resp.Structured[i].GetFields()["question"].GetStringValue())
	}
	fmt.Println()
}

// This example demonstraits how to extract text from an image
// The text is returned as an array of strings
func extractTextFromImage(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	b, _ := ioutil.ReadFile("sample.jpg")

	req := &aigateway.ExtractTextFromImageRequest{
		Document: b,
	}
	resp, err := c.ExtractTextFromImage(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ExtractTextFromImage example")
	for i := 0; i < len(resp.TextLines); i++ {
		fmt.Printf("Line %d: %s\n", i+1, resp.TextLines[i])
	}
	fmt.Println()
}

func generateImage(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	req := &aigateway.GenerateImageRequest{
		Prompt: "Create an eye catching image of a drone in a city scape. The drone should be the main focus of the image.",
	}
	resp, err := c.GenerateImage(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("GenerateImage example")
	fmt.Println(resp.ResponseUrl)
	fmt.Println()
}

func objectDetection(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	b, _ := ioutil.ReadFile("sydney.jpg")

	req := &aigateway.DetectObjectsInImageRequest{
		Image: b,
	}
	resp, err := c.DetectObjectsInImage(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ObjectDetection example")
	fmt.Println("Objects detected")
	for _, o := range resp.GetObjects() {
		fmt.Printf("%s x %d @ %f%% confidence\n", o.GetName(), o.GetCount(), o.GetConfidence())
	}
	fmt.Println()
}

func objectDetectionWithModeration(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	b, _ := ioutil.ReadFile("weapon.jpg")

	req := &aigateway.DetectObjectsInImageRequest{
		Image: b,
	}
	resp, err := c.DetectObjectsInImage(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ObjectDetectionWithModeration example")
	fmt.Println("Objects detected")
	for _, o := range resp.GetObjects() {
		fmt.Printf("%s x %d @ %f%% confidence\n", o.GetName(), o.GetCount(), o.GetConfidence())
	}

	fmt.Println()
	fmt.Println("Moderation labels")
	for _, o := range resp.GetModerationLabels() {
		fmt.Printf("%s @ %f%% confidence\n", o.GetName(), o.GetConfidence())
	}
	fmt.Println()
}

func ppeDetectionWithModeration(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	b, _ := ioutil.ReadFile("ppe.jpeg")

	req := &aigateway.DetectPPEInImageRequest{
		Image: b,
	}
	resp, err := c.DetectPPEInImage(ctx, req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("PPEDetectionWithModeration example")
	fmt.Println("PPE Detected")
	for _, p := range resp.GetPersons() {
		for _, parts := range p.GetParts() {
			fmt.Printf("Body Part: name - %s, PPE Type - %s, confidence- %f%% \n", parts.GetName(), parts.GetPpeType().String(), parts.GetConfidence())
		}
		fmt.Printf("%d person id \n", p.GetId())
	}

	fmt.Println()
	fmt.Println("Moderation labels")
	for _, o := range resp.GetModerationLabels() {
		fmt.Printf("%s @ %f%% confidence\n", o.GetName(), o.GetConfidence())
	}
	fmt.Println()
}
