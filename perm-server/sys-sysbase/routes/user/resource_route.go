package user

import (
	"../../../sys-common/models/baseModel"
	"../../../sys-common/supports"
	"../../../sys-common/supports/commonConst"
	"../../../sys-common/supports/responseHandle"
	"../../service/userService"
	userVO "../../vo/uservo"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
)

func ResoureHub(party iris.Party) {
	var menu = party.Party("/resoure") //菜单管理

	menu.Post("/addResource", hero.Handler(AddResource))       // 添加资源
	menu.Post("/updateResource", hero.Handler(UpdateResource)) // 更新资源
	menu.Post("/deleteResource", hero.Handler(DeleteResource)) // 删除资源

	menu.Post("/findResourceList", hero.Handler(FindResourceList))                             // 给角色添加权限
	menu.Post("/findResourceTree", hero.Handler(FindResourceTree))                             // 给角色添加权限
	menu.Post("/findResourceTreeAll", hero.Handler(FindResourceTreeAll))                       // 获取菜单树带根
	menu.Post("/findResourceTreeAllWithoutRoot", hero.Handler(FindResourceTreeAllWithoutRoot)) // 获取菜单树不带带根
	menu.Post("/findPagesResource", hero.Handler(FindPagesResource))                           // 给角色添加权限
}

//=============================   新增   =========================================
func AddResource(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.CreateResource(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   更新   =========================================
func UpdateResource(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.UpdateResource(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   删除   =========================================
func DeleteResource(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil || paramRVo.Ids == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res = userService.DeleteResource(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   获取菜单资源列表   =========================================
func FindResourceList(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("用户id为空", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, "用户id为空", nil)
		return
	}

	res = userService.FindResourceListByUser(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   获取用户所拥有菜单   =========================================
func FindResourceTree(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("用户id为空", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, "用户id为空", nil)
		return
	}

	res = userService.FindResourceTreeByUser(paramRVo)

	responseHandle.Response(ctx, res)

}

func FindPagesResource(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("用户id为空", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, "用户id为空", nil)
		return
	}

	tmpRes := userService.FindPagesResource(paramRVo)
	if !tmpRes.Code {
		ctx.Application().Logger().Errorf("用户id为空", "", err.Error())
		return
	}

	supports.Ok(ctx, tmpRes.Msg, tmpRes.Data)

	return
}

// 获取菜单树
func FindResourceTreeAll(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		subRtVo  []*userVO.ResourceTreeVo
		rtVo     *userVO.ResourceTreeVo
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("用户id为空", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, "用户id为空", nil)
		return
	}

	subRtVo, res := userService.FindResourceTreeAll(paramRVo)
	if res == nil {
		supports.Ok(ctx, "", "")
		return
	}

	if !res.Code {
		supports.Error(ctx, iris.StatusBadRequest, res.Msg, nil)
		return
	}

	rtVo = &userVO.ResourceTreeVo{
		Id:           "0",
		ParentId:     "-1",
		ResourceName: "总菜单",
		Subs:         subRtVo,
		Children:     subRtVo,
		Label:        "总菜单",
	}
	supports.Ok(ctx, res.Msg, rtVo)
}

func FindResourceTreeAllWithoutRoot(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(userVO.ResourceParamVo)
		subRtVo  []*userVO.ResourceTreeVo
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("用户id为空", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, "用户id为空", nil)
		return
	}

	subRtVo, res := userService.FindResourceTreeAll(paramRVo)
	res = &supports.SimResult{
		Code: res.Code,
		Msg:  res.Msg,
		Data: subRtVo,
	}
	responseHandle.Response(ctx, res)
}
