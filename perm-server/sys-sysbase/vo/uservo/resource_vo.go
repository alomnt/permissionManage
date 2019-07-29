package uservo

import (
	"../../../sys-common/models/baseModel"
	"../../../sys-common/vo"
	"../../models/user"
	"time"
)

// 前端需要的数据结构
type ResourceVO struct {
	Id                    string    `json:"id"`
	CreateTime            time.Time `json:"createTime"`
	CreatorDepartmentId   string    `json:"creatorDepartmentId"`
	CreatorDepartmentName string    `json:"creatorDepartmentName"`
	CreatorId             string    `json:"creatorId"`
	CreatorName           string    `json:"creatorName"`
	IsDeleted             int64     `json:"isDeleted"`
	UpdateTime            time.Time `json:"updateTime"`
	UpdatorDepartmentId   string    `json:"updatorDepartmentId"`
	UpdatorDepartmentName string    `json:"updatorDepartmentName"`
	UpdatorId             string    `json:"updatorId"`
	UpdatorName           string    `json:"updatorName"`
	Version               int64     `json:"version"`

	ResourceType  string `json:"resourceType"`
	ResourceName  string `json:"resourceName"`
	ResourceLevel int64  `json:"resourceLevel"`
	ResourceOrder int64  `json:"resourceOrder"`
	ResourceSate  int64  `json:"resourceSate"`
	ParentId      string `json:"parentId"`
	Url           string `json:"url"`
	Icon          string `json:"icon"`
	Description   string `json:"description"`
}

// 请求参数
type ResourceParamVo struct {
	*baseModel.CurrentUser
	baseModel.BaseModel
	*vo.PageVo
	Id            string `json:"id"`
	UserId        string `json:"userId"`
	ResourceType  string `json:"resourceType"`
	ResourceName  string `json:"resourceName"`
	ResourceLevel int64  `json:"resourceLevel"`
	ResourceOrder int64  `json:"resourceOrder"`
	ResourceSate  int64  `json:"resourceSate"`
	ParentId      string `json:"parentId"`
	Url           string `json:"url"`
	Icon          string `json:"icon"`
	Description   string `json:"description"`
	Token         string `json:"token"`

	Ids string `json:"ids"` //批量删除入参
}

// 树结构数据
type ResourceTreeVo struct {
	Id           string            `json:"id"`
	ParentId     string            `json:"parentId"`
	ResourceName string            `json:"title"`
	Url          string            `json:"index"`
	Icon         string            `json:"icon"`
	Subs         []*ResourceTreeVo `json:"subs"`
	Children     []*ResourceTreeVo `json:"children"`
	Label        string            `json:"label"`
}

// 列表数据结构
type ResourceListVo struct {
	Id            string `json:"id"`
	ResourceType  string `json:"resourceType"`
	ResourceName  string `json:"resourceName"`
	ResourceLevel int64  `json:"resourceLevel"`
	ResourceOrder int64  `json:"resourceOrder"`
	ResourceSate  int64  `json:"resourceSate"`
	ParentId      string `json:"parentId"`
	Url           string `json:"url"`
	Icon          string `json:"icon"`
	Description   string `json:"description"`
}

// 用户列表，不带token
func ResourceToResourceTreeVo(resource *user.Resource) (rtVo *ResourceTreeVo) {
	rtVo = new(ResourceTreeVo)
	rtVo.Id = resource.Id
	rtVo.ParentId = resource.ParentId
	rtVo.ResourceName = resource.ResourceName
	rtVo.Url = resource.Url
	rtVo.Icon = resource.Icon
	rtVo.Label = resource.ResourceName
	//golog.Infof("%s", rtVo)
	return
}

func ResourceToResourceTreeVos(resources []*user.Resource) (rtVos []*ResourceTreeVo) {
	for _, r := range resources {
		rtVo := ResourceToResourceTreeVo(r)
		rtVos = append(rtVos, rtVo)
	}
	return
}

func ResourceToResourceListVo(resource *user.Resource) (rtVo *ResourceListVo) {
	rtVo = new(ResourceListVo)
	rtVo.Id = resource.Id
	rtVo.ResourceType = resource.ResourceType
	rtVo.ResourceName = resource.ResourceName
	rtVo.ResourceLevel = resource.ResourceLevel
	rtVo.ResourceLevel = resource.ResourceLevel
	rtVo.ResourceOrder = resource.ResourceOrder
	rtVo.ResourceSate = resource.ResourceSate
	rtVo.ParentId = resource.ParentId
	rtVo.Url = resource.Url
	rtVo.Icon = resource.Icon
	rtVo.Description = resource.Description
	//golog.Infof("%s", rtVo)
	return
}

func ResourceToResourceListVos(resources []*user.Resource) (rtVos []*ResourceListVo) {
	for _, r := range resources {
		rtVo := ResourceToResourceListVo(r)
		rtVos = append(rtVos, rtVo)
	}
	return
}

func ResourceParamVoToResource(rp *ResourceParamVo) (result *user.Resource) {
	currentUser := rp.CurrentUser
	if rp.Id != "" {
		result = user.FindResourceById(rp.Id)
		result.UpdatorId = currentUser.UserAccount
		result.UpdatorName = currentUser.UserName
		result.UpdatorDepartmentId = currentUser.DepartmentId
		result.UpdatorDepartmentName = currentUser.DepartmentName
	} else {
		result = new(user.Resource)
		result.CreatorId = currentUser.UserAccount
		result.CreatorName = currentUser.UserName
		result.CreatorDepartmentId = currentUser.DepartmentId
		result.CreatorDepartmentName = currentUser.DepartmentName

		result.ResourceLevel = rp.ResourceLevel
		result.ResourceOrder = rp.ResourceOrder
		result.ResourceSate = rp.ResourceSate
	}
	result.ResourceType = "0"
	result.ResourceName = rp.ResourceName
	result.ParentId = rp.ParentId
	result.Url = rp.Url
	result.Icon = rp.Icon
	result.Description = rp.Description

	return
}
