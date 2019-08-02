package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//固定路由的get方法
func (this *MainController) Get() {
	beego.Info(" 固定路由的 get类型的方法 ")
	this.Ctx.Output.Body([]byte(" 固定路由 Get 请求"))
}

//处理PostInfo请求
func (this *MainController) Post() {
	beego.Info(" 固定路由的  Post类型的方法请求 ")
	this.Ctx.Output.Body([]byte(" 固定路由 Post 请求"))
}

func (this *MainController) Delete() {
	beego.Info(" 固定路由的 Delete类型的方法请求 ")
	this.Ctx.Output.Body([]byte("固定路由的 Delete 请求"))
}

func (this *MainController) Options() {
	beego.Info(" 固定路由的 options 类型的方法请求 ")
	this.Ctx.Output.Body([]byte("固定路由的 options 请求"))
}
