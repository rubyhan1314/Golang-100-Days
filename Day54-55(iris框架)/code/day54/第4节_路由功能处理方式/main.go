package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {

	app := iris.New()

	/**
	 * 1.handle方式处理请求
	 */
	//GET: http://localhost:8002/userinfo
	app.Handle("GET", "/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		app.Logger().Error(" request path :", path)
	})

	//post
	app.Handle("POST", "/postcommit", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(" post reuqest, the url is : ", path)
		context.HTML(path)
	})

	//      http://localhost:8002?date=20190310&city=beijing
	//GET： http://localhost:8002/weather/2019-03-10/beijing
	//      http://localhost:8002/weather/2019-03-11/beijing
	//      http://localhost:8002/weather/2019-03-11/tianjin

	app.Get("/weather/{date}/{city}", func(context context.Context) {
		path := context.Path()
		date := context.Params().Get("date")
		city := context.Params().Get("city")
		context.WriteString(path + "  , " + date + " , " + city)
	})
	
	/**
	 * 1、Get 正则表达式 路由
	 * 使用：context.Params().Get("name") 获取正则表达式变量
	 *
	 */
	// 请求1：/hello/1  /hello/2  /hello/3 /hello/10000
	//正则表达式：{name}
	app.Get("/hello/{name}", func(context context.Context) {
		//获取变量
		path := context.Path()

		app.Logger().Info(path)
		//获取正则表达式变量内容值
		name := context.Params().Get("name")
		context.HTML("<h1>" + name + "</h1>")
	})

	/**
	 * 2、自定义正则表达式变量路由请求 {unit64:uint64}进行变量类型限制
	 */
	// 123
	// davie
	app.Get("/api/users/{userid:uint64}", func(context context.Context) {
		userID, err := context.Params().GetUint("userid")

		if err != nil {
			//设置请求状态码，状态码可以自定义
			context.JSON(map[string]interface{}{
				"requestcode": 201,
				"message":     "bad request",
			})
			return
		}

		context.JSON(map[string]interface{}{
			"requestcode": 200,
			"user_id":     userID,
		})
	})

	//自定义正则表达式路由请求 bool
	app.Get("/api/users/{isLogin:bool}", func(context context.Context) {
		isLogin, err := context.Params().GetBool("isLogin")
		if err != nil {
			context.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			context.WriteString(" 已登录 ")
		} else {
			context.WriteString(" 未登录 ")
		}

		//正则表达式所支持的数据类型
		context.Params()
	})

	app.Run(iris.Addr(":8002"), iris.WithoutServerError(iris.ErrServerClosed))

}
