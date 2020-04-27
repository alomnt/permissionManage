package uservo

import (
	"permissionManage/perm-server/sys-common/models/baseModel"
	"permissionManage/perm-server/sys-sysbase/models/user"
	"time"
)

// 前端需要的数据结构
type UserRoleVO struct {
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
type UserRoleParamVo struct {
	*baseModel.CurrentUser

	CheckType string   `json:"checkType"`
	UserId    string   `json:"userId"`
	RoleIds   []string `json:"roleIds"`
}

func UserRoleParamVoToUserRole(paramVO *UserRoleParamVo) (result []*user.UserRole) {
	currentUser := paramVO.CurrentUser
	//roleIdArry := strings.Split(paramVO.RoleIds, ",")
	roleIdArry := paramVO.RoleIds
	for _, roleId := range roleIdArry {
		tmpRes := new(user.UserRole)
		tmpRes.CreatorId = currentUser.UserAccount
		tmpRes.CreatorName = currentUser.UserName
		tmpRes.CreatorDepartmentId = currentUser.DepartmentId
		tmpRes.CreatorDepartmentName = currentUser.DepartmentName
		tmpRes.UpdatorId = currentUser.UserAccount
		tmpRes.UpdatorName = currentUser.UserName
		tmpRes.UpdatorDepartmentId = currentUser.DepartmentId
		tmpRes.UpdatorDepartmentName = currentUser.DepartmentName
		tmpRes.UpdateTime = time.Now()

		tmpRes.UserId = paramVO.UserId
		tmpRes.RoleId = roleId

		result = append(result, tmpRes)
	}

	return
}
