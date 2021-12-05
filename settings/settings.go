package settings

import "github.com/gin-gonic/gin"

var ENGINE *gin.Engine

func init() {

	ENGINE = gin.Default()

}
