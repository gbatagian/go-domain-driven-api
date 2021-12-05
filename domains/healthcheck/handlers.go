package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthcheckGet() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			map[string]string{
				"Feeling": "Great",
			},
		)
	}

}
