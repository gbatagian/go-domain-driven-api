package app

import (
	"go-domain-driven-api/domains/greetme"
	"go-domain-driven-api/domains/healthcheck"
)

func (app *App) GetRoutes() {
	// Healthcheck endpoint
	app.Router = healthcheck.IncludeDomainURLS(app.Router)

	// Greetme endpoint
	app.Router = greetme.IncludeDomainURLS(app.Router)
}
