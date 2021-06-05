package services

import (
	"base.com/utils"
	"github.com/kataras/iris/v12"
)

func init() {
	p := utils.NewRouterParty("tporder")
	p.RegisterRouter("POST", "/save", order_save)
	p.RegisterRouter("POST", "/checktoken", user_gettoken)

}

func order_save(ctx iris.Context) {

}
