package uservo

import (
	"permissionManage/perm-server/sys-common/models/baseModel"
	"permissionManage/perm-server/sys-common/vo"
	"permissionManage/perm-server/sys-sysbase/models/user"
)

// 前端需要的数据结构
type DepartmentVO struct {
	*user.User
	Token string `json:"token"`
}

// 请求参数
type DepartmentParamVo struct {
	*baseModel.CurrentUser
	baseModel.BaseModel
	*vo.PageVo
	DepName  string `json:"depName"`
	DepCode  string `json:"depCode" form:"depCode"`
	DepLevel int64  `json:"depLevel" form:"depLevel"`
	DepOrder int64  `json:"depOrder" form:"depOrder"`
	ParentId string `json:"parentId" form:"parentId"`
	Token    string `json:"token"`

	Ids string `json:"ids"` //批量删除入参
}

// 列表数据结构
type DepartmentListVo struct {
	Id       string `json:"id"`
	DepName  string `json:"depName"`
	DepCode  string `json:"depCode" form:"depCode"`
	DepLevel int64  `json:"depLevel" form:"depLevel"`
	DepOrder int64  `json:"depOrder" form:"depOrder"`
	ParentId string `json:"parentId" form:"parentId"`
}

// 树结构数据
type DepartmentTreeVo struct {
	Id       string              `json:"id"`
	ParentId string              `json:"parentId"`
	Subs     []*DepartmentTreeVo `json:"subs"`
	Child    []*DepartmentTreeVo `json:"child"`
	Name     string              `json:"name"`
	DepLevel int64               `json:"depLevel" form:"depLevel"`
}

func DepartmentToDepartmentTreeVo(resource *user.Department) (rtVo *DepartmentTreeVo) {
	rtVo = new(DepartmentTreeVo)
	rtVo.Id = resource.Id
	rtVo.ParentId = resource.ParentId
	rtVo.Name = resource.DepName
	rtVo.DepLevel = resource.DepLevel
	//golog.Infof("%s", rtVo)
	return
}

func DepartmentToDepartmentTreeVos(resources []*user.Department) (rtVos []*DepartmentTreeVo) {
	for _, r := range resources {
		rtVo := DepartmentToDepartmentTreeVo(r)
		rtVos = append(rtVos, rtVo)
	}
	return
}

func DepartmentParamVoVoToDepartment(paramVo *DepartmentParamVo) (result *user.Department) {
	currentUser := paramVo.CurrentUser
	if paramVo.Id != "" {
		result = user.FindDepartmentById(paramVo.Id)
		result.UpdatorId = currentUser.UserAccount
		result.UpdatorName = currentUser.UserName
		result.UpdatorDepartmentId = currentUser.DepartmentId
		result.UpdatorDepartmentName = currentUser.DepartmentName
	} else {
		result = new(user.Department)
		result.CreatorId = currentUser.UserAccount
		result.CreatorName = currentUser.UserName
		result.CreatorDepartmentId = currentUser.DepartmentId
		result.CreatorDepartmentName = currentUser.DepartmentName
	}
	result.DepName = paramVo.DepName
	result.DepCode = paramVo.DepCode
	result.DepOrder = paramVo.DepOrder
	result.ParentId = paramVo.ParentId
	result.DepLevel = paramVo.DepLevel
	return
}

func DepartmentToResourceListVo(dep *user.Department) (depvo *DepartmentListVo) {
	depvo = new(DepartmentListVo)
	depvo.Id = dep.Id
	depvo.DepName = dep.DepName
	depvo.DepCode = dep.DepCode
	depvo.DepOrder = dep.DepOrder
	depvo.ParentId = dep.ParentId
	depvo.DepLevel = dep.DepLevel
	return
}

func DepartmentToDepartmentListVos(deps []*user.Department) (depVos []*DepartmentListVo) {
	for _, dep := range deps {
		depVo := DepartmentToResourceListVo(dep)
		depVos = append(depVos, depVo)
	}
	return
}
