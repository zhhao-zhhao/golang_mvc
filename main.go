package main

//go get -u github.com/kataras/iris
import (
	"fmt"

	//"./service/tetService"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	//tetService.printTxt()
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(ctx iris.Context) {

		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Get("/zh/zhanghao", func(ctx iris.Context) {

		ctx.ViewData("message", "Hello zhanghao!")
		ctx.View("work/zh.html")
	})
	app.Get("/ping", func(ctx iris.Context) {
		fmt.Println("------in cntroller ping-------")
		ctx.WriteString("pong")
	})
	// Method: GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})
	app.Run(iris.Addr(":8080"))
}
