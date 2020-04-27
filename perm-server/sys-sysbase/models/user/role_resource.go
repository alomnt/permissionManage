package user

import (
	"permissionManage/perm-server/sys-common/db"
	"permissionManage/perm-server/sys-common/utils"
	"time"
)

// 角色-资源关联表
type RoleResource struct {
	Id                    string    `xorm:"pk varchar(32) notnull unique 'id'"`
	CreateTime            time.Time `xorm:"'create_time'"`
	CreatorDepartmentId   string    `xorm:"varchar(50) 'creator_department_id'"`
	CreatorDepartmentName string    `xorm:"varchar(255) 'creator_department_name'"`
	CreatorId             string    `xorm:"varchar(50) 'creator_id'"`
	CreatorName           string    `xorm:"varchar(255) 'creator_name'"`
	IsDeleted             int64     `xorm:"INT(1) notnull 'is_deleted'"`
	UpdateTime            time.Time `xorm:"'update_time'"`
	UpdatorDepartmentId   string    `xorm:"varchar(50) 'updator_department_id'"`
	UpdatorDepartmentName string    `xorm:"varchar(255) 'updator_department_name'"`
	UpdatorId             string    `xorm:"varchar(50) 'updator_id'"`
	UpdatorName           string    `xorm:"varchar(255) 'updator_name'"`
	Version               int64     `xorm:"INT(11) 'version'"`

	RoleId     string `xorm:"varchar(50) notnull 'role_id'"`
	ResourceId string `xorm:"varchar(50) notnull 'resource_id'"`
}

func (RoleResource) TableName() string {
	return "cd_sys_user_role_resource"
}

//
func CreateRelationRoleResource(roleResources ...*RoleResource) (int64, error) {
	e := db.MasterEngine()
	var roleResourceEntitys []RoleResource
	for _, r := range roleResources {
		// 生成uuid
		r.Id = utils.GetUuid()
		r.CreateTime = time.Now()
		r.Version = 0
		r.IsDeleted = 0

		roleResourceEntitys = append(roleResourceEntitys, *r)
	}
	return e.Insert(roleResourceEntitys)
}

//=========================   物理删除	============================
func RemoveRoleResource(roleId string) (effect int64, err error) {
	e := db.MasterEngine()

	rr := new(RoleResource)
	effect, err = e.Where("role_id = ?", roleId).Delete(rr)

	return
}
