package routers

import (
	"github.com/astaxie/beego"
	"BeegoDemo2/controllers"
)

func init() {
	
	//获取用户信息
	//get
	beego.Router("/getUserInfo", &controllers.CustomController{}, "GET:GetUserInfo")
	//post
	beego.Router("/registerInfo", &controllers.CustomController{}, "POST:RegisterInfo")
}
