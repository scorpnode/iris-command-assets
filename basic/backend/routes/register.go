package routes

import (
	"github.com/kataras/iris"
)

func init() {

	// define the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", iris.Map{"Title": iris.StatusText(iris.StatusInternalServerError)})
	})

	// register the rest of the routes here
	iris.Handle("GET", "/", Index())

}
