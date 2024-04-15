package app

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAppCreation(t *testing.T) {
	handler := http.NewServeMux()
	app := App{Addr: "127.0.0.1:8080", Router: handler}
	app.RegisterURLS()

	routingIndex := reflect.ValueOf(*handler).FieldByName("index")
	segments := routingIndex.FieldByName("segments") // the registered routes

	if len(segments.MapKeys()) != 1 {
		t.Errorf("Expected routes not registered")
	}
}
