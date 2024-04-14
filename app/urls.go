package app

import (
	"go-domain-driven-api/domains/greetme"
	"go-domain-driven-api/domains/healthcheck"
)

func (app *App) RegisterRoutes() {
	// healthCheck endpoint
	healthcheck.IncludeDomainURLS(app.Handler)

	// GreetMe endpoint
	greetme.IncludeDomainURLS(app.Handler)
}
