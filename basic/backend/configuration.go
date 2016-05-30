package main

import "github.com/kataras/iris"

func init() {
	// set the configs
	configuration := iris.Config()

	configuration.Render.Template.Directory = "../frontend/templates"
	configuration.Render.Template.Layout = "layout.html"

}
