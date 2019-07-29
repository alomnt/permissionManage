package user

import (
	"../../../sys-common/db"
	"../../../sys-common/utils"
	"time"
	"strconv"
)

// table  cd_sys_user_department
type Department struct {
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

	DepName  string `xorm:"varchar(100) 'dep_name' index"`
	DepCode  string `xorm:"varchar(50) 'dep_code'"`
	DepLevel int64  `xorm:"INT(11) 'dep_level'"`
	DepOrder int64  `xorm:"INT(11) 'dep_order'"`
	ParentId string `xorm:"varchar(100) 'parent_id'"`
}

func (Department) TableName() string {
	return "cd_sys_user_department"
}

//=========================   新增	============================
func CreateDepartment(deps ...*Department) (int64, error) {
	e := db.MasterEngine()
	var roleEntitys []Department
	for _, d := range deps {
		// 生成uuid
		d.Id = utils.GetUuid()
		d.CreateTime = time.Now()
		d.Version = 0
		d.IsDeleted = 0

		roleEntitys = append(roleEntitys, *d)
	}
	return e.Insert(roleEntitys)
}

//=========================   更新	============================
func UpdateDepartment(eParam *Department) (int64, error) {
	e := db.MasterEngine()
	eParam.UpdateTime = time.Now()
	eParam.Version = eParam.Version + 1
	aff, err := e.Id(eParam.Id).Update(eParam)
	if err != nil {
		err.Error()
	}
	return aff, err
}

//=========================   逻辑删除	============================
func DeleteDepartmentByIds(ids []string) (effect int64, err error) {
	e := db.MasterEngine()
	for _, id := range ids {
		eModel := FindDepartmentById(id)
		if eModel != nil {
			eModel.IsDeleted = 1
			_, err := e.Id(id).Update(eModel)
			if err != nil {
				err.Error()
			}
		}
	}
	return
}

//=========================   物理删除	============================
func RemoveDepartmentByIds(ids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	cr := new(Department)
	for _, v := range ids {
		i, err1 := e.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}

//=========================   查询	============================
// 根据id查询
func FindDepartmentById(id string) (eRes *Department) {
	eRes = new(Department)
	e := db.MasterEngine()
	_, err := e.Id(id).Get(eRes)
	if err != nil {
		err.Error()
		return
	}

	return
}

// ===== 根据父id查询当前级别菜单最大序号
func GetDepartmentMaxOrderByParentId(parentId string) (int64, error) {
	e := db.MasterEngine()

	sql := "SELECT MAX(ud.dep_order) AS maxOrder FROM cd_sys_user_department ud WHERE ud.parent_id = '" + parentId + "'"
	s := e.SQL(sql)
	re, err := s.QueryString(sql)
	if err != nil || len(re) == 0 {
		return int64(0), err
	}
	if re[0]["maxOrder"] == "" {
		return int64(0), nil
	}
	maxOrder, err := strconv.ParseInt(re[0]["maxOrder"], 10, 64)

	return maxOrder, err
}

// ===== 根据父id查询当前菜单级别
func GetDepartmentParentLevelByParentId(parentId string) (int64, error) {
	e := db.MasterEngine()

	sql := "SELECT MAX(ud.dep_level) AS parentLevel FROM cd_sys_user_department ud WHERE ud.id = '" + parentId + "'"
	s := e.SQL(sql)
	re, err := s.QueryString(sql)
	if err != nil || len(re) == 0 {
		return int64(0), err
	}
	level, err := strconv.ParseInt(re[0]["parentLevel"], 10, 64)

	return level, err
}
