package sys

import (
	"permissionManage/perm-server/sys-common/middleware/casbins"

	"github.com/kataras/golog"
)

var (
	// 定义系统初始的角色
	Components = [][]string{
		{"0", "角色管理", "role_admin*", "角色管理"},
		{"0", "demo角色", "role_demo*", "demo角色"},
	}
)

// 创建系统默认角色
func CreateSystemRole() bool {
	e := casbins.GetEnforcer()

	for _, v := range Components {
		p := e.GetFilteredPolicy(0, v[0])
		if len(p) == 0 {
			if ok := e.AddPolicy(v); !ok {
				golog.Fatalf("初始化角色[%s]权限失败。%s", v)
			}
		}
	}
	return true
}
