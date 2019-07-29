package main

import (
	_ "../sys-common/inits"
	conf "../sys-common/inits/parse"
	"../sys-sysbase/routes"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	routes.Hub(app)
	app.Run(iris.Addr(":"+conf.O.Port), iris.WithConfiguration(conf.C))
}
