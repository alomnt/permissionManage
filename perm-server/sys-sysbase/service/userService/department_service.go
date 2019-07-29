package userService

import (
	"../../../sys-common/db"
	"../../../sys-common/supports"
	"../../../sys-common/vo"
	"../../models/user"
	"../../vo/uservo"
	"fmt"
	"github.com/kataras/golog"
	"strings"
)

//=============================   新增   =========================================
func CreateDepartment(paramVo *uservo.DepartmentParamVo) (res *supports.SimResult) {
	mod := uservo.DepartmentParamVoVoToDepartment(paramVo)

	maxOrder, orderErr := user.GetDepartmentMaxOrderByParentId(paramVo.ParentId)
	if orderErr != nil {
		orderErr.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
		return
	}
	mod.DepOrder = maxOrder + 1

	if paramVo.ParentId == "-1" {
		mod.DepLevel = 1
	} else {
		level, levelErr := user.GetDepartmentParentLevelByParentId(paramVo.ParentId)
		if levelErr != nil {
			levelErr.Error()
			res = &supports.SimResult{
				Code: false,
				Msg:  supports.ADD_FAIL,
			}
			return
		}
		mod.DepLevel = level + 1
	}

	_, err := user.CreateDepartment(mod)
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
func UpdateDepartment(paramVo *uservo.DepartmentParamVo) (res *supports.SimResult) {
	eModel := uservo.DepartmentParamVoVoToDepartment(paramVo)
	_, err := user.UpdateDepartment(eModel)
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

//=============================   确认是否能删除   =========================================
func IsDepartmentCanDelete(paramVo *uservo.DepartmentParamVo) (res *supports.SimResult) {
	UserParamVo := &uservo.UserParamVO{
		DepartmentId: paramVo.Id,
	}
	res = new(supports.SimResult)
	users, err := FindUserList(UserParamVo)
	if err != nil {
		err.Error()
		res.Code = false
		res.Msg = supports.DELETE_FAIL
		res.Data = false
		return
	}
	if users == nil || len(users) == 0 {
		res.Code = true
		res.Data = true

		return
	}

	res.Code = false
	res.Msg = ""
	res.Data = false

	return
}

//=============================   删除   =========================================
func DeleteDepartment(paramVo *uservo.DepartmentParamVo) (res *supports.SimResult) {
	idsStr := paramVo.Ids
	_, err := user.DeleteDepartmentByIds(strings.Split(idsStr, ","))
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
func FindDepartmentList(paramVo *uservo.DepartmentParamVo) (res *supports.SimResult) {
	e := db.MasterEngine()
	tmpRes := make([]*user.Department, 0)

	sql := fmt.Sprintf(" SELECT * FROM cd_sys_user_department ud " +
			" WHERE ud.is_deleted = false")

	err := e.SQL(sql).Find(&tmpRes)
	if err != nil {
		res = &supports.SimResult{
			Code: false,
			Msg:  "",
			Data: tmpRes,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Data: uservo.DepartmentToDepartmentListVos(tmpRes),
	}

	return
}

func FindDepartments(paramVo *uservo.DepartmentParamVo) (rs []*user.Department, err error) {
	e := db.MasterEngine()
	result := make([]*user.Department, 0)

	sql := fmt.Sprintf(" SELECT * FROM cd_sys_user_department ud " +
			" WHERE ud.is_deleted = false")

	err = e.SQL(sql).Find(&result)

	return result, err
}

// 分页
func FindDepartmentPages(paramVo *uservo.DepartmentParamVo) (res *supports.SimResult) {
	e := db.MasterEngine()
	listVos := make([]*uservo.DepartmentListVo, 0)
	pageResTmp := make([]*user.Department, 0)

	pageParam := vo.GetPageParam(paramVo.PageVo)

	sql := " is_deleted = false "
	if paramVo.DepName != "" {
		sql += " and dep_name like '%" + paramVo.DepName + "%'"
	}
	s := e.Where(sql).Limit(pageParam.PageSize, pageParam.Start)
	count, err := s.FindAndCount(&pageResTmp)
	if err != nil {
		golog.Error(err)
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.QUERY_FAIL,
		}
		return
	}

	// 数据库查询出的数据类型转化为页面展示的数据类型
	listVos = uservo.DepartmentToDepartmentListVos(pageResTmp)

	pageResult := &vo.PageResult{
		CurrentPage: pageParam.CurrentPage,
		PageSize:    pageParam.PageSize,
		Total:       count,
		Data:        listVos,
	}
	res = &supports.SimResult{
		Code: true,
		Msg:  "",
		Data: pageResult,
	}
	return
}

//=============================   获取部门树   =========================================
func FindDepartmentTreeAll(resourceParam *uservo.DepartmentParamVo) (res *supports.SimResult) {
	var (
		err            error
		departmentList []*user.Department
	)

	departmentList, err = FindDepartments(resourceParam)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.QUERY_FAIL,
		}
		return
	}
	tmpRtVO := uservo.DepartmentToDepartmentTreeVos(departmentList)
	rtVO := buildDepartmentTree(tmpRtVO)

	res = &supports.SimResult{
		Code: true,
		Msg:  "",
		Data: rtVO,
	}
	return
}

func buildDepartmentTree(departmentTreeVo []*uservo.DepartmentTreeVo) (dtVO []*uservo.DepartmentTreeVo) {
	for _, dt := range departmentTreeVo {
		subs := getSubDepartmentTree(dt.Id, departmentTreeVo)
		if subs == nil {
			if dt.ParentId == "-1" {
				dtVO = append(dtVO, dt)
			}
			continue
		}
		dt.Subs = subs
		dt.Child = subs
		buildDepartmentTree(subs)
		if dt.ParentId == "-1" {
			dtVO = append(dtVO, dt)
		}
	}
	return
}
func getSubDepartmentTree(parentId string, departmentTreeVoAll []*uservo.DepartmentTreeVo) (dtVO []*uservo.DepartmentTreeVo) {
	for _, dt := range departmentTreeVoAll {
		if dt.ParentId == parentId {
			dtVO = append(dtVO, dt)
		}
	}

	return
}
