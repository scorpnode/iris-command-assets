package main

import (
	// just import them in order to call their init() to load the api and the routes
	_ "github.com/iris-contrib/iris-command-assets/basic/backend/api"
	_ "github.com/iris-contrib/iris-command-assets/basic/backend/routes"
	"github.com/kataras/iris"
)

func main() {
	// set the favicon
	iris.Favicon("../frontend/public/images/favicon.ico")

	// set static folder(s)
	iris.Static("/public", "../frontend/public", 1)

	// start the server
	iris.Listen("127.0.0.1:80")
}
