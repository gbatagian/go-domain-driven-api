package utils

import (
	"fmt"
	"net/http"
)

type URL struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func IncludeURLS(hander *http.ServeMux, urls []URL) {
	for _, url := range urls {
		hander.HandleFunc(fmt.Sprintf("%s %s", url.Method, url.Path), url.Handler)
	}
}
