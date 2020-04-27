package uservo

import (
	"permissionManage/perm-server/sys-common/models/baseModel"
	"permissionManage/perm-server/sys-common/vo"
	"permissionManage/perm-server/sys-sysbase/models/user"
)

// 前端需要的数据结构
type RoleVO struct {
	Id          string `json:"id"`
	PType       string `json:"pType"`
	RoleType    string `json:"roleType" form:"roleType"`
	RoleName    string `json:"roleName" form:"roleName"`
	RoleCode    string `json:"roleCode" form:"roleCode"`
	Description string `json:"description" form:"description"`

	Token string `json:"token"`
}

// 请求参数
type RoleParamVo struct {
	*baseModel.CurrentUser
	baseModel.BaseModel
	*vo.PageVo
	Id          string `json:"id"`
	PType       string `json:"pType"`
	RoleType    string `json:"roleType" form:"roleType"`
	RoleName    string `json:"roleName" form:"roleName"`
	RoleCode    string `json:"roleCode" form:"roleCode"`
	Description string `json:"description" form:"description"`
	Token       string `json:"token"`

	ResourceType string `json:"resourceType"`
	Ids          string `json:"ids"`    //批量删除入参
	UserId       string `json:"userId"` //用户id
}

// 列表数据结构
type RoleListVo struct {
	Id          string `json:"id"`
	PType       string `json:"pType"`
	RoleType    string `json:"roleType" form:"roleType"`
	RoleName    string `json:"roleName" form:"roleName"`
	RoleCode    string `json:"roleCode" form:"roleCode"`
	Description string `json:"description" form:"description"`
}

func RoleParamVoToRole(paramVo *RoleParamVo) (result *user.Role) {
	currentUser := paramVo.CurrentUser
	if paramVo.Id != "" {
		result = user.FindRoleById(paramVo.Id)
		result.UpdatorId = currentUser.UserAccount
		result.UpdatorName = currentUser.UserName
		result.UpdatorDepartmentId = currentUser.DepartmentId
		result.UpdatorDepartmentName = currentUser.DepartmentName
	} else {
		result = new(user.Role)
		result.CreatorId = currentUser.UserAccount
		result.CreatorName = currentUser.UserName
		result.CreatorDepartmentId = currentUser.DepartmentId
		result.CreatorDepartmentName = currentUser.DepartmentName
	}

	result.PType = paramVo.PType
	result.RoleType = "p"
	result.RoleName = paramVo.RoleName
	result.RoleCode = paramVo.RoleCode
	result.Description = paramVo.Description
	return
}

func RoleToRoleLListVo(rl *user.Role) (rlvo *RoleListVo) {
	rlvo = new(RoleListVo)
	rlvo.Id = rl.Id
	rlvo.PType = rl.PType
	rlvo.RoleType = rl.RoleType
	rlvo.RoleName = rl.RoleName
	rlvo.RoleCode = rl.RoleCode
	rlvo.Description = rl.Description
	return
}

func RoleToRoleListVos(rls []*user.Role) (rlvos []*RoleListVo) {
	for _, r := range rls {
		rtVo := RoleToRoleLListVo(r)
		rlvos = append(rlvos, rtVo)
	}
	return
}

func RoleToRoleVo(param *user.Role) (result *RoleListVo) {
	result = new(RoleListVo)
	result.Id = param.Id
	result.PType = param.PType
	result.RoleType = param.RoleType
	result.RoleName = param.RoleName
	result.RoleCode = param.RoleCode
	result.Description = param.Description
	return
}
