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

var IncludeDomainURLS = func(handler *http.ServeMux) {
	utils.IncludeURLS(handler, URLS)
}
