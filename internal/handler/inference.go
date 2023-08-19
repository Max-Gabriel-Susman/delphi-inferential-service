package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Max-Gabriel-Susman/delphi-go-kit/delphiweb"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/inference"
)

type inferenceGroup struct {
	*inference.API
}

func InferenceEndpoints(app *delphiweb.App, api *inference.API) {
	ig := inferenceGroup{API: api}

	app.Handle("GET", "/health", ig.HealthCheck)
	app.Handle("POST", "/generate", ig.GenerateTokens)
}

func (ig inferenceGroup) HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: implement tracer support

	// TODO: implemnet 503 Text generation inference is down logic

	fmt.Println("HealthCheck served")

	return delphiweb.Respond(ctx, w, struct{}{}, http.StatusOK)
}

func (ig inferenceGroup) GenerateTokens(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: implement tracer support

	// TODO: separate the business logic from the handler
	var request inference.GenerateInferenceRequest
	if err := delphiweb.Decode(r.Body, &request); err != nil {
		return err
	}

	fmt.Println("un-tokenized input: ", request.Inputs) // delete

	tokenizedInput := request.TokenizeInput()

	fmt.Println("tokenized input: ", tokenizedInput) // delete

	response := inference.GeneratedInferenceResponse{
		Details: inference.GeneratedInferenceResponseDetails{
			Tokens: tokenizedInput,
		},
	}

	fmt.Println("pre reversed tokens: ", response.Details.Tokens) // delete

	// reverse tokens // just for testing
	response.ReverseTokens()

	fmt.Println("post reversed tokens: ", response.Details.Tokens) // delete

	fmt.Println("Tokens generated") // delete

	return delphiweb.Respond(ctx, w, response, http.StatusOK)
}
