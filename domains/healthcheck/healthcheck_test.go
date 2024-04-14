package healthcheck

import (
	"encoding/json"
	"fmt"
	"go-domain-driven-api/settings"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessfulCall(t *testing.T) {
	// arrange
	baseURL := "/healthcheck"
	testHandler := settings.DefaultSettings.Router
	RegisterDomainURLS(testHandler)

	req, _ := http.NewRequest("GET", baseURL, nil)
	resp := httptest.NewRecorder()

	// act
	testHandler.ServeHTTP(resp, req)

	// assert response code
	if resp.Code != 200 {
		t.Errorf("GET /healthcheck request unsuccessful")
	}

	// assert response payload
	var responsePayload map[string]string
	json.NewDecoder(resp.Body).Decode(&responsePayload)

	if fmt.Sprint(responsePayload) != fmt.Sprint(map[string]string{"feeling": "great"}) {
		t.Errorf(fmt.Sprintf("Unexpected GET /healthcheck response payload: %v", responsePayload))
	}
}
