package settings

import "net/http"

type Settings struct {
	Host   string
	Port   int
	Router *http.ServeMux
}

var DefaultSettings = Settings{
	Host:   "0.0.0.0",
	Port:   8080,
	Router: http.NewServeMux(),
}
