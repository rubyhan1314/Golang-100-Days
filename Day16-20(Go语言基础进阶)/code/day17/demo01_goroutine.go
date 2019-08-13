package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	一个goroutine打印数字，另外一个goroutine打印字母，观察运行结果。。

	并发的程序的运行结果，每次都不一定相同。
	不同计算机设备执行，效果也不相同。


	go语言的并发：go关键字
		系统自动创建并启动主goroutine，执行对应的main()
		用于自己创建并启动子goroutine，执行对应的函数

		go 函数()//go关键创建并启动goroutine，然后执行对应的函数()，该函数执行结束，子goroutine也随之结束。

			子goroutine中执行的函数，往往没有返回值。
			如果有也会被舍弃。
	 */

	 //1.先创建并启动子goroutine，执行printNum()
	 go printNum()

	 //2.main中打印字母
	 for i:=1;i<=100;i++{
	 	fmt.Printf("\t主goroutine中打印字母：A %d\n",i)
	 }

	 time.Sleep(1*time.Second)
	 fmt.Println("main...over...")

}

func printNum() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("子goroutine中打印数字：%d\n", i)
	}
}
