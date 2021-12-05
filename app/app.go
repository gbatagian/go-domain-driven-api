package app

import "github.com/gin-gonic/gin"

type App struct {
	Router *gin.Engine
}

func (app *App) Run() {

	app.Router.Run()

}

func RunWithEngine(engine *gin.Engine) {

	app := App{Router: engine}
	app.GetRoutes()
	app.Run()

}
