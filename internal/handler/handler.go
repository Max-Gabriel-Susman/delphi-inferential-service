package handler

import (
	"net/http"

	"github.com/Max-Gabriel-Susman/delphi-go-kit/delphiweb"
)

var _ http.Handler = (*delphiweb.App)(nil)

func API(d Deps) *delphiweb.App {
	app := delphiweb.NewApp()
	return app
}
