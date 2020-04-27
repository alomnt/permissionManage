package uservo

import (
	"permissionManage/perm-server/sys-common/models/baseModel"
	"permissionManage/perm-server/sys-common/utils"
	"permissionManage/perm-server/sys-common/vo"
	"permissionManage/perm-server/sys-common/vo/baseVo"
	"permissionManage/perm-server/sys-sysbase/models/user"
)

// 前端需要的数据结构
type UserParamVO struct {
	*baseModel.CurrentUser
	baseModel.BaseModel
	*vo.PageVo
	Id           string `json:"id" form:"id"`
	UserNumber   string `json:"userNumber" form:"userNumber"`
	UserAccount  string `json:"userAccount" form:"userAccount"`
	UserPassword string `json:"userPassword" form:"userPassword"`
	UserName     string `json:"userName" form:"userName"`
	Sex          string `json:"sex" form:"sex"`
	DepartmentId string `json:"departmentId" form:"departmentId"`
	Mobile       string `json:"mobile" form:"mobile"`
	Email        string `json:"email" form:"email"`
	Address      string `json:"address" form:"address"`
	Token        string `json:"token"`

	Ids string `json:"ids"` //批量删除入参
}

// 前端需要的数据结构
type UserVO struct {
	baseVo.BaseVO
	Id           string `json:"id" form:"id"`
	UserNumber   string `json:"userNumber" form:"userNumber"`
	UserAccount  string `json:"userAccount" form:"userAccount"`
	UserPassword string `json:"userPassword" form:"userPassword"`
	UserName     string `json:"userName" form:"userName"`
	Sex          string `json:"sex" form:"sex"`
	DepartmentId string `json:"departmentId" form:"departmentId"`
	Mobile       string `json:"mobile" form:"mobile"`
	Email        string `json:"email" form:"email"`
	Address      string `json:"address" form:"address"`
}

// 前端需要的数据结构
type UserSimpleVO struct {
	Id           string `json:"id" form:"id"`
	UserNumber   string `json:"userNumber" form:"userNumber"`
	UserAccount  string `json:"userAccount" form:"userAccount"`
	UserPassword string `json:"userPassword" form:"userPassword"`
	UserName     string `json:"userName" form:"userName"`
	Sex          string `json:"sex" form:"sex"`
	DepartmentId string `json:"departmentId" form:"departmentId"`
	Mobile       string `json:"mobile" form:"mobile"`
	Email        string `json:"email" form:"email"`
	Address      string `json:"address" form:"address"`
	Token        string `json:"token"`
}

// 前端需要的数据结构
type UserListVO struct {
	Id             string `json:"id" form:"id"`
	UserNumber     string `json:"userNumber" form:"userNumber"`
	UserAccount    string `json:"userAccount" form:"userAccount"`
	UserName       string `json:"userName" form:"userName"`
	Sex            string `json:"sex" form:"sex"`
	DepartmentId   string `json:"departmentId" form:"departmentId"`
	DepartmentName string `json:"departmentName" form:"departmentName"`
	Mobile         string `json:"mobile" form:"mobile"`
	Email          string `json:"email" form:"email"`
	Address        string `json:"address" form:"address"`
}

// 携带token
func UserToUserSimpleVO(user *user.User) (uSVO *UserSimpleVO) {
	uSVO = &UserSimpleVO{
		Id:           user.Id,
		UserNumber:   user.UserNumber,
		UserAccount:  user.UserAccount,
		UserPassword: user.UserPassword,
		UserName:     user.UserName,
		Sex:          user.Sex,
		DepartmentId: user.DepartmentId,
		Mobile:       user.Mobile,
		Email:        user.Email,
		Address:      user.Address,
	}
	return
}
func UserToUserSimpleVOs(userList []*user.User) (uSVOList []*UserSimpleVO) {
	for _, user := range userList {
		uSVOList = append(uSVOList, UserToUserSimpleVO(user))
	}
	return
}

