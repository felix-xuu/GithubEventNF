package main

import (
	"GithubEventNF/main/fun"
	"github.com/kataras/iris/v12"
)

func main() {
	fun.ConfigParse()
	app := iris.New()
	app.Post(fun.ParsedConfig.Server.Path, fun.Hook)
	_ = app.Listen(":" + fun.ParsedConfig.Server.Port)
}
