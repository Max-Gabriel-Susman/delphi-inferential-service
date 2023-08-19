package handler

import (
	"context"
	"net/http"
	"fmt"

	"github.com/Max-Gabriel-Susman/delphi-go-kit/delphiweb"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/inference"
)

type inferenceGroup struct {
	*inference.API
}

func InferenceEndpoints(app *delphiweb.App, api *inference.API) {
	ig := inferenceGroup{API: api}

	app.Handle("GET", "/health", ig.HealthCheck)
	app.Handle("GET", "/generate", ig.GenerateTokens)
}

func (ig inferenceGroup) HealthCheck(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: implement tracer support

	// TODO: implemnet 503 Text generation inference is down logic

	fmt.Println("HealthCheck served")

	return delphiweb.Respond(ctx, w, struct{}{}, http.StatusOK)
}

func (ig inferenceGroup) GenerateTokens(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: implement tracer support

	

	fmt.Println("Tokens generated")

	return delphiweb.Respond(ctx, w, struct{}{}, http.StatusOK)
}
