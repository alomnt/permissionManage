package userService

import (
	"../../../sys-common/db"
	"../../../sys-common/middleware/jwts"
	"../../../sys-common/models/baseModel"
	"../../../sys-common/supports"
	"../../../sys-common/utils"
	"../../../sys-common/vo"
	"../../models/user"
	"../../vo/uservo"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"strings"
	"fmt"
	"strconv"
)

//=============================   新增   =========================================
func CreateUser(paramVo *uservo.UserParamVO) (res *supports.SimResult) {
	eModel := uservo.UserParamVOToUser(paramVo)

	maxUserNumber, numErr := user.GetMaxUserNumber()
	if numErr != nil {
		numErr.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
		return
	}
	eModel.UserNumber = strconv.FormatInt(maxUserNumber+1, 10)

	_, err := user.CreateUser(eModel)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.ADD_SUCCESS,
	}

	return
}

//=============================   更新   =========================================
func UpdateUser(paramVo *uservo.UserParamVO) (res *supports.SimResult) {
	eModel := uservo.UserParamVOToUser(paramVo)
	_, err := user.UpdateUser(eModel)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.UPDATE_FAIL,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.UPDATE_SUCCESS,
	}

	return
}

//=============================   重置密码   =========================================
func ResetPass(paramVo *uservo.UserParamVO) (res *supports.SimResult) {
	eModel := uservo.ResetPassVOToUser(paramVo)
	_, err := user.UpdateUser(eModel)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.UPDATE_FAIL,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.UPDATE_SUCCESS,
	}

	return
}

//=============================   删除   =========================================
func DeleteUser(paramVo *uservo.UserParamVO) (res *supports.SimResult) {
	idsStr := paramVo.Ids
	_, err := user.DeleteUserByIds(strings.Split(idsStr, ","))
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.DELETE_FAIL,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.DELETE_SUCCESS,
	}

	return
}

//=============================   查询   =========================================
// 根据id查询
func FindUserById(paramVo *uservo.UserParamVO) (res *supports.SimResult) {
	res = &supports.SimResult{
		Code: true,
		Data: uservo.UserToUserSimpleVO(user.FindUserById(paramVo.Id)),
	}
	return
}

// 登出
func LoginOut(ctx iris.Context) {
	supports.Ok(ctx, supports.LoginSuccess, nil)
	return
}

// 登录
func Login(paramUVo *uservo.UserParamVO) (uSVO *uservo.UserParamVO, res *supports.SimResult) {
	var (
		err         error
		mUser       = new(user.User) // 查询数据后填充了数据的user
		currentUser *baseModel.CurrentUser
		ckPassword  bool
		token       string
	)

	mUser.UserAccount = paramUVo.UserAccount
	has, err := user.GetUserByUsername(mUser)
	currentUser, errCurr := baseModel.InitCurrentUserByUserAccount(paramUVo.UserAccount)
	if err != nil && errCurr != nil {
		golog.Println("用户[%s]登录失败。%s", paramUVo.UserName, err.Error())
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.LoginFailur,
		}
		return
	}

	if !has { // 用户名不正确
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.UsernameFailur,
		}
		return
	}

	ckPassword = utils.CheckPWD(paramUVo.UserPassword, mUser.UserPassword)
	if !ckPassword {
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.PasswordFailur,
		}
		return
	}

	token, err = jwts.GenerateToken(mUser, currentUser)
	golog.Infof("用户[%s], 登录生成token [%s]", mUser.UserName, token)
	if err != nil {
		golog.Println("用户[%s]登录，生成token出错。%s", paramUVo.UserName, err.Error())
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.TokenCreateFailur,
		}
		return
	}

	uSVO = uservo.UserToUserParamVO(token, mUser)
	res = &supports.SimResult{
		Code: true,
		Msg:  supports.LoginSuccess,
	}
	return
}

//========================== 分页 ========================
func FindPagesUser(paramVo *uservo.UserParamVO) (res *supports.SimResult) {
	e := db.MasterEngine()
	userListVos := make([]*uservo.UserListVO, 0)
	userDepartments := make([]*user.UserDepartment, 0)

	pageParam := vo.GetPageParam(paramVo.PageVo)

	sql := " cd_sys_user_user.is_deleted = false "
	if paramVo.UserName != "" {
		sql += " and cd_sys_user_user.user_name like '%" + paramVo.UserName + "%'"
	}
	s := e.Join("LEFT", "cd_sys_user_department",
		"cd_sys_user_department.id = cd_sys_user_user.department_id").Where(sql).Limit(pageParam.PageSize, pageParam.Start)
	count, err := s.FindAndCount(&userDepartments)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
	}

	userListVos = uservo.UserToUserListVOs(userDepartments)

	pageResult := &vo.PageResult{
		CurrentPage: pageParam.CurrentPage,
		PageSize:    pageParam.PageSize,
		Total:       count,
		Data:        userListVos,
	}
	res = &supports.SimResult{
		Code: true,
		Msg:  "",
		Data: pageResult,
	}
	golog.Println(count, err)
	return
}

func FindUserList(paramVo *uservo.UserParamVO) (rs []*user.User, err error) {
	e := db.MasterEngine()
	result := make([]*user.User, 0)

	sql := fmt.Sprintf(" SELECT * FROM cd_sys_user_user uu " +
			" WHERE uu.is_deleted = false")
	if paramVo.DepartmentId != "" {
		sql = sql + " AND uu.department_id = '" + paramVo.DepartmentId + "'"
	}

	err = e.SQL(sql).Find(&result)

	return result, err
}

func FindUserRoleSByUserId(paramVo *uservo.UserParamVO) (res *supports.SimResult) {

	return FindRoluesByUserId(paramVo.Id)
}

//========================== 	用户角色		 ========================
func HandlePermission(paramVo *uservo.UserRoleParamVo) (res *supports.SimResult) {
	checkType := paramVo.CheckType
	if checkType == "1" {
		return CreateUserRole(paramVo)
	} else if checkType == "0" {

		return RemoveUserRoles(paramVo)
	}

	return
}

// 授权
func CreateUserRole(paramVo *uservo.UserRoleParamVo) (res *supports.SimResult) {
	userRoles := uservo.UserRoleParamVoToUserRole(paramVo)

	// 添加新的关系
	_, err := user.CreateRelationUserRole(userRoles...)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.ADD_SUCCESS,
	}

	return
}

// 撤权
func RemoveUserRoles(paramVo *uservo.UserRoleParamVo) (res *supports.SimResult) {
	//roleIdArry := strings.Split(paramVo.RoleIds, ",")
	roleIdArry := paramVo.RoleIds
	_, err := user.RemoveUserRole(paramVo.UserId, roleIdArry)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.DELETE_FAIL,
		}
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.DELETE_SUCCESS,
	}

	return
}
