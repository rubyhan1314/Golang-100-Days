package main

import "github.com/kataras/iris"

/**
 * 程序主入口
 */
func main() {

	app := iris.New()

	app.Run(iris.Addr(":9000"))

}
