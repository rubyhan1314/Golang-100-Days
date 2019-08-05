package main

import (
	_ "myblog/routers"
	"github.com/astaxie/beego"
	"myblog/utils"
)

func main() {

	utils.InitMysql()
	beego.Run()
}

