package controllers

import (
	"fmt"
	"myblog/models"
)

type HomeController struct {
	BaseController
}

/**
 * 请求：http://localhost:8080/
 * 请求类型：Get
 * 请求描述：
 */
func (this *HomeController) Get() {

	//分页
	// http://localhost:8080/

	//标签
	// http://localhost:8080?tag=web

	tag := this.GetString("tag")
	fmt.Println("tag:", tag)
	page, _ := this.GetInt("page")
	var artList []models.Article

	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.QueryArticlesWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		//设置分页
		artList, _ = models.FindArticleWithPage(page)

		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

	this.TplName = "home.html"
}
