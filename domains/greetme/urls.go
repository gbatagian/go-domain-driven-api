package greetme

import (
	"go-domain-driven-api/utils"
	"net/http"
)

var URLS = []utils.URL{
	// All the domain urls must be initialized here.
	{
		Method:  "POST",
		Path:    "/greetme",
		Handler: GreetMePost(),
	},
}

func RegisterDomainURLS(router *http.ServeMux) {
	utils.RegisterURLS(router, URLS)
}
