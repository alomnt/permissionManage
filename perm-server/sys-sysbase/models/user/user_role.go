package user

import (
	"../../../sys-common/utils"
	"../../../sys-common/db"
	"time"
)

// 用户-角色关联表
type UserRole struct {
	Id                    string    `xorm:"pk varchar(32) notnull unique" json:"id" form:"id"`
	CreateTime            time.Time `json:"create_time" form:"create_time"`
	CreatorDepartmentId   string    `xorm:"varchar(50)" json:"creator_department_id" form:"creator_department_id"`
	CreatorDepartmentName string    `xorm:"varchar(255)" json:"creator_department_name" form:"creator_department_name"`
	CreatorId             string    `xorm:"varchar(50)"json:"creator_id" form:"creator_id"`
	CreatorName           string    `xorm:"varchar(255)"json:"creator_name" form:"creator_name"`
	IsDeleted             int64     `xorm:"INT(1) notnull" json:"isDeleted" form:"isDeleted"`
	UpdateTime            time.Time `json:"update_time" form:"update_time"`
	UpdatorDepartmentId   string    `xorm:"varchar(50)" json:"updator_department_id" form:"updator_department_id"`
	UpdatorDepartmentName string    `xorm:"varchar(255)" json:"updator_department_name" form:"updator_department_name"`
	UpdatorId             string    `xorm:"varchar(50)"json:"updator_id" form:"updator_id"`
	UpdatorName           string    `xorm:"varchar(255)"json:"updator_name" form:"updator_name"`
	Version               int64     `xorm:"INT(11)" json:"version" form:"version"`

	UserId string `xorm:"varchar(50) notnull" json:"user_id"`
	RoleId string `xorm:"varchar(50) notnull" json:"role_id"`
}

func (UserRole) TableName() string {
	return "cd_sys_user_user_role"
}

//=========================   新增用户角色	============================
func CreateRelationUserRole(userRoles ...*UserRole) (int64, error) {
	e := db.MasterEngine()
	var userRolEntitys []UserRole
	for _, ur := range userRoles {
		// 生成uuid
		ur.Id = utils.GetUuid()
		ur.CreateTime = time.Now()
		ur.Version = 0
		ur.IsDeleted = 0

		userRolEntitys = append(userRolEntitys, *ur)
	}
	return e.Insert(userRolEntitys)
}

//=========================   物理删除	============================
func RemoveUserRole(userId string, roleIds []string) (effect int64, err error) {
	e := db.MasterEngine()

	ur := new(UserRole)
	effect, err = e.Where("user_id = ?", userId).In("role_id", roleIds).Delete(ur)
	return
}
