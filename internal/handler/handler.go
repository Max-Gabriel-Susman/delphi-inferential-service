package handler

import (
	"net/http"

	"github.com/Max-Gabriel-Susman/delphi-go-kit/delphiweb"
	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/inference"
)

var _ http.Handler = (*delphiweb.App)(nil)

func API(d Deps) *delphiweb.App {
	app := delphiweb.NewApp()
	// dbrConn := database.NewDBR(d.DB)
	// accountAPI := account.NewAPI(account.NewMySQLStore(dbrConn))
	// AccountEndpoints(app, accountAPI)
	inferenceAPI := inference.NewAPI()
	InferenceEndpoints(app, inferenceAPI)
	return app
}
