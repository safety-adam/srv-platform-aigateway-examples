# AIGateway quick start

This repository contains examples and quick start guides for the AI Gateway. 

Follow the instructions below to get started.

ℹ️ Make sure you are connected to the **Development VPN** before running the examples.

1. Clone this repository

```bash
% git clone git@github.com:SafetyCulture/srv-platform-aigateway-examples.git
```

2. **Connect to the Development VPN**

3. Change to the root of the repository

4. Run the examples with the following example

```bash
srv-platform-aigateway-examples% go run .
```

# Tutorial

This tutorial will take you step by step through the features of the AIGateway. At the end of this tutorial you will have a working example that you can use as a starting point for your own AIGateway integrations.

## Prerequisites

- [Go](https://golang.org/doc/install) installed
- AWS VPN Client installed
- Development environment VPN configuration installed
- VPN connection to the development environment established

## Clone the repository

Clone the repository's `tutorial-start` branch to your local machine.

```bash
% git clone --branch tutorial-start git@github.com:SafetyCulture/srv-platform-aigateway.git
```

## Running the tutorial

After each step below, run the project using the command below to see the changes.

```bash
% go run .
```

## Step 1: Generating text

In this step we will generate text using the `CompleteText` endpoint.

In the `main.go` file, add the following code under the `main` function.

```go
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
```

In the `main` function, add the following code to call the `completeText` function.

```go
completeText(outCtx, aiClient)
```

Save the file and run the project (see "Running the tutorial" above).

## Step 2: Generating structured text / data

In this step we will generate structured text using the same `CompleteText` endpoint.

In the `main.go` file, add the following code under the `main` function.

```go
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
```

In the `main` function, add the following code to call the `completeTextStructured` function.

```go
completeTextStructured(outCtx, aiClient)
```

Save the file and run the project (see "Running the tutorial" above).

## Step 3: Generating an image

In this step we will generate an image using the `GenerateImage` endpoint.

In the `main.go` file, add the following code under the `main` function.

```go
// This example demonstrates how to generate an image using the GenerateImage endpoint.
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
```

In the `main` function, add the following code to call the `generateImage` function.

```go
generateImage(outCtx, aiClient)
```

Save the file and run the project (see "Running the tutorial" above).

## Step 4: Object detection with moderation

In this step we will detect objects in an image using the `DetectObjects` endpoint.

In the `main.go` file, add the following code under the `main` function.

```go
// This example demonstrates how to detect objects in an image using the DetectObjects endpoint.
func objectDetectionWithModeration(ctx context.Context, c aigateway.AIGatewayServiceClient) {
	b, _ := os.ReadFile("weapon.jpg")

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
```

In the `main` function, add the following code to call the `objectDetectionWithModeration` function.

```go
objectDetectionWithModeration(outCtx, aiClient)
```

Save the file and run the project (see "Running the tutorial" above).

## Conclusion

Congratulations! You have completed the tutorial. You can now use this project as a starting point for your own AIGateway integrations.

There are more complex examples in the `main` branch of the repository. Feel free to explore these examples to see how to use the other endpoints of the AIGateway.
