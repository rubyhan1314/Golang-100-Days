package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{})
	//注册
	beego.Router("/register", &controllers.RegisterController{})
	//登录
	beego.Router("/login", &controllers.LoginController{})

	//退出
	beego.Router("/exit", &controllers.ExitController{})
}
