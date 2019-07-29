package sys

import (
	"../../../sys-sysbase/models/user"
	"../../db"
	"../../middleware/casbins"
	"../../utils"
	"time"

	"github.com/kataras/golog"
)

const (
	userNumber   = "0"
	userAccount  = "root"
	userPassword = "123456"
	userName     = "超级管理员"
)

// 检查超级用户是否存在
func CheckRootExit() bool {
	e := db.MasterEngine()
	// root is existed?
	exit, err := e.Exist(&user.User{UserAccount: userAccount})
	if err != nil {
		golog.Fatalf("@@@ When check Root User is exited? happened error. %s", err.Error())
	}
	if exit {
		golog.Info("@@@ Root User is existed.")

		// 初始化rbac_model
		r := user.User{UserAccount: userAccount}
		if exit, _ := e.Get(&r); exit {
			casbins.SetRbacModel(r.Id)
			CreateSystemRole()
		}
	}
	return exit
}

func CreateRoot() {
	newRoot := user.User{
		Id:           utils.GetUuid(),
		UserNumber:   userNumber,
		UserAccount:  userAccount,
		UserPassword: utils.AESEncrypt([]byte(userPassword)),
		UserName:     userName,
		CreateTime:   time.Now(),
	}

	e := db.MasterEngine()
	if _, err := e.Insert(&newRoot); err != nil {
		golog.Fatalf("@@@ When create Root User happened error. %s", err.Error())
	}
	rooId := newRoot.Id
	casbins.SetRbacModel(rooId)

	addAllpolicy(rooId)
}

func addAllpolicy(rooId string) {
	// add policy for root
	//p := casbins.GetEnforcer().AddPolicy(utils.FmtRolePrefix(newRoot.Id), "/*", "ANY", ".*")
	e := casbins.GetEnforcer()
	p := e.AddPolicy(rooId, "/*", "ANY", ".*", "", "", "", "", "", "超级用户")
	if !p {
		golog.Fatalf("初始化用户[%s]权限失败.", userName)
	}

	//
	for _, v := range Components {
		e.AddGroupingPolicy(rooId, v[0])
	}
}
