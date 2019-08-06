package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Teacher struct {
	Name  string `json:"Name"`
	City  string `json:"City"`
	Other string `json:"Other"`
}

//程序主入口
func main() {

	app := iris.New()

	/**
	 * 通过Get请求
	 * 返回 WriteString
	 */
	app.Get("/getHello", func(context context.Context) {
		context.WriteString(" Hello world ")
	})

	/**
	 * 通过Get请求
	 * 返回 HTML数据
	 */
	app.Get("/getHtml", func(context context.Context) {
		context.HTML("<h1> Davie, 12 </h1>")
	})

	/**
	 * 通过Get请求
	 * 返回 Json数据
	 */
	app.Get("/getJson", func(context context.Context) {
		context.JSON(iris.Map{"message": "hello word", "requestCode": 200})
	})

	//POST
	app.Post("/user/postinfo", func(context context.Context) {
		//context.WriteString(" Post Request ")
		//user := context.Params().GetString("user")

		var tec Teacher
		err := context.ReadJSON(&tec)
		if err != nil {
			panic(err.Error())
		} else {
			app.Logger().Info(tec)
			context.WriteString(tec.Name)
		}

		//fmt.Println(user)
		//teacher := Teacher{}
		//err := context.ReadForm(&teacher)
		//if err != nil {
		//	panic(err.Error())
		//} else {
		//	context.WriteString(teacher.Name)
		//}
	})

	//PUT
	app.Put("/getHello", func(context context.Context) {
		context.WriteString(" Put Request ")
	})

	app.Delete("/DeleteHello", func(context context.Context) {
		context.WriteString(" Delete Request ")
	})

	//返回json数据
	app.Get("/getJson", func(context context.Context) {
		context.JSON(iris.Map{"message": "hello word", "requestCode": 200})
	})

	app.Get("/getStuJson", func(context context.Context) {
		context.JSON(Student{Name: "Davie", Age: 18})
	})

	//返回xml数据
	app.Get("/getXml", func(context context.Context) {
		context.XML(Person{Name: "Davie", Age: 18})
	})

	//返回Text
	app.Get("/helloText", func(context context.Context) {
		context.Text(" text hello world ")
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))

}

//json结构体
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//xml结构体
type Person struct {
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}
