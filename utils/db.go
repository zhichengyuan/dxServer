package utils

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "140258Nn"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test"
)

//Db数据库连接池
var DB *sql.DB

type Msql struct {
}

//启动
func (db *Msql) InitDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

//查询多条
func (db *Msql) QueryList(sql string, args interface{}) []interface{} {
	res := make([]interface{}, 13)
	rows, e := DB.Query(sql)
	if e != nil {
		fmt.Println("query incur error")
	}
	for rows.Next() {
		// data := make([]interface{}, 0)
		// e := rows.Scan(&data)
		// if e != nil {
		// 	fmt.Println("query error", e)
		// }
		rows.Scan(&res)

		res = append(res, rows)
	}
	rows.Close()
	defer DB.Close()
	return res
	// DB.QueryRow("select * from user where id=1").Scan(user.age, user.id, user.name, user.phone, user.sex)

	// stmt, e := DB.Prepare("select * from user where id=?")
	// query, e := stmt.Query(1)
	// query.Scan()

}

//单行
func (db *Msql) Query(sql string) *sql.Row {

	res := DB.QueryRow(sql)
	defer DB.Close()
	return res
}

//增删改
func (db *Msql) Insert(args map[string]string, sql string) bool {
	//开启事务
	vals := make([]interface{}, 0)
	for _, v := range args {
		// sqlStr += "(?, ?, ?),"
		vals = append(vals, v)

	}
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare(sql) //"INSERT INTO table_name (``) VALUES (?)"
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行

	res, err := stmt.Exec(vals)
	if err != nil {
		fmt.Println("Exec fail", err)
		return false
	}
	id, err := res.LastInsertId()
	if err != nil {
		return false
	}
	if id > 0 {
		//提交
		tx.Commit()
		return true
	} else {
		tx.Rollback()
		fmt.Println("Rowback Insert")
		return false
	}

	//获得上一个插入自增的id
	return true
}

// //删除
// func DeleteUser(user User) bool {
// 	//开启事务
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		fmt.Println("tx fail")
// 	}
// 	//准备sql语句
// 	stmt, err := tx.Prepare("DELETE FROM user WHERE id = ?")
// 	if err != nil {
// 		fmt.Println("Prepare fail")
// 		return false
// 	}
// 	//设置参数以及执行sql语句
// 	res, err := stmt.Exec(user.id)
// 	if err != nil {
// 		fmt.Println("Exec fail")
// 		return false
// 	}
// 	//提交事务
// 	tx.Commit()
// 	//获得上一个insert的id
// 	fmt.Println(res.LastInsertId())
// 	return true
// }
