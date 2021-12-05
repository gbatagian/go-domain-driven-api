package main

import (
	"go-domain-driven-api/app"
	"go-domain-driven-api/settings"
)

func main() {

	app.RunWithEngine(settings.ENGINE)

}
