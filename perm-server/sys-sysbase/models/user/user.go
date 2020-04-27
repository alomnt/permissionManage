package user

import (
	"permissionManage/perm-server/sys-common/db"
	"permissionManage/perm-server/sys-common/utils"
	"strconv"
	"time"
)

// table  cd_sys_user_user
type User struct {
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

	UserNumber   string `xorm:"varchar(50) 'user_number'"`
	UserAccount  string `xorm:"varchar(50) 'user_account'"`
	UserPassword string `xorm:"varchar(50) 'user_password'"`
	UserName     string `xorm:"varchar(50) 'user_name'"`
	Sex          string `xorm:"varchar(50) 'sex'"`
	DepartmentId string `xorm:"varchar(50) 'department_id'"`
	Mobile       string `xorm:"varchar(50) 'mobile'"`
	Email        string `xorm:"varchar(50) 'email'"`
	Address      string `xorm:"varchar(255) 'address'"`
}

type UserDepartment struct {
	User       `xorm:"extends"`
	Department `xorm:"extends"`
}

func (User) TableName() string {
	return "cd_sys_user_user"
}

func (UserDepartment) TableName() string {
	return "cd_sys_user_user"
}

//=========================   新增	============================
func CreateUser(users ...*User) (int64, error) {
	e := db.MasterEngine()
	var userEntitys []User
	for _, u := range users {
		u.Id = utils.GetUuid()
		u.CreateTime = time.Now()
		u.IsDeleted = 0
		// 密码加密
		u.UserPassword = utils.AESEncrypt([]byte(u.UserPassword))
		userEntitys = append(userEntitys, *u)
	}
	return e.Insert(userEntitys)
}

//=========================   更新	============================
func UpdateUser(eParam *User) (int64, error) {
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
func DeleteUserByIds(ids []string) (effect int64, err error) {
	e := db.MasterEngine()
	for _, id := range ids {
		user := FindUserById(id)
		if user != nil {
			user.IsDeleted = 1
			_, err := e.Id(id).Update(user)
			if err != nil {
				err.Error()
			}
		}
	}
	return
}

//=========================   物理删除	============================
func RemoveUserByIds(ids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	cr := new(User)
	for _, v := range ids {
		i, err1 := e.Id(v).Delete(cr)
		effect += i
		err = err1
	}
	return
}

//=========================   查询	============================
//// 根据id查询
//func FindUserById(userId string) *User {
//	user := new(User)
//	e := db.MasterEngine()
//	e.Get(&User{Id: userId})
//
//	return user
//}

// 根据id查询
func FindUserById(id string) (user *User) {
	user = new(User)
	e := db.MasterEngine()
	_, err := e.Id(id).Get(user)
	if err != nil {
		err.Error()
		return
	}

	return
}

// ===== 查询用户最大编号
func GetMaxUserNumber() (int64, error) {
	e := db.MasterEngine()

	sql := "SELECT MAX(CONVERT(uu.user_number,SIGNED)) AS maxUserNumber FROM cd_sys_user_user uu"
	s := e.SQL(sql)
	re, err := s.QueryString(sql)
	if err != nil || len(re) == 0 {
		return int64(0), err
	}
	if re[0]["maxUserNumber"] == "" {
		return int64(0), nil
	}
	maxUserNumber, err := strconv.ParseInt(re[0]["maxUserNumber"], 10, 64)

	return maxUserNumber, err
}

func GetUserByUsername(user *User) (bool, error) {
	e := db.MasterEngine()
	return e.Get(user)
}

func UpdateUserById(user *User) (int64, error) {
	e := db.MasterEngine()
	return e.Id(user.Id).Update(user)
}

func DeleteByUsers(uids []int64) (effect int64, err error) {
	e := db.MasterEngine()

	u := new(User)
	for _, v := range uids {
		i, err1 := e.Id(v).Delete(u)
		effect += i
		err = err1
	}
	return
}
