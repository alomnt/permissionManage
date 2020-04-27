package routes

import (
	"permissionManage/perm-server/sys-common/middleware"
	"permissionManage/perm-server/sys-common/supports"

	conf "permissionManage/perm-server/sys-common/inits/parse"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	rcover "github.com/kataras/iris/v12/middleware/recover"
	"permissionManage/perm-server/sys-sysbase/routes/user"
)

// 所有的路由
func Hub(app *iris.Application) {
	preSettring(app)
	var main = corsSetting(app)

	// 用户模块路由
	user.UserHub(main)       // 用户API模块
	user.ResoureHub(main)    // 菜单资源
	user.RoleHub(main)       // 角色
	user.DepartmentHub(main) // 部门
}

func corsSetting(app *iris.Application) (main iris.Party) {
	var (
		crs context.Handler
	)

	crs = cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //允许通过的主机名称
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
		//AllowCredentials: true,
	})

	/* 定义路由 */
	main = app.Party("/perm", crs).AllowMethods(iris.MethodOptions)
	main.Use(middleware.ServeHTTP)
	//main := app.Party("/")
	//main.Use(middleware.ServeHTTP)

	return main
}

func preSettring(app *iris.Application) {
	app.Logger().SetLevel(conf.O.LogLevel)

	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(
		rcover.New(),
		customLogger,
		//middleware.ServeHTTP
	)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		supports.Error(ctx, iris.StatusNotFound, supports.NotFound, nil)
	})
	//app.OnErrorCode(iris.StatusForbidden, customLogger, func(ctx iris.Context) {
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//})
	//捕获所有http错误:
	//app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	ctx.JSON(utils.Error(500, "服务器内部错误", nil))
	//})
}
