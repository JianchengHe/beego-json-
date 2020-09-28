package controllers

import (
	"beego190604/db_mysql"
	"beego190604/models"
	"crypto/md5"
	"encoding/hex"
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
		//r.Ctx.WriteString("数据接收错误，请重试")
		result := models.Result{
			Code:0,
			Message:"数据接收错误，请重试",
			Data:nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
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
		r.Data["json"] = &result
		r.ServeJSON()
		return
	}
	//直接调用保存数据的一个函数，并判断保存后的结果
	row,err := db_mysql.AddUser(user)
	if err != nil {
		//r.Ctx.WriteString("注册用户信息失败，请重试")
		result := models.Result{
			Code:0,
			Message:"注册用户信息失败，请重试",
			Data:nil,
		}
		r.Data["json"] = &result
		r.ServeJSON()
		return
	}
	fmt.Println(row)
	md5Hash := md5.New()
	md5Hash.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(md5Hash.Sum(nil))
	result := models.Result{
		Code:1,
		Message:"恭喜，注册用户信息成功",
		Data: user,
	}
	//json.Marshal(result) 编码
	r.Data["json"] = &result
	r.ServeJSON()//将result编码为json格式返回给前端
	//r.Ctx.WriteString("恭喜，注册用户信息成功")
}