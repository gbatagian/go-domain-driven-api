package utils

import "github.com/gin-gonic/gin"

type URL struct {
	Method     string
	Path       string
	Controller gin.HandlerFunc
}

func IncludeURLS(router *gin.Engine, urls []URL) *gin.Engine {

	for _, url := range urls {

		if url.Method == "GET" {
			router.GET(url.Path, url.Controller)
		} else if url.Method == "POST" {
			router.POST(url.Path, url.Controller)
		}

	}

	return router

}
