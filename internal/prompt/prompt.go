package prompt

import "github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/clients/generate"

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type PromptResponse struct {
	Response string `json:"response"`
}

type ErrorResponse struct { // TODO: consider movingig into delphierrors package
	Error     string `json:"error"`
	ErrorType string `json:"error_type"`
}

type API struct {
	GenerationClient generate.Client
}

func NewAPI() *API {
	generationClient := generate.NewClient("generate", "http://localhost:8081")
	return &API{GenerationClient: *generationClient}
}
