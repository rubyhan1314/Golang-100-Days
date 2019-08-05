package controllers

import (
	"myblog/models"
	"github.com/astaxie/beego"
)

/**
 * 相册控制器
 */
type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		beego.BeeLogger.Error(err.Error())
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
	//this.TplName = "album.html"
}
