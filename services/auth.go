package services

import (
	"base.com/utils"
	local "dx.com/utils"
	"github.com/kataras/iris/v12"
)

func init() {
	p := utils.NewRouterParty("user")
	p.RegisterRouter("POST", "/gettoken", user_gettoken)
	p.RegisterRouter("POST", "/login", user_login)

}

func user_gettoken(ctx iris.Context) {

	param := make(map[string]interface{})
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.Write(utils.GetResult(0, "param format error", ""))
		return
	}

	user_name := param["user_name"].(string)
	pwd := param["password"].(string)
	token, err := local.JwtCreate(user_name, pwd)
	if err != nil {
		ctx.Write(utils.GetResult(1, "error", err))
		return
	}
	ctx.Write(utils.GetResult(0, "get token success", token))
	// type UserInfo struct {
	// 	Username string
	// 	PWD      string
	// 	IsActive int64
	// }
	// user := UserInfo{}
	// db := utils.Msql{}
	// res := db.Query("select is_active from dx_user where user_name=" + param["username"].(string) + " and " + "password=" + param["password"].(string))
	// res.Scan(&user.IsActive)

	// if user.IsActive != 1 {
	// 	ctx.Write(utils.GetResult(-1, "", nil))
	// 	return
	// } else {

	// 	rest := make(map[string]interface{})
	// 	token, err := local.JwtCreate(user.Username, user.PWD)
	// 	if err == nil {
	// 		rest["token"] = token

	// 	} else {
	// 		rest["token"] = ""
	// 	}
	// 	ctx.Write(utils.GetResult(0, "", rest))

	// }
}

func user_login(ctx iris.Context) {

	param := make(map[string]interface{})
	if err := ctx.ReadJSON(&param); err != nil {
		ctx.Write(utils.GetResult(0, "param format error", ""))
		return
	}
	type UserInfo struct {
		Username string
		PWD      string
		IsActive int64
	}
	user := UserInfo{}
	db := utils.Msql{}
	res := db.Query("select is_active from dx_user where user_name=" + param["username"].(string) + " and " + "password=" + param["password"].(string))
	res.Scan(&user.IsActive)
}
