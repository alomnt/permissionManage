package responseHandle

import (
	"github.com/kataras/iris/v12"
	"permissionManage/perm-server/sys-common/supports"
)

func Response(ctx iris.Context, res *supports.SimResult) {
	if res == nil {
		supports.Ok(ctx, "", "")
		return
	}

	if !res.Code {
		supports.Error(ctx, iris.StatusInternalServerError, res.Msg, nil)
		return
	}

	supports.Ok(ctx, res.Msg, res.Data)
}
