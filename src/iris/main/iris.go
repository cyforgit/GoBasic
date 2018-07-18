package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("message", "Hello world!")

		ctx.View("index.html")
	})

	app.Run(iris.Addr(":8080"))

}
