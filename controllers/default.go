package controllers

import (
	"beego190604/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller//匿名字段
}

func (c *MainController) Get() {
	//name1 := c.GetString("name")
	//age1,err := c.GetInt("age")
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(name1,age1)
	//获取get类型的请求参数
	/*name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)*/

	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)

	//以吴亦凡和20和male为正确数据
	if name != "吴亦凡"||age != "18"||sex != "male"{
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜你数据正确"))
	//以admin和18为正确数据验证
	/*if name != "admin" || age != "18"{
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))*/
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1019015673@qq.com"
	c.TplName = "index.tpl"
}


//该post方法是处理post类型的请求的时候
//func (c *MainController) Post(){
	/*fmt.Println("post类型的请求")
	user := c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：",user)
	pwd := c.Ctx.Request.FormValue("pwd")
	fmt.Println("密码为：",pwd)

	//与固定值进行比较
	if user != "admin" || pwd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据正确"))
	//request 请求，response 响应
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1019015673@qq.com"
	c.TplName = "index.tpl"*/
	//user := c.Ctx.Request.FormValue("user")
	//pwd := c.Ctx.Request.FormValue("pwd")
	//fmt.Println(user,pwd)
	////以user为JianchengHe，密码为123456
	//if user != "JianchengHe"||pwd != "123456"{
	//	c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
	//	return
	//}
	//	c.Ctx.ResponseWriter.Write([]byte("恭喜你，数据验证正确"))


	//解析数据
	//body := c.Ctx.Request.Body
	/*dataByes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	//json包解析
	var person models.JianchengHe_info
	err = json.Unmarshal(dataByes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("用户名：",person.Name,"年龄",person.Age)
	c.Ctx.WriteString("用户名是："+person.Name)
}*/

func (c *MainController) Post() {
	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil{
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	//json包解析
	 var person2 models.Person2
	err = json.Unmarshal(dataBytes,&person2)
	if err != nil {
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	fmt.Println("姓名：",person2.Name,"生日：",person2.Birthday,"地址：",person2.Address,"昵称：",person2.Nick)
}
