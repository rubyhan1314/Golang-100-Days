package main

import (
	"github.com/kataras/iris"
	"os"
	"encoding/json"
	"fmt"
)

/**
 * Iris配置设置案例
 */
func main() {

	//1.新建app实例
	app := iris.New()

	//一、通过程序代码对应用进行全局配置
	app.Configure(iris.WithConfiguration(iris.Configuration{
		//如果设置为true，当人为中断程序执行时，则不会自动正常将服务器关闭。如果设置为true，需要自己自定义处理。
		DisableInterruptHandler: false,
		//该配置项表示更正并将请求的路径重定向到已注册的路径
		//比如：如果请求/home/ 但找不到此Route的处理程序，然后路由器检查/home处理程序是否存在，如果是，（permant）将客户端重定向到正确的路径/home。
		//默认为false
		DisablePathCorrection: false,
		//
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon,02 Jan 2006 15:04:05 GMT",
		Charset:                           "utf-8",
	}))

	//二、通过读取tml配置文件读取服务配置
	//注意：要在run方法运行之前执行
	app.Configure(iris.WithConfiguration(iris.TOML("/Users/hongweiyu/go/src/irisDemo/5-路由组及Iris配置/configs/iris.tml")))

	//三、通过读取yaml配置文件读取服务配置
	//同样要在run方法运行之前执行
	app.Configure(iris.WithConfiguration(iris.YAML("/Users/hongweiyu/go/src/irisDemo/5-路由组及Iris配置/configs/iris.yml")))

	//四、通过json配置文件进行应用配置
	file, _ := os.Open("/Users/hongweiyu/go/src/irisDemo/5-路由组及Iris配置/config.json")
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	conf := Coniguration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Port)

	//2.运行服务，端口监听
	app.Run(iris.Addr(":8009"))
}

type Coniguration struct {
	AppName string `json:"appname"`
	Port    int    `json:"port"`
}
