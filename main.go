package main

import (
	_ "beego190604/routers"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	//定于config变量接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目的应用昵称：",appName)
	port,err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请检查后重试")
	}
	fmt.Println("应用监听端口：",port)

	driver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	db,err := sql.Open(driver,dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset = utf8")
	if err != nil {
		panic("数据库连接打开失败，请重试")
	}
	fmt.Println(db)
	beego.Run()
}

