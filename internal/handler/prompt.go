package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Max-Gabriel-Susman/delphi-go-kit/delphiweb"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/clients/generate"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/prompt"
)

type promptGroup struct {
	*prompt.API
}

func PromptEndpoints(app *delphiweb.App, api *prompt.API) {
	pg := promptGroup{API: api}

	app.Handle("GET", "/health", pg.HealthCheck) // TODO: consider healthchecing text generation inference and other services behind this once from client
	app.Handle("POST", "/prompt", pg.Prompt)
}

func (pg promptGroup) HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: implement tracer support

	// TODO: implemnet 503 Text generation inference is down logic

	fmt.Println("HealthCheck served") // delete

	return delphiweb.Respond(ctx, w, struct{}{}, http.StatusOK)
}

func (pg promptGroup) Prompt(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: implement tracer support

	// TODO: separate the business logic from the handler
	// var request prompt.PromptRequest // we need new types
	var promptRequest prompt.PromptRequest // we need new types
	if err := delphiweb.Decode(r.Body, &promptRequest); err != nil {
		return err
	}

	generateRequest := generate.GenerateInferenceRequest{
		Inputs: promptRequest.Prompt,
	}

	fmt.Println("prompt is: ", generateRequest.Inputs) // delete

	fmt.Println("pre generate req") // delete

	// TODO: make a call to text generation inference service hither
	generateResponse, err := pg.GenerationClient.Generate(ctx, generateRequest)
	if err != nil {
		return err
	}
	fmt.Println("post generate req") // delete

	promptResponse := prompt.PromptResponse{
		Response: generateResponse.GeneratedInference,
	}

	fmt.Println("generated inference is: ", generateResponse.GeneratedInference) // delete

	fmt.Println("prompt response is: ", promptResponse.Response) // delete

	return delphiweb.Respond(ctx, w, promptResponse, http.StatusOK)
}
