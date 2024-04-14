package utils

import (
	"net/http"
	"reflect"
	"testing"
)

func TestIncludeURLS(t *testing.T) {
	urls := []URL{
		{
			Method:  "GET",
			Path:    "/foo",
			Handler: func(w http.ResponseWriter, r *http.Request) {},
		},
		{
			Method:  "POST",
			Path:    "/bar",
			Handler: func(w http.ResponseWriter, r *http.Request) {},
		},
		{
			Method:  "PUT",
			Path:    "/baz",
			Handler: func(w http.ResponseWriter, r *http.Request) {},
		},
	}
	testHandler := http.NewServeMux()

	RegisterURLS(testHandler, urls)

	routingIndex := reflect.ValueOf(*testHandler).FieldByName("index")
	segments := routingIndex.FieldByName("segments") // the registered routes

	if len(segments.MapKeys()) != 3 {
		t.Errorf("Expected routes not registered")
	}
}
