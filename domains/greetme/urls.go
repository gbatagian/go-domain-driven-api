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

func IncludeDomainURLS(handler *http.ServeMux) {
	utils.IncludeURLS(handler, URLS)
}
