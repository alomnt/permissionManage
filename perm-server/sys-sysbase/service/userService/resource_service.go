package userService

import (
	"fmt"
	"github.com/kataras/golog"
	"permissionManage/perm-server/sys-common/db"
	"permissionManage/perm-server/sys-common/supports"
	"permissionManage/perm-server/sys-common/vo"
	"permissionManage/perm-server/sys-sysbase/models/user"
	"permissionManage/perm-server/sys-sysbase/vo/uservo"
	"strings"
)

func CreateResource(resourceParam *uservo.ResourceParamVo) (res *supports.SimResult) {
	resource := uservo.ResourceParamVoToResource(resourceParam)
	maxOrder, orderErr := user.GetResourceMaxOrderByParentId(resourceParam.ParentId)
	if orderErr != nil {
		orderErr.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
		return
	}
	resource.ResourceOrder = maxOrder + 1

	if resourceParam.ParentId == "0" {
		resource.ResourceLevel = 1
	} else {
		level, levelErr := user.GetResourceLevelByParentId(resourceParam.ParentId)
		if levelErr != nil {
			levelErr.Error()
			res = &supports.SimResult{
				Code: false,
				Msg:  supports.ADD_FAIL,
			}
			return
		}
		resource.ResourceLevel = level + 1
	}

	_, err := user.CreateResource(resource)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Msg:  supports.ADD_SUCCESS,
	}

	return
}

func UpdateResource(resourceParam *uservo.ResourceParamVo) (res *supports.SimResult) {
	resource := uservo.ResourceParamVoToResource(resourceParam)
	_, err := user.UpdateResource(resource)
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

func DeleteResource(resourceParam *uservo.ResourceParamVo) (res *supports.SimResult) {
	idsStr := resourceParam.Ids
	_, err := user.DeleteResourceByIds(strings.Split(idsStr, ","))
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

func FindResourceListByUser(resourceParam *uservo.ResourceParamVo) (res *supports.SimResult) {
	var (
		err          error
		resourceList []*user.Resource
	)
	result := make([]*uservo.ResourceParamVo, 0)

	resourceList, err = user.GetResourceByUserId(resourceParam.UserId)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  "",
			Data: result,
		}
		return
	}

	res = &supports.SimResult{
		Code: true,
		Data: uservo.ResourceToResourceListVos(resourceList),
	}

	return
}

// 构建 菜单 tree
func FindResourceTreeByUser(resourceParam *uservo.ResourceParamVo) (res *supports.SimResult) {
	var (
		err          error
		resourceList []*user.Resource
	)

	resourceList, err = user.GetResourceByUserId(resourceParam.UserId)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.QUERY_FAIL,
		}
		return
	}
	tmpRtVO := uservo.ResourceToResourceTreeVos(resourceList)
	rtVO := buildResourceTree(tmpRtVO)
	res = &supports.SimResult{
		Code: true,
		Msg:  supports.LoginSuccess,
		Data: rtVO,
	}

	return
}
func FindResourceTreeAll(resourceParam *uservo.ResourceParamVo) (rtVO []*uservo.ResourceTreeVo, res *supports.SimResult) {
	var (
		err          error
		resourceList []*user.Resource
	)

	resourceList, err = FindResourceList(resourceParam)
	if err != nil {
		err.Error()
		res = &supports.SimResult{
			Code: false,
			Msg:  supports.ADD_FAIL,
		}
		return
	}
	tmpRtVO := uservo.ResourceToResourceTreeVos(resourceList)
	rtVO = buildResourceTree(tmpRtVO)

	res = &supports.SimResult{
		Code: true,
		Msg:  "",
		Data: rtVO,
	}
	return
}

func buildResourceTree(resourceTreeVo []*uservo.ResourceTreeVo) (rtVO []*uservo.ResourceTreeVo) {
	for _, rt := range resourceTreeVo {
		subs := getSubResourceTree(rt.Id, resourceTreeVo)
		if subs == nil {
			if rt.ParentId == "0" {
				rtVO = append(rtVO, rt)
			}
			continue
		}
		rt.Subs = subs
		rt.Children = subs
		buildResourceTree(subs)
		if rt.ParentId == "0" {
			rtVO = append(rtVO, rt)
		}
	}
	return
}
func getSubResourceTree(parentId string, resourceTreeVoAll []*uservo.ResourceTreeVo) (rtVO []*uservo.ResourceTreeVo) {
	for _, rt := range resourceTreeVoAll {
		if rt.ParentId == parentId {
			rtVO = append(rtVO, rt)
		}
	}

	return
}

func FindResourceList(resourceParam *uservo.ResourceParamVo) (rs []*user.Resource, err error) {
	e := db.MasterEngine()
	result := make([]*user.Resource, 0)

	sql := fmt.Sprintf(" SELECT * FROM cd_sys_user_resource ur " +
		" WHERE ur.is_deleted = false")

	err = e.SQL(sql).Find(&result)

	return result, err
}

// 分页
func FindPagesResource(paramVo *uservo.ResourceParamVo) (res *supports.SimResult) {
	e := db.MasterEngine()
	resourceListVos := make([]*uservo.ResourceListVo, 0)
	resources := make([]*user.Resource, 0)

	pageParam := vo.GetPageParam(paramVo.PageVo)

	sql := " is_deleted = false "
	if paramVo.ResourceName != "" {
		sql += " and resource_name like '%" + paramVo.ResourceName + "%'"
	}
	if paramVo.ResourceType != "" {
		sql += " and resource_type = '" + paramVo.ResourceType + "'"
	}
	s := e.Where(sql).Limit(pageParam.PageSize, pageParam.Start)
	count, err := s.FindAndCount(&resources)

	resourceListVos = uservo.ResourceToResourceListVos(resources)

	pageResult := &vo.PageResult{
		CurrentPage: pageParam.CurrentPage,
		PageSize:    pageParam.PageSize,
		Total:       count,
		Data:        resourceListVos,
	}
	res = &supports.SimResult{
		Code: true,
		Msg:  "",
		Data: pageResult,
	}
	golog.Println(count, err)
	return
}
