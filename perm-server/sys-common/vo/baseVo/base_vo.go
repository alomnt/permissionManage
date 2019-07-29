package baseVo

import (
	"time"
)

// 基础信息
type BaseVO struct {
	Id                    string    `json:"id" form:"id"`
	CreateTime            time.Time `json:"createTime" form:"createTime"`
	CreatorDepartmentId   string    `json:"creatorDepartmentId" form:"creatorDepartmentId"`
	CreatorDepartmentName string    `json:"creatorDepartmentName" form:"creatorDepartmentName"`
	CreatorId             string    `json:"creatorId" form:"creatorId"`
	CreatorName           string    `json:"creatorName" form:"creatorName"`
	IsDeleted             int64     `json:"isDeleted" form:"isDeleted"`
	UpdateTime            time.Time `json:"updateTime" form:"updateTime"`
	UpdatorDepartmentId   string    `json:"updatorDepartmentId" form:"updatorDepartmentId"`
	UpdatorDepartmentName string    `json:"updatorDepartmentName" form:"updatorDepartmentName"`
	UpdatorId             string    `json:"updatorId" form:"updatorId"`
	UpdatorName           string    `json:"updatorName" form:"updatorName"`
	Version               int64     `json:"version" form:"version"`
	Token                 string    `json:"token"`
}
