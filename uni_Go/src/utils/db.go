package utils

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

//用户结构体
type Users struct {
	DataId    string `db:"id"`
	DataName  string `db:"Bookname"`
	Dataautor string `db:"autor"`
	DataTime  string `db:"time"`
	DataNum   string `db:"number"`
	DataUrl   string `db:"url"`
	Datadata  string `db:"data"`
}

//数据库指针
var db *sqlx.DB

//初始化数据库连接，init()方法系统会在动在main方法之前执行。
func init() {
	fmt.Printf("10086")
	database, err := sqlx.Open("mysql", "root:chenyonyu1@@tcp(gz-cynosdbmysql-grp-4edrl381.sql.tencentcdb.com:24729)/data")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}
func SelectData(str string) []Users {
	var user Users
	var fal []Users
	fmt.Printf("1005586")
	sql := str

	//执行SQL语句
	res, err := db.Query(sql)
	if err != nil {
		fmt.Println("exec failed,", err)
		return fal
	}
	var data []Users
	for res.Next() {
		err := res.Scan(&user.DataId, &user.DataName, &user.Dataautor, &user.DataTime, &user.DataNum, &user.DataUrl, &user.Datadata)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return fal
		}
		fmt.Printf("user:%#v\n", user)
		data = append(data, user)

	}
	fmt.Printf("user:%#v\n", data[0].DataId)
	return data
}

func InsertData(sqlStr string, username string, password string, email string) bool {
	/*sqlStr := "insert into login(username,password,email) value(?,?,?)"*/
	_, err := db.Exec(sqlStr, username, password, email)
	if err != nil {
		fmt.Println("插入失败")
		log.Panic(err.Error())
		return false
	} else {
		return true
		fmt.Println("插入成功")
	}
	return false
}
