package main

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	i := 0
	// let's simulate a panic every next request
	app.Get("/", func(ctx iris.Context) {
		i++
		if i%2 == 0 {
			panic("a panic here")
		}
		ctx.Writef("Hello, refresh one time more to get panic!")
	})

	// http://localhost:8080, refresh it 5-6 times.
	app.Run(iris.Addr(":8080"))

}
