package userService

import (
	"../../../sys-common/db"
	"../../../sys-common/supports"
	"../../../sys-common/vo"
	"../../models/user"
	"../../vo/uservo"
	"fmt"
	"strings"
)

//=============================   新增   =========================================
func CreateRole(paramVo *uservo.RoleParamVo) (res *supports.SimResult) {
	role := uservo.RoleParamVoToRole(paramVo)

	_, err := user.CreateRole(role)
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
func UpdateRole(paramVo *uservo.RoleParamVo) (res *supports.SimResult) {
	eModel := uservo.RoleParamVoToRole(paramVo)
	_, err := user.UpdateRole(eModel)
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
func DeleteRole(paramVo *uservo.RoleParamVo) (res *supports.SimResult) {
	idsStr := paramVo.Ids
	_, err := user.DeleteRoleByIds(strings.Split(idsStr, ","))
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
func FindRoleById(roleParam *uservo.RoleParamVo) (res *supports.SimResult) {
	res = &supports.SimResult{
		Code: true,
		Data: uservo.RoleToRoleVo(user.FindRoleById(roleParam.Id)),
	}
	return
}
func FindRoleList(roleParam *uservo.RoleParamVo) (res *supports.SimResult) {
	e := db.MasterEngine()
	result := make([]*user.Role, 0)

	sql := fmt.Sprintf(" SELECT * FROM cd_sys_user_role ur " +
			" WHERE ur.is_deleted = false")

	err := e.SQL(sql).Find(&result)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  "",
			Data: nil,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Data: uservo.RoleToRoleListVos(result),
	}

	return
}

// 分页
func FindRolePages(paramVo *uservo.RoleParamVo) (res *supports.SimResult) {
	e := db.MasterEngine()
	roleListVos := make([]*uservo.RoleListVo, 0)
	roles := make([]*user.Role, 0)

	pageParam := vo.GetPageParam(paramVo.PageVo)

	sql := " is_deleted = false "
	if paramVo.RoleName != "" {
		sql += " and role_name like '%" + paramVo.RoleName + "%'"
	}
	s := e.Where(sql).Limit(pageParam.PageSize, pageParam.Start)
	count, err := s.FindAndCount(&roles)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
	}

	roleListVos = uservo.RoleToRoleListVos(roles)

	pageResult := &vo.PageResult{
		CurrentPage: pageParam.CurrentPage,
		PageSize:    pageParam.PageSize,
		Total:       count,
		Data:        roleListVos,
	}
	res = &supports.SimResult{
		Code: true,
		Msg:  "",
		Data: pageResult,
	}
	return
}

// 查询角色所拥有的菜单
func GetRoleResourceByRoleIds(paramVo *uservo.RoleParamVo) (res *supports.SimResult) {
	e := db.MasterEngine()
	sql := fmt.Sprintf("SELECT" +
			" ur.id" +
			" FROM cd_sys_user_resource ur" +
			" LEFT JOIN cd_sys_user_role_resource urr ON ur.id = urr.resource_id" +
			" WHERE ur.is_deleted = 0" +
			" AND ur.resource_type = '" + paramVo.ResourceType + "'" +
			" AND urr.role_id = '" + paramVo.Id + "'")
	tmpRes, err := e.QueryString(sql)

	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
	}
	resArry := make([]string, 0)
	for _, tmp := range tmpRes {
		resArry = append(resArry, tmp["id"])
	}

	res = &supports.SimResult{
		Code: true,
		Data: resArry,
	}
	return
}

//=============================   角色资源关联关系   =========================================
func FindRoluesByUserId(userId string) (res *supports.SimResult) {
	e := db.MasterEngine()
	result := make([]*user.Role, 0)
	sql := "SELECT * FROM cd_sys_user_role ur WHERE ur.id IN(SELECT uur.role_id FROM cd_sys_user_user_role uur WHERE uur.user_id = '" +
		"" + userId + "') AND ur.is_deleted = FALSE"
	err := e.SQL(sql).Find(&result)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  "",
			Data: nil,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Data: uservo.RoleToRoleListVos(result),
	}

	return
}

// 根据用户id查询用户未分配的角色
func FindUnDistributeRoluesByUserId(userId string) (res *supports.SimResult) {
	e := db.MasterEngine()
	result := make([]*user.Role, 0)
	sql := "SELECT * FROM cd_sys_user_role ur WHERE ur.id NOT IN(SELECT uur.role_id FROM cd_sys_user_user_role uur WHERE uur.user_id = '" +
			"" + userId + "') AND ur.is_deleted = FALSE"
	err := e.SQL(sql).Find(&result)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  "",
			Data: nil,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Data: uservo.RoleToRoleListVos(result),
	}

	return
}

func CreateRoleResource(paramVo *uservo.RoleResourceParamVo) (res *supports.SimResult) {
	roleResource := uservo.RoleResourceParamVoToRoleResource(paramVo)

	// 删除现有的关系
	user.RemoveRoleResource(paramVo.RoleId)
	// 添加新的关系
	_, err := user.CreateRelationRoleResource(roleResource...)
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
