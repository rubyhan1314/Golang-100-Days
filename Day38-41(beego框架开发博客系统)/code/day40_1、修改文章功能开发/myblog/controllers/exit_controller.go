package controllers

type ExitController struct {
	BaseController
}

func (this *ExitController)Get(){
	//清除该用户登录状态的数据
	this.DelSession("loginuser")

	this.Redirect("/",302)
}

