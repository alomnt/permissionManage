package user

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"permissionManage/perm-server/sys-common/models/baseModel"
	"permissionManage/perm-server/sys-common/supports"
	"permissionManage/perm-server/sys-common/supports/commonConst"
	"permissionManage/perm-server/sys-common/supports/responseHandle"
	"permissionManage/perm-server/sys-sysbase/service/userService"
	"permissionManage/perm-server/sys-sysbase/vo/uservo"
)

func RoleHub(party iris.Party) {
	var menu = party.Party("/role") //菜单管理

	menu.Post("/addRole", hero.Handler(AddRole))       // 添加角色
	menu.Post("/updateRole", hero.Handler(UpdateRole)) // 更新角色
	menu.Post("/deleteRole", hero.Handler(DeleteRole)) // 删除角色

	menu.Post("/findRoleById", hero.Handler(FindRoleById))                                     // 根据id查询
	menu.Post("/findRoleByUserId", hero.Handler(FindRoleByUserId))                             // 根据id查询
	menu.Post("/findRoleList", hero.Handler(FindRoleList))                                     // 列表
	menu.Post("/findRolePages", hero.Handler(FindRolePages))                                   // 分页
	menu.Post("/findUnDistributeRoluesByUserId", hero.Handler(FindUnDistributeRoluesByUserId)) // 根据用户id查询用户未分配的角色

	menu.Post("/createRoleResource", hero.Handler(CreateRoleResource))             // 新增角色和菜单关系
	menu.Post("/getRoleResourceByRoleIds", hero.Handler(GetRoleResourceByRoleIds)) // 查询角色所拥有的菜单
}

//=============================   新增   =========================================
func AddRole(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.CreateRole(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   更新   =========================================
func UpdateRole(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.UpdateRole(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   删除   =========================================
func DeleteRole(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil || paramRVo.Ids == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res = userService.DeleteRole(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   获取菜单资源列表   =========================================
// 根据id查询详情
func FindRoleByUserId(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
	)

	err = ctx.ReadJSON(&paramRVo)
	if err != nil || paramRVo.Id == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindRoleById(paramRVo)
	responseHandle.Response(ctx, res)
}

// 根据用户id查询用户所拥有的角色
func FindRoleById(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
	)

	err = ctx.ReadJSON(&paramRVo)
	if err != nil || paramRVo.Id == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindRoleById(paramRVo)
	responseHandle.Response(ctx, res)
}

// 分页
func FindRolePages(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindRolePages(paramRVo)
	responseHandle.Response(ctx, res)
}

// 列表
func FindRoleList(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindRoleList(paramRVo)
	responseHandle.Response(ctx, res)
}

// 查询角色所拥有的菜单
func GetRoleResourceByRoleIds(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.GetRoleResourceByRoleIds(paramRVo)
	responseHandle.Response(ctx, res)
}

//=============================   根据用户id查询用户未分配的角色   =========================================
func FindUnDistributeRoluesByUserId(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindUnDistributeRoluesByUserId(paramRVo.UserId)
	responseHandle.Response(ctx, res)
}

//=============================   角色资源关联关系   =========================================
func CreateRoleResource(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.RoleResourceParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.CreateRoleResource(paramRVo)

	responseHandle.Response(ctx, res)
}
