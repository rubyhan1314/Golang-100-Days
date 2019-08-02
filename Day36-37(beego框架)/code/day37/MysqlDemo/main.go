package main

import (
	_ "BlogProject/MysqlDemo/routers"
	"github.com/astaxie/beego"
	//测试数据库连接
	_ "BlogProject/MysqlDemo/models"
)

func main() {
	beego.Run()
}
