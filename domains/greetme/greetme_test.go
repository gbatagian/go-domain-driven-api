package greetme

import (
	"encoding/json"
	"fmt"
	"go-domain-driven-api/settings"
	"go-domain-driven-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessfulCall(t *testing.T) {
	// arrange
	baseURL := "/greetme"
	router := settings.DefaultSettings.Router
	RegisterDomainURLS(router)

	requestPayload := map[string]string{
		"name":  "George",
		"title": "Mr.",
	}
	req, _ := http.NewRequest(
		"POST",
		baseURL,
		utils.ToJsonBytesStream(requestPayload),
	)
	resp := httptest.NewRecorder()

	// act
	router.ServeHTTP(resp, req)

	// assert response code
	if resp.Code != 200 {
		t.Errorf("POST /greetme request unsuccessful")
	}

	// assert response payload
	var responsePayload map[string]string
	json.NewDecoder(resp.Body).Decode(&responsePayload)

	if fmt.Sprint(responsePayload) != fmt.Sprint(map[string]string{"Greetings": "Mr. George"}) {
		t.Errorf(fmt.Sprintf("Unexpected POST /greetme response payload: %v", responsePayload))
	}
}
