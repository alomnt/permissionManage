package user

import (
	"../../../sys-common/models/baseModel"
	"../../../sys-common/supports"
	"../../../sys-common/supports/commonConst"
	"../../../sys-common/supports/responseHandle"
	"../../service/userService"
	"../../vo/uservo"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func DepartmentHub(party iris.Party) {
	var menu = party.Party("/department") //菜单管理

	menu.Post("/addDepartment", hero.Handler(AddDepartment))                 // 添加部门
	menu.Post("/updateDepartment", hero.Handler(UpdateDepartment))           // 更新部门
	menu.Post("/isDepartmentCanDelete", hero.Handler(IsDepartmentCanDelete)) // 删除部门
	menu.Post("/deleteDepartment", hero.Handler(DeleteDepartment))           // 删除部门

	menu.Post("/findDepartmentList", hero.Handler(FindDepartmentList))       // 列表
	menu.Post("/findDepartmentPages", hero.Handler(FindDepartmentPages))     // 分页
	menu.Post("/findDepartmentTreeAll", hero.Handler(FindDepartmentTreeAll)) // 获取菜单树带根
}

//=============================   新增   =========================================
func AddDepartment(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.CreateDepartment(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   更新   =========================================
func UpdateDepartment(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.UpdateDepartment(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   确认是否能删除   =========================================
func IsDepartmentCanDelete(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil || paramRVo.Id == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res = userService.IsDepartmentCanDelete(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   删除   =========================================
func DeleteDepartment(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil || paramRVo.Ids == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res = userService.DeleteDepartment(paramRVo)

	responseHandle.Response(ctx, res)
}

// 获取菜单资源列表
func FindDepartmentPages(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindDepartmentPages(paramRVo)
	if res == nil {
		supports.Ok(ctx, "", "")
		return
	}

	if !res.Code {
		supports.Error(ctx, iris.StatusInternalServerError, res.Msg, nil)
		return
	}

	supports.Ok(ctx, res.Msg, res.Data)
}

func FindDepartmentList(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	tmpRes := userService.FindDepartmentList(paramRVo)
	if !tmpRes.Code {
		ctx.Application().Logger().Errorf("", tmpRes.Msg, err.Error())
		supports.Error(ctx, iris.StatusInternalServerError, tmpRes.Msg, nil)
		return
	}

	supports.Ok(ctx, tmpRes.Msg, tmpRes.Data)

	return
}

// 获取菜单树
func FindDepartmentTreeAll(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.DepartmentParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindDepartmentTreeAll(paramRVo)
	responseHandle.Response(ctx, res)
}
