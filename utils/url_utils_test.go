package utils

import (
	"go-domain-driven-api/settings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var urls []URL
var testRouter *gin.Engine

func init() {

	urls = []URL{
		{
			Method:     "GET",
			Path:       "/foo",
			Controller: func(context *gin.Context) {},
		},
		{
			Method:     "GET",
			Path:       "/bar",
			Controller: func(context *gin.Context) {},
		},
		{
			Method:     "POST",
			Path:       "/baz",
			Controller: func(context *gin.Context) {},
		},
	}
	testRouter = settings.ENGINE

}

func TestIncludeURLS(t *testing.T) {

	testRouter = IncludeURLS(testRouter, urls)
	assert.Equal(t, len(testRouter.Routes()), 3)

	var urls []string
	for _, route := range testRouter.Routes() {
		urls = append(urls, route.Path)
	}
	assert.Equal(t, urls, []string{"/foo", "/bar", "/baz"})

}
