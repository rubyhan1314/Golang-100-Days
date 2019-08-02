package routers

/**
 * 正则路由：
 *    在固定路由的基础上，beego支持通过正则表达式来解析http请求
 */
func init() {

	//*全匹配
	//beego.Router("/*", &controllers.RegController{})

	//:id变量匹配
	//beego.Router("/getUser/:name", &controllers.RegController{})

	////自定义正则表达式匹配
	// /getUser/davie
	//beego.Router("/getUser/:name([0-9]+)", &controllers.RegController{})

	//*.*匹配
	//  http://localhost:8080/upload/file/img/hellworld.png
	//beego.Router("/upload/*.*", &controllers.RegController{})

	//int类型匹配
	//beego.Router("/getUserInfo/:id:int", &controllers.RegController{})

	//string类型匹配
	//beego.Router("/getUserInfo/:username:string", &controllers.RegController{})

}
