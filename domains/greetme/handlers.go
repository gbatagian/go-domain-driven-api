package greetme

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GreetmePost() gin.HandlerFunc {

	return func(context *gin.Context) {

		// Serialize request payload.
		var requestBody GreetmeSerializer
		err := context.BindJSON(&requestBody)
		if err != nil {
			log.Fatal(err)
		}

		// Generate respond.
		context.JSON(
			http.StatusOK,
			map[string]string{
				"Greetings": strings.Trim(
					fmt.Sprintf("%s %s", requestBody.Title, requestBody.Name),
					" ",
				),
			},
		)

	}

}
