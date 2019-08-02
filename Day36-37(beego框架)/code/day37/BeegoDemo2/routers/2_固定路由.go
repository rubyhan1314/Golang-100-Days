package routers

/**
 * 固定路由:
 *    beego框架支持我们通过beego.Router函数来进行路由注册;第一个参数url，代表用户的请求；
 * 第二个参数ControllerInterface用来指定进行请求逻辑处理的对象，通常我们命名为xxConttoller
 * 固定路由依然是根据http的请求方法的类型来自动执行对应的controller的方法，比如Get类型的请求会自动执行Get方法，Post请求会自动执行Post方法；以此来一一对应
 */

func init() {

	//固定路由的Get请求
	//beego.Router("/GetInfo", &controllers.MainController{})

	////固定路由的Post请求
	//beego.Router("/PostInfo", &controllers.MainController{})

	////固定路由的Delete请求
	//beego.Router("/DeleteInfo", &controllers.MainController{})

	////固定路由的Options请求
	//beego.Router("/OptionsInfo", &controllers.MainController{})
}
