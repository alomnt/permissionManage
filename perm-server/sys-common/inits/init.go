package inits

import (
	"permissionManage/perm-server/sys-common/inits/parse"
	"permissionManage/perm-server/sys-common/inits/sys"
	"permissionManage/perm-server/sys-sysbase/models/user"

	"github.com/kataras/golog"
)

func init() {
	parse.AppOtherParse()
	parse.DBSettingParse()

	initRootUser()
}

func initRootUser() {
	// root is existed?
	if sys.CheckRootExit() {
		return
	}

	// create root user
	sys.CreateRoot()

	ok := sys.CreateSystemRole()
	if ok {
		addRoleMenu()
	}

}

func addRoleMenu() {
	// 添加role-menu关系
	rMenus := []*user.RoleResource{
		{RoleId: "", ResourceId: ""},
	}
	effect, err := user.CreateRelationRoleResource(rMenus...)
	if err != nil {
		golog.Fatalf("**@@@@@@@@@@@0> %d, %s", effect, err.Error())
	}
	golog.Infof("@@@@@@@@@-> %s, %s", effect, err.Error())
}
