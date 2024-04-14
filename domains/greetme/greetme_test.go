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
	testHandler := settings.DefaultSettings.Handler
	IncludeDomainURLS(testHandler)

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
	testHandler.ServeHTTP(resp, req)

	// assert response code
	if resp.Code != 200 {
		t.Errorf("POST /greetme request unsuccessful")
	}

	var responsePayload map[string]string
	json.NewDecoder(resp.Body).Decode(&responsePayload)

	if fmt.Sprint(responsePayload) != fmt.Sprint(map[string]string{"Greetings": "Mr. George"}) {
		t.Errorf(fmt.Sprintf("Unexpected POST /greetme response payload: %v", responsePayload))
	}
}
