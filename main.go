package main

import (
	"fmt"

	BaseService "base.com"
	"base.com/utils"
	_ "dx.com/services"
)

func main() {
	// bs:=
	db := utils.Msql{}
	db.InitDB()
	bs := BaseService.BaseService{}
	bs.Run("8870")
	fmt.Println()

	// u := User{}
	// // data := db.QueryList("SELECT * FROM dx_user", nil)
	// // db.Query("SELECT user_name FROM dx_user WHERE password = ?", 123456, name)
	// // var user_name string
	// // u := []string{"name", "is"}

	// data := db.Query("SELECT user_name  FROM dx_user WHERE password = ?", "123456", u.Name)
	// fmt.Println(data)

}
