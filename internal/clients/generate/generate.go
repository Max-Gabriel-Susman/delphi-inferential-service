package generate

import (
	"context"
	"fmt"
)

func (c *Client) Generate(ctx context.Context, inferenceRequest GenerateInferenceRequest) (GeneratedInferenceResponse, error) {
	// if span, ok := tracer.SpanFromContext(ctx); ok {
	// 	span.SetTag(ext.ManualDrop, true)
	// } TODO: add tracing
	fmt.Println("Token Generation requested") // delete

	fmt.Println("inference request Inputs: ", inferenceRequest.Inputs) // delete
	const path = "/generate"
	req, err := c.NewRequest(ctx, "POST", path, inferenceRequest)
	if err != nil {
		return GeneratedInferenceResponse{}, err
	}
	var response GeneratedInferenceResponse
	// TODO: implement request headers
	if err := c.DoAndDecode(req, &response); err != nil {
		return GeneratedInferenceResponse{}, err
	}
	fmt.Println("Token Generation completed") // delete
	return response, nil
}
