package main

import (
	_ "permissionManage/perm-server/sys-common/inits"
	conf "permissionManage/perm-server/sys-common/inits/parse"
	"permissionManage/perm-server/sys-sysbase/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	routes.Hub(app)
	app.Run(iris.Addr(":"+conf.O.Port), iris.WithConfiguration(conf.C))
}
