package settings

import (
	"net/http"
	"os"
)

type Settings struct {
	Host   string
	Port   string
	Router *http.ServeMux
}

var DefaultSettings = Settings{
	Host:   "0.0.0.0",
	Port:   "8080",
	Router: http.NewServeMux(),
}

var EnvSettings = Settings{
	Host:   getApiHost(),
	Port:   getApiPort(),
	Router: http.NewServeMux(),
}

func getApiHost() string {
	host := os.Getenv("API_HOST")
	if host == "" {
		return DefaultSettings.Host
	}
	return host
}

func getApiPort() string {
	port := os.Getenv("API_PORT")
	if port == "" {
		return DefaultSettings.Port
	}
	return port
}
