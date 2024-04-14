package settings

import "net/http"

type Settings struct {
	Host    string
	Port    int
	Handler *http.ServeMux
}

var DefaultSettings = Settings{
	Host:    "127.0.0.1",
	Port:    8080,
	Handler: http.NewServeMux(),
}
