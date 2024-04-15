package app

import (
	"fmt"
	"log"
	"net/http"

	"go-domain-driven-api/settings"
)

type App struct {
	Addr   string
	Router *http.ServeMux
}

func (a *App) Run() {
	log.Printf("Starting server at %s", a.Addr)

	s := http.Server{Addr: a.Addr, Handler: a.Router}
	s.ListenAndServe()
}

func RunWithSettings(s settings.Settings) {
	app := App{
		Addr:   fmt.Sprintf("%s:%s", s.Host, s.Port),
		Router: s.Router,
	}
	app.RegisterURLS()
	app.Run()
}
