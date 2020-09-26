package routers

import (
	"beego190604/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &controllers.MainController{})
    beego.Router("/register",&controllers.MainController{})
}
