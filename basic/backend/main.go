package main

import (
	"github.com/kataras/iris"
)

func init() {
	// set the configs
	configuration := iris.Config()

	configuration.Render.Template.Directory = "../frontend/templates"
	configuration.Render.Template.Layout = "layout.html"

	// enable websocket server by setting its endpoint to 127.0.0.1/ws
	configuration.Websocket.Endpoint = "/ws"
}

func main() {
	// set the favicon
	iris.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	iris.Static("/public", "../frontend/public", 1)

	makeRouter()

	// start the server
	err := iris.ListenWithErr("127.0.0.1:80")
	if err != nil {
		panic(err)
	}
}

func makeRouter() {
	// define the custom errors
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("errors/404.html", iris.Map{"Title": "Not found"})
	})

	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Render("errors/500.html", iris.Map{"Title": "Server internal error"})
	})

	// set the routes
	iris.Handle("GET", "/", Index())
	// iris.Get("/", Index().Serve)
}
