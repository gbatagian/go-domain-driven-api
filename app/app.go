package app

import (
	"fmt"
	"net/http"

	"go-domain-driven-api/settings"
)

type App struct {
	Addr    string
	Handler *http.ServeMux
}

func (a *App) Run() {
	s := http.Server{Addr: a.Addr, Handler: a.Handler}
	s.ListenAndServe()
}

func RunWithSettings(s settings.Settings) {
	app := App{
		Addr:    fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler: s.Handler,
	}
	app.RegisterRoutes()
	app.Run()
}
