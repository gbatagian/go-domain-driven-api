package app

import (
	"go-domain-driven-api/settings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppCreation(t *testing.T) {

	_app := App{Router: settings.ENGINE}
	_app.GetRoutes()

	routes := _app.Router.Routes()
	expected_number_of_routes := 2
	assert.Equal(t, len(routes), expected_number_of_routes)

	var urls []string
	for _, route := range routes {
		urls = append(urls, route.Path)
	}
	expected_urls := []string{
		"/healthcheck",
		"/greetme",
	}
	assert.Equal(t, urls, expected_urls)

}
