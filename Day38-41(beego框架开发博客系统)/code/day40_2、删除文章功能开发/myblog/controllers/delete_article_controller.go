package controllers

import (
	"fmt"
	"log"
	"myblog/models"
)

/**
 * 删除文章控制器
 */
type DeleteArticleController struct {
	BaseController
}

//点击删除后重定向到首页
func (this *DeleteArticleController) Get() {
	artID, _ := this.GetInt("id")
	fmt.Println("删除 id:", artID)

	_, err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	this.Redirect("/", 302)
}
