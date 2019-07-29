package responseHandle

import (
	"../../../sys-common/supports"
	"github.com/kataras/iris"
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
