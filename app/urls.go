package app

import (
	"go-domain-driven-api/domains/healthcheck"
)

func (app *App) RegisterRoutes() {
	// Healthcheck endpoint
	healthcheck.IncludeDomainURLS(app.Handler)

	// Greetme endpoint
	// greetme.IncludeDomainURLS(app.Handler)
}
