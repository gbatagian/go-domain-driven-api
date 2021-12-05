package greetme

import (
	"go-domain-driven-api/settings"
	"go-domain-driven-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var baseURL string
var testRouter *gin.Engine

func init() {

	baseURL = "/greetme"
	testRouter = settings.ENGINE
	testRouter = IncludeDomainURLS(testRouter)

}

func TestSuccessfulCall(t *testing.T) {

	request_body := map[string]interface{}{
		"name":  "George",
		"title": "Mr.",
	}
	req, err := http.NewRequest("POST", baseURL, utils.ToJsonBytesStream((request_body)))
	assert.Nil(t, err)

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)

	expected_response := map[string]interface{}{
		"Greetings": "Mr. George",
	}
	assert.Equal(t, resp.Body.String(), utils.ToJsonString(expected_response))

}
