package healthcheck

import (
	"go-domain-driven-api/utils"

	"github.com/gin-gonic/gin"
)

var URLS []utils.URL
var IncludeDomainURLS (func(router *gin.Engine) *gin.Engine)

func init() {

	// All the domain urls must be initialized here.
	URLS = []utils.URL{
		{
			Method:     "GET",
			Path:       "/healthcheck",
			Controller: HealthcheckGet(),
		},
	}

	IncludeDomainURLS = func(router *gin.Engine) *gin.Engine {
		return utils.IncludeURLS(router, URLS)
	}

}
