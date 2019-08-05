package controllers

import (
	"fmt"
	"myblog/models"
)

/**
 * 更新文章的控制器
 */
type UpdateArticleController struct {
	BaseController
}

//当访问/update路径的时候回触发Get()方法，响应的页面是通过TplName这个属性指定返回给客户端的页面
func (this *UpdateArticleController) Get() {
	id, _ := this.GetInt("id")
	fmt.Println(id)

	//获取id所对应的文章信息
	art := models.QueryArticleWithId(id)

	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id

	this.TplName = "write_article.html"
}

//修改文章
func (this *UpdateArticleController) Post() {
	id, _ := this.GetInt("id")
	fmt.Println("postid:", id)

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	//实例化model，修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	//返回数据给浏览器
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}

	this.ServeJSON()
}
