package main

import (
	"time"
	"fmt"
)

func main() {
	/*
	1. func NewTimer(d Duration) *Timer
			创建一个计时器，d时间以后触发


	 */
	 //timer := time.NewTimer(3 *time.Second)
	 //fmt.Printf("%T\n",timer) //*time.Timer
	 //fmt.Println(time.Now()) //2019-08-15 11:32:17.065452 +0800 CST m=+0.000614404
	 //
	 ////此处等待channel中的数值，会阻塞3秒
	 //ch2 := timer.C
	 //fmt.Println(<-ch2 ) //2019-08-15 11:32:20.068101 +0800 CST m=+3.003327715


	 //新建一个计时器
	 timer2 := time.NewTimer(5*time.Second)
	 //开始goroutine，来处理触发后的事件
	 go func() {
	 	<- timer2.C
	 	fmt.Println("Timer 2 结束了。。。开始。。。。")
	 }()

	 time.Sleep(3*time.Second)
	 flag := timer2.Stop()
	 if flag{
	 	fmt.Println("Timer 2 停止了。。。")
	 }
}
