package healthcheck

import (
	"go-domain-driven-api/utils"
	"net/http"
)

var URLS = []utils.URL{
	// All the domain urls must be initialized here.
	{
		Method:  "GET",
		Path:    "/healthcheck",
		Handler: HealthCheckGet(),
	},
}

func RegisterDomainURLS(router *http.ServeMux) {
	utils.RegisterURLS(router, URLS)
}
