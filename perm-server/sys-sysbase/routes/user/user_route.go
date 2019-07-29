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

func UserHub(party iris.Party) {
	var u = party.Party("/user")
	//p.Use(middleware.BasicAuth)s
	u.Post("/registe", hero.Handler(Registe)) // 注册
	u.Post("/login", hero.Handler(Login))     // 登录
	u.Post("/logout", hero.Handler(LogOut)) // 登出

	u.Post("/addUser", hero.Handler(AddUser))       // 新增用户
	u.Post("/updateUser", hero.Handler(UpdateUser)) // 更新用户
	u.Post("/resetPass", hero.Handler(ResetPass)) // 重置密码
	u.Post("/deleteUesr", hero.Handler(DeleteUesr)) // 删除用户

	u.Post("/findUserById", hero.Handler(FindUserById))                   // 根据id查询
	u.Post("/findPagesUser", hero.Handler(FindPagesUser))                 // 分页
	u.Post("/findUserRoleSByUserId", hero.Handler(FindUserRoleSByUserId)) // 根据用户id获取用户所拥有的角色
	u.Post("/handlePermission", hero.Handler(HandlePermission))           // 授权撤权处理
	//u.Delete("/{uids:string}", hero.Handler(userService.DeleteUsers)) // 删除用户
}

func Registe(ctx iris.Context) {

}

func Login(ctx iris.Context) {
	var (
		err      error
		paramUVo = new(uservo.UserParamVO)
		uVo      *uservo.UserParamVO
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramUVo); err != nil {
		ctx.Application().Logger().Errorf("用户[%s]登录失败。%s", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.LoginFailur, nil)
		return
	}

	uVo, res = userService.Login(paramUVo)
	if !res.Code {
		supports.Error(ctx, iris.StatusBadRequest, res.Msg, nil)
	}

	supports.Ok(ctx, res.Msg, uVo)
}

func LogOut(ctx iris.Context) {
	supports.Ok(ctx, supports.LogoutSuccess, nil)
	return
}

//=============================   新增   =========================================
func AddUser(ctx iris.Context) {
	var (
		err      error
		paramUVo = new(uservo.UserParamVO)
		res      *supports.SimResult
	)
	if err = ctx.ReadJSON(&paramUVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramUVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.CreateUser(paramUVo)
	if !res.Code {
		supports.Error(ctx, iris.StatusBadRequest, res.Msg, nil)
	}

	supports.Ok(ctx, res.Msg, nil)
}

//=============================   更新   =========================================
func UpdateUser(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserParamVO)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.UpdateUser(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   重置密码   =========================================
func ResetPass(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserParamVO)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res = userService.ResetPass(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   删除   =========================================
func DeleteUesr(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserParamVO)
		res      *supports.SimResult
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil || paramRVo.Ids == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res = userService.DeleteUser(paramRVo)

	responseHandle.Response(ctx, res)
}

//=============================   查询   =========================================
// 根据id查询详情
func FindUserById(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserParamVO)
	)

	err = ctx.ReadJSON(&paramRVo)
	if err != nil || paramRVo.Id == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindUserById(paramRVo)
	responseHandle.Response(ctx, res)
}

//=============================   分页查询用户   =========================================
func FindPagesUser(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserParamVO)
	)

	if err = ctx.ReadJSON(&paramRVo); err != nil {
		ctx.Application().Logger().Errorf("", "", err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	tmpRes := userService.FindPagesUser(paramRVo)
	if !tmpRes.Code {
		ctx.Application().Logger().Errorf("json", supports.QUERY_FAIL, err.Error())
		return
	}

	supports.Ok(ctx, tmpRes.Msg, tmpRes.Data)

	return
}

// 根据id查询详情
func FindUserRoleSByUserId(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserParamVO)
	)

	err = ctx.ReadJSON(&paramRVo)
	if err != nil || paramRVo.Id == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}

	res := userService.FindUserRoleSByUserId(paramRVo)

	responseHandle.Response(ctx, res)
}

// 根据id查询详情
func HandlePermission(ctx iris.Context) {
	var (
		err      error
		paramRVo = new(uservo.UserRoleParamVo)
	)

	err = ctx.ReadJSON(&paramRVo)
	if err != nil || paramRVo.UserId == "" {
		ctx.Application().Logger().Errorf("", supports.PARAM_IS_EMPTY, err.Error())
		supports.Error(ctx, iris.StatusBadRequest, supports.PARAM_IS_EMPTY, nil)
		return
	}
	paramRVo.CurrentUser = baseModel.GetCurrentUset(ctx.Params().Get(commonConst.CURRENT_USER_KEY))

	res := userService.HandlePermission(paramRVo)

	responseHandle.Response(ctx, res)
}
