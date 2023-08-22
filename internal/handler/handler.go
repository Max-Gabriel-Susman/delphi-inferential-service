package handler

import (
	"net/http"

	"github.com/Max-Gabriel-Susman/delphi-go-kit/delphiweb"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/prompt"
)

var _ http.Handler = (*delphiweb.App)(nil)

func API(d Deps) *delphiweb.App {
	app := delphiweb.NewApp()
	promptAPI := prompt.NewAPI()
	PromptEndpoints(app, promptAPI)
	return app
}
