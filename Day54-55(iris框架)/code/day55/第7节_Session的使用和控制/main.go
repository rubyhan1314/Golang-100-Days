package main

import (
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
)

var (
	USERNAME = "userName"
	ISLOGIN  = "isLogin"
)

/**
 * Session的使用和控制
 */
func main() {

	app := iris.New()

	sessionID := "mySession"

	//1、创建session并进行使用
	sess := sessions.New(sessions.Config{
		Cookie: sessionID,
	})

	/**
	 * 用户登录功能
	 */
	app.Post("/login", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(" 请求Path：", path)
		userName := context.PostValue("name")
		passwd := context.PostValue("pwd")

		if userName == "davie" && passwd == "pwd123" {
			session := sess.Start(context)

			//用户名
			session.Set(USERNAME, userName)
			//登录状态
			session.Set(ISLOGIN, true)

			context.WriteString("账户登录成功 ")

		} else {
			session := sess.Start(context)
			session.Set(ISLOGIN, false)
			context.WriteString("账户登录失败，请重新尝试")
		}
	})

	/**
	 * 用户退出登录功能
	 */
	app.Get("/logout", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(" 退出登录 Path :", path)
		session := sess.Start(context)
		//删除session
		session.Delete(ISLOGIN)
		session.Delete(USERNAME)
		context.WriteString("退出登录成功")
	})


	app.Get("/query", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(" 查询信息 path :", path)
		session := sess.Start(context)

		isLogin, err := session.GetBoolean(ISLOGIN)
		if err != nil {
			context.WriteString("账户未登录,请先登录 ")
			return
		}

		if isLogin {
			app.Logger().Info(" 账户已登录 ")
			context.WriteString("账户已登录")
		} else {
			app.Logger().Info(" 账户未登录 ")
			context.WriteString("账户未登录")
		}

	})


	//2、session和db绑定使用
	db, err := boltdb.New("sessions.db", 0600)
	if err != nil {
		panic(err.Error())
	}

	//程序中断时，将数据库关闭
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})

	//session和db绑定
	sess.UseDatabase(db)

	app.Run(iris.Addr(":8009"), iris.WithoutServerError(iris.ErrServerClosed))
}
