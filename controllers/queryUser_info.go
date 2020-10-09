package controllers

import (
	"beego190604/db_mysql"
	"beego190604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type QueryUser_info struct {
	beego.Controller
}

func (q *QueryUser_info) Post(){
	//解析前端用户的注册信息
	dataBytes,err := ioutil.ReadAll(q.Ctx.Request.Body)
	if err != nil {
		//r.Ctx.WriteString("数据接收错误，请重试")
		result := models.Result{
			Code:0,
			Message:"数据接收错误，请重试",
			Data:nil,
		}
		q.Data["json"] = &result
		q.ServeJSON()
		return
	}
	//json包解析
	var user models.User
	err = json.Unmarshal(dataBytes,&user)
	if err != nil {
		//r.Ctx.WriteString("数据解析错误，请重试。")
		result := models.Result{
			Code:0,
			Message:"数据解析错误，请重试。",
			Data:nil,
		}
		q.Data["json"] = &result
		q.ServeJSON()
		return
	}
	//到数据库匹配信息
	user_num,err := db_mysql.QueryUser(user.Name1)
	if err != nil {
		fmt.Println(err.Error())
	}
	users,err := db_mysql.QueryUserName(user.Name1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user_num)

	if err != nil {
		fmt.Println(err.Error())
		result := models.Result{
			Code:0,
			Message:"数据匹配错误，请重试。",
			Data:nil,
		}
		q.Data["json"] = &result
		q.ServeJSON()
		return
	}
	if user_num > 0  {
		result := models.Result{
			Code:1,
			Message:"恭喜，数据返回成功",
			Data:users,
		}
		q.Data["json"] = &result
		q.ServeJSON()//将result编码为json格式返回给前端
		//r.Ctx.WriteString("恭喜，注册用户信息成功")
	}
}