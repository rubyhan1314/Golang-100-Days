package controllers

import "github.com/astaxie/beego"

//自定义路由控制器
type CustomController struct {
	beego.Controller
}

/**
 * 获取用户信息
 */
func (this *CustomController) GetUserInfo() {
	beego.Info("获取用户信息")

	username := this.GetString("username")

	userid := this.GetString("userid")

	this.Ctx.Output.Body([]byte("获取用户信息请求,用户名：" + username + " , 用户编号：" + userid))
}

/**
 *注册用户信息
 */
func (this *CustomController) RegisterInfo() {
	beego.Info("注册用户信息")
	username := this.GetString("username")
	userid := this.GetString("userid")
	this.Ctx.Output.Body([]byte("用户注册信息，username：" + username + " , userid：" + userid))
}
