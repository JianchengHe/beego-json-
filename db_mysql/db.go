package db_mysql

import (
	"beego190604/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB
func Connect() {
	//定于config变量接收并赋值为全局配置变量
	config := beego.AppConfig
	//获取配置选项
	appName := config.String("appname")
	fmt.Println("项目的应用昵称：", appName)
	port, err := config.Int("httpport")
	if err != nil {
		//配置信息解析错误
		panic("项目配置信息解析错误，请检查后重试")
	}
	fmt.Println("应用监听端口：", port)

	driver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	db, err := sql.Open(driver, dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8")
	if err != nil { //err不为nil,表示连接数据库时出现错误，程序就在此中断就不用在执行了。
		//早发现，早解决
		//panic:让程序进入一种恐慌状态，程序会终止执行。
		panic("数据库连接打开失败，请重试")
	}
	Db = db
	fmt.Println(db)
}

//将用户信息保存到数据中的函数
func AddUser(u models.User)(int64,error) {
	//1、将密码进行哈希计算，得到密码hash
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)

	//execute
	result,err := Db.Exec("insert into user(name1,birthday,address,password)"+
	"values(?,?,?,?)", u.Name1,u.Birthday,u.Address,u.Password)
	if err != nil {

		return -1,err
	}
	row,err := result.RowsAffected()
	if err != nil {

		return -1,err
	}
	return row,nil
}

//根据数据库里的用户名查询用户信息
func QueryUserName(name string) ([]models.User,error) {
	rows,err := Db.Query("select * from user where name1 = ?",name)
	if err != nil {
		return nil,err
	}
	users := make([]models.User,0)
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Name1,&user.Password,&user.Address,&user.Birthday)
		if err!= nil{
			return nil,err
		}
		users = append(users,user)
	}
	return users,nil
}

func QueryUser(name string)(int,error){
	rows := Db.QueryRow("select count(name1) user_num from user where name1 = ?",name)
	var user_num int
	err :=rows.Scan(&user_num)
	if err != nil {
		return 0,nil
	}
	return user_num,nil
}