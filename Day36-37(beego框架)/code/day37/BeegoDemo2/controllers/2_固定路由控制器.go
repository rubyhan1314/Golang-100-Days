package controllers

import "github.com/astaxie/beego"

type ChangelessController struct {
	beego.Controller
}

//固定路由的get方法
func (this *ChangelessController) Get() {
	this.Ctx.Output.Body([]byte("固定路由Get请求成功"))

}

//固定路由的Post方法
func (this *ChangelessController) Post() {
	this.Ctx.Output.Body([]byte("固定路由Post请求成功"))
}

//固定路由的delete方法
func (this *ChangelessController) Delete() {
	this.Ctx.ResponseWriter.Write([]byte("固定路由Delete请求成功"))
}

//固定路由的options方法
func (this *ChangelessController) Options() {
	this.Ctx.ResponseWriter.Write([]byte("固定路由Options请求成功"))
}
