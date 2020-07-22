package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	api := app.Party("/api")
	app.Logger().SetLevel("debug")
	logger := app.Logger()
	mvc.Configure(api, func(mvc *mvc.Application) {
		mvc.Register(
			logger,
			NewEnvironment,
			NewDatabase,
		)
		mvc.Handle(new(Controller))
	})
	app.Run(iris.Addr("0.0.0.0:8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
