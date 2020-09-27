package controllers

import (
	"beego190604/db_mysql"
	"beego190604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterControllers struct {
	beego.Controller
}

func (r *RegisterControllers) Post(){
	//解析前端用户的注册信息
	dataBytes,err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	//json包解析
	var user models.User
	err = json.Unmarshal(dataBytes,&user)
	if err != nil {
		r.Ctx.WriteString("数据解析错误，请重试。")
		return
	}
	//直接调用保存数据的一个函数，并判断保存后的结果
	row,err := db_mysql.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("注册用户信息失败，请重试")
		return
	}
	fmt.Println(row)
	r.Ctx.WriteString("恭喜，数据保存成功")
}