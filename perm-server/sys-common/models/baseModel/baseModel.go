package baseModel

import (
	"permissionManage/perm-server/sys-common/caches"
	"permissionManage/perm-server/sys-common/db"
	"permissionManage/perm-server/sys-common/utils"
	"time"
)

// table  cd_sys_user_user
type BaseModel struct {
	Id                    string    `json:"id" form:"id"`
	CreateTime            time.Time `json:"create_time" form:"create_time"`
	CreatorDepartmentId   string    `json:"creator_department_id" form:"creator_department_id"`
	CreatorDepartmentName string    `json:"creator_department_name" form:"creator_department_name"`
	CreatorId             string    `json:"creator_id" form:"creator_id"`
	CreatorName           string    `json:"creator_name" form:"creator_name"`
	IsDeleted             int64     `json:"isDeleted" form:"isDeleted"`
	UpdateTime            time.Time `json:"update_time" form:"update_time"`
	UpdaterDepartmentId   string    `json:"updater_department_id" form:"updater_department_id"`
	UpdaterDepartmentName string    `json:"updater_department_name" form:"updater_department_name"`
	UpdaterId             string    `json:"updater_id" form:"updater_id"`
	UpdaterName           string    `json:"updater_name" form:"updater_name"`
	Version               int64     `json:"version" form:"version"`
}

type CurrentUser struct {
	UserId         string `json:"userId"`
	UserAccount    string `json:"userAccount"`
	UserName       string `json:"userName"`
	DepartmentId   string `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
	Roles          string `json:"roles"`
	Resource       string `json:"resource"`
}

func GetCreateBaseModel(currentUser *CurrentUser) (baseModel *BaseModel) {
	baseModel.Id = utils.GetUuid()
	baseModel.CreateTime = time.Now()
	baseModel.CreatorId = currentUser.UserAccount
	baseModel.CreatorName = currentUser.UserName
	baseModel.CreatorDepartmentId = currentUser.DepartmentId
	baseModel.CreatorDepartmentName = currentUser.DepartmentName
	baseModel.IsDeleted = 0

	return baseModel
}

func InitCurrentUserByUserAccount(userAccount string) (res *CurrentUser, err error) {
	e := db.MasterEngine()

	sql := "SELECT " +
		" u.id AS userId, " +
		" u.user_account AS userAccount, " +
		" u.user_name AS userName, " +
		" u.department_id AS departmentId, " +
		" ud.dep_name AS departmentName " +
		" FROM cd_sys_user_user u " +
		"	LEFT JOIN cd_sys_user_department ud ON u.department_id = ud.id " +
		" WHERE u.is_deleted = FALSE " +
		" AND u.user_account = '" + userAccount + "'"
	s := e.SQL(sql)
	queryResult, err := s.QueryString(sql)
	if err != nil || len(queryResult) == 0 {
		return
	}

	res = new(CurrentUser)
	res.UserId = queryResult[0]["userId"]
	res.UserName = queryResult[0]["userName"]
	res.UserAccount = queryResult[0]["userAccount"]
	res.DepartmentId = queryResult[0]["departmentId"]
	res.DepartmentName = queryResult[0]["departmentName"]

	return
}

func GetCurrentUset(userId string) *CurrentUser {
	curr := new(CurrentUser)
	currByt, err := caches.GetJsonByte("currentUser_" + userId)
	if err != nil {
		err.Error()
	}
	res := utils.ByteToStruct(currByt, curr)
	return res.(*CurrentUser)
}
