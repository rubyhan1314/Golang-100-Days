package routers

/**
 *基础路由：
 *  beego框架提供了常见的http的请求类型方法的路由方案，比如：get，post，head，options，delete等方法
 */

func init() {

	//基础路由get方法
	//beego.Get("/get", func(context *context.Context) {
	//	beego.Info("基础路由中的get请求")
	//	context.Output.Body([]byte("基础路由中的get请求 get method"))
	//})

	//beego.Get("/getUserInfo", func(context *context.Context) {
	//	beego.Info("获取用户信息")
	//	context.Output.Body([]byte("获取用户信息请求"))
	//})

	//基础路由post方法
	//beego.Post("/post", func(this *context.Context) {
	//	beego.Info("基础路由中的post请求")
	//	this.Output.Body([]byte(" 基础路由 的post 请求"))
	////})

	//基础路由中的delete方法
	//beego.Delete("/deleteInfo", func(context *context.Context) {
	//	beego.Info("基础路由中的delete方法")
	//	context.ResponseWriter.Write([]byte("基础路由中的deleteInfo 方法"))
	//})

	//基础路由中的head方法
	//beego.Head("/head", func(context *context.Context) {
	//	beego.Info("基础路由中的head方法")
	//
	//})

	//基础路由中的options方法
	//beego.Options("/options", func(context *context.Context) {
	//	beego.Info("基础路由中的options方法")
	//})
}
