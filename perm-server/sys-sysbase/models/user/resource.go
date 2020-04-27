package user

import (
	"fmt"
	"permissionManage/perm-server/sys-common/db"
	"permissionManage/perm-server/sys-common/utils"
	"strconv"
	"time"
)

// cd_sys_user_resource
type Resource struct {
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

	ResourceType  string `xorm:"varchar(50) 'resource_type'"`
	ResourceName  string `xorm:"varchar(100) 'resource_name'"`
	ResourceLevel int64  `xorm:"INT(11) 'resource_level'"`
	ResourceOrder int64  `xorm:"INT(11) 'resource_order'"`
	ResourceSate  int64  `xorm:"tinyint(2) 'resource_sate'"`
	ParentId      string `xorm:"varchar(50) 'parent_id'"`
	Url           string `xorm:"varchar(100) 'url'"`
	Icon          string `xorm:"varchar(100) 'icon'"`
	Description   string `xorm:"varchar(255) 'description'"`
}

func (Resource) TableName() string {
	return "cd_sys_user_resource"
}

//=========================   新增	============================
func CreateResource(resources ...*Resource) (int64, error) {
	e := db.MasterEngine()
	var resourcesEntitys []Resource
	for _, r := range resources {
		// 生成uuid
		r.Id = utils.GetUuid()
		r.CreateTime = time.Now()
		r.Version = 0
		r.IsDeleted = 0

		resourcesEntitys = append(resourcesEntitys, *r)
	}
	return e.Insert(resourcesEntitys)
}

//=========================   更新	============================
func UpdateResource(resource *Resource) (int64, error) {
	e := db.MasterEngine()
	resource.UpdateTime = time.Now()
	resource.Version = resource.Version + 1
	aff, err := e.Id(resource.Id).Update(resource)
	if err != nil {
		err.Error()
	}
	return aff, err
}

//=========================   逻辑删除	============================
func DeleteResourceByIds(ids []string) (effect int64, err error) {
	e := db.MasterEngine()
	for _, id := range ids {
		tmpResource := FindResourceById(id)
		if tmpResource != nil {
			tmpResource.IsDeleted = 1
			_, err := e.Id(id).Update(tmpResource)
			if err != nil {
				err.Error()
			}
		}
	}
	return
}

//=========================   物理删除	============================
func RemoveResourceByIds(ids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	cr := new(Resource)
	for _, v := range ids {
		i, err1 := e.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}

//=========================   查询	============================
func FindResourceById(id string) (resource *Resource) {
	resource = new(Resource)
	e := db.MasterEngine()
	_, err := e.Id(id).Get(resource)
	if err != nil {
		err.Error()
		return
	}

	return
}

// ===== 根据父id查询当前级别菜单最大序号
func GetResourceMaxOrderByParentId(parentId string) (int64, error) {
	e := db.MasterEngine()

	sql := "SELECT MAX(ur.resource_order) AS maxOrder FROM cd_sys_user_resource ur WHERE ur.parent_id = '" + parentId + "'"
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
func GetResourceLevelByParentId(parentId string) (int64, error) {
	e := db.MasterEngine()

	sql := "SELECT MAX(ur.resource_level) AS parentLevel FROM cd_sys_user_resource ur WHERE ur.id = '" + parentId + "'"
	s := e.SQL(sql)
	re, err := s.QueryString(sql)
	if err != nil || len(re) == 0 {
		return int64(0), err
	}
	level, err := strconv.ParseInt(re[0]["parentLevel"], 10, 64)

	return level, err
}

// ===== 根据用户Id获取该用户所拥有的菜单
func GetResourceByUserId(uId string) ([]*Resource, error) {
	e := db.MasterEngine()
	sql := fmt.Sprintf(" SELECT * FROM cd_sys_user_resource ur "+
		" WHERE ur.id IN ( SELECT urr.resource_id FROM cd_sys_user_role_resource urr "+
		" WHERE urr.role_id IN ( SELECT uur.role_id FROM cd_sys_user_user_role uur "+
		" WHERE uur.user_id = '%s' ) ) "+
		" AND ur.resource_type = 0  "+
		" AND ur.is_deleted = 0 "+
		" ORDER BY ur.resource_order ASC", uId)

	result := make([]*Resource, 0)
	err := e.SQL(sql).Find(&result)

	return result, err
}
