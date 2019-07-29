package uservo

import (
	"../../../sys-common/models/baseModel"
	"../../models/user"
	"strings"
	"time"
)

// 前端需要的数据结构
type RoleResourceVO struct {
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

	RoleId     string `json:"roleId"`
	ResourceId string `json:"resourceId"`
}

// 请求参数
type RoleResourceParamVo struct {
	*baseModel.CurrentUser

	RoleId      string `json:"roleId"`
	ResourceIds string `json:"resourceIds"`
}

func RoleResourceParamVoToRoleResource(paramVO *RoleResourceParamVo) (result []*user.RoleResource) {
	currentUser := paramVO.CurrentUser
	resourceIdArry := strings.Split(paramVO.ResourceIds, ",")
	for _, resourceId := range resourceIdArry {
		tmpRes := new(user.RoleResource)
		tmpRes.CreatorId = currentUser.UserAccount
		tmpRes.CreatorName = currentUser.UserName
		tmpRes.CreatorDepartmentId = currentUser.DepartmentId
		tmpRes.CreatorDepartmentName = currentUser.DepartmentName
		tmpRes.UpdatorId = currentUser.UserAccount
		tmpRes.UpdatorName = currentUser.UserName
		tmpRes.UpdatorDepartmentId = currentUser.DepartmentId
		tmpRes.UpdatorDepartmentName = currentUser.DepartmentName
		tmpRes.UpdateTime = time.Now()

		tmpRes.RoleId = paramVO.RoleId
		tmpRes.ResourceId = resourceId

		result = append(result, tmpRes)
	}

	return
}
