package utils

import (
	"base.com/utils"
	"github.com/kataras/iris/v12"
)

func CheckAuth(ctx iris.Context) bool {

	token := ctx.GetHeader("token")
	if token == "" {
		ctx.Write(utils.GetResult(1, "token error", nil))
		return false
	}

	_, err := utils.JwtParse(token)
	if err != nil {
		ctx.Write(utils.GetResult(1, "auth error", nil))
		return false
	}

	return true
}
