package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {

	c.Data["wechat"] = "微信：13167582311"
	c.Data["qq"] = "QQ：861574834"
	c.Data["tel"] = "Tel：13167582311"

	c.TplName = "aboutme.html"
}
