package api

import "github.com/kataras/iris"

func init() {
	// register the api here
	iris.API("/users", UserAPI{})
}