func UserToUserListVO(userDepartment *user.UserDepartment) (uSVO *UserListVO) {
	uSVO = new(UserListVO)
	uSVO.Id = userDepartment.User.Id
	uSVO.UserNumber = userDepartment.User.UserNumber
	uSVO.UserAccount = userDepartment.User.UserAccount
	uSVO.UserName = userDepartment.User.UserName
	if userDepartment.User.Sex == "0" {
		uSVO.Sex = "男"
	} else {
		uSVO.Sex = "女"
	}
	uSVO.DepartmentId = userDepartment.Department.Id
	uSVO.DepartmentName = userDepartment.Department.DepName
	uSVO.Mobile = userDepartment.Mobile
	uSVO.Email = userDepartment.Email
	uSVO.Address = userDepartment.Address
	return
}
func UserToUserListVOs(userDepartments []*user.UserDepartment) (uSVOList []*UserListVO) {
	for _, userDepartment := range userDepartments {
		uSVOList = append(uSVOList, UserToUserListVO(userDepartment))
	}
	return
}

// 前端传递数据转化为数据库数据类型
func UserSimpleVOToUser(uSVO *UserSimpleVO) (result *user.User) {
	result = new(user.User)
	result.UserNumber = uSVO.UserNumber
	result.UserAccount = uSVO.UserAccount
	result.UserPassword = uSVO.UserPassword
	result.UserName = uSVO.UserName
	result.Sex = uSVO.Sex
	result.DepartmentId = uSVO.DepartmentId
	result.Mobile = uSVO.Mobile
	result.Email = uSVO.Email
	result.Address = uSVO.Address
	return
}

// 携带token
func UserToUserParamVO(token string, user *user.User) (uSVO *UserParamVO) {
	uSVO = &UserParamVO{
		Id:           user.Id,
		UserNumber:   user.UserNumber,
		UserAccount:  user.UserAccount,
		UserPassword: user.UserPassword,
		UserName:     user.UserName,
		Sex:          user.Sex,
		DepartmentId: user.DepartmentId,
		Mobile:       user.Mobile,
		Email:        user.Email,
		Address:      user.Address,
		Token:        token,
	}
	return
}

// 前端传递数据转化为数据库数据类型
func UserParamVOToUser(upVO *UserParamVO) (result *user.User) {
	currentUser := upVO.CurrentUser
	if upVO.Id != "" {
		result = user.FindUserById(upVO.Id)
		result.UpdatorId = currentUser.UserAccount
		result.UpdatorName = currentUser.UserName
		result.UpdatorDepartmentId = currentUser.DepartmentId
		result.UpdatorDepartmentName = currentUser.DepartmentName
	} else {
		result = new(user.User)
		result.CreatorId = currentUser.UserAccount
		result.CreatorName = currentUser.UserName
		result.CreatorDepartmentId = currentUser.DepartmentId
		result.CreatorDepartmentName = currentUser.DepartmentName
	}
	result.UserNumber = upVO.UserNumber
	result.UserAccount = upVO.UserAccount
	result.UserPassword = upVO.UserPassword
	result.UserName = upVO.UserName
	result.Sex = upVO.Sex
	result.DepartmentId = upVO.DepartmentId
	result.Mobile = upVO.Mobile
	result.Email = upVO.Email
	result.Address = upVO.Address
	return
}

// 前端传递数据转化为数据库数据类型
func ResetPassVOToUser(upVO *UserParamVO) (result *user.User) {
	currentUser := upVO.CurrentUser
	result = user.FindUserById(upVO.Id)
	result.UpdatorId = currentUser.UserAccount
	result.UpdatorName = currentUser.UserName
	result.UpdatorDepartmentId = currentUser.DepartmentId
	result.UpdatorDepartmentName = currentUser.DepartmentName

	result.UserPassword = utils.AESEncrypt([]byte(upVO.UserPassword))
	return
}
