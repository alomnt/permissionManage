package user

import (
	"../../../sys-common/db"
	"../../../sys-common/utils"
	"time"
)

// table  cd_sys_user_role
type Role struct {
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

	PType       string `xorm:"varchar(100) index 'p_type'"`
	RoleType    string `xorm:"varchar(50) 'role_type'"`
	RoleName    string `xorm:"varchar(100) 'role_name'"`
	RoleCode    string `xorm:"varchar(20) 'role_code'"`
	Description string `xorm:"varchar(255) 'description'"`
}

func (Role) TableName() string {
	return "cd_sys_user_role"
}

//=========================   新增	============================
func CreateRole(roles ...*Role) (int64, error) {
	e := db.MasterEngine()
	var roleEntitys []Role
	for _, r := range roles {
		// 生成uuid
		r.Id = utils.GetUuid()
		r.CreateTime = time.Now()
		r.Version = 0
		r.IsDeleted = 0

		roleEntitys = append(roleEntitys, *r)
	}
	return e.Insert(roleEntitys)
}

//=========================   更新	============================
func UpdateRole(role *Role) (int64, error) {
	e := db.MasterEngine()
	role.UpdateTime = time.Now()
	role.Version = role.Version + 1
	aff, err := e.Id(role.Id).Update(role)
	if err != nil {
		err.Error()
	}
	return aff, err
}

//=========================   逻辑删除	============================
func DeleteRoleByIds(ids []string) (effect int64, err error) {
	e := db.MasterEngine()
	for _, id := range ids {
		role := FindRoleById(id)
		if role != nil {
			role.IsDeleted = 1
			_, err := e.Id(id).Update(role)
			if err != nil {
				err.Error()
			}
		}
	}
	return
}

//=========================   物理删除	============================
func RemoveRoleByIds(ids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	cr := new(Role)
	for _, v := range ids {
		i, err1 := e.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}

//=========================   查询	============================
// 根据id查询
func FindRoleById(id string) (role *Role) {
	role = new(Role)
	e := db.MasterEngine()
	_, err := e.Id(id).Get(role)
	if err != nil {
		err.Error()
		return
	}

	return
}
