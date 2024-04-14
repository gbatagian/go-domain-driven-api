package app

import (
	"go-domain-driven-api/domains/greetme"
	"go-domain-driven-api/domains/healthcheck"
)

func (app *App) RegisterURLS() {
	// healthCheck endpoint
	healthcheck.RegisterDomainURLS(app.Router)

	// GreetMe endpoint
	greetme.RegisterDomainURLS(app.Router)
}
