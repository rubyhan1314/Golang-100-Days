package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	关闭通道：close(ch)
		子goroutine：写出10个数据
			每写一个，阻塞一次，主goroutine读取一次，解除阻塞

		主goroutine，读取数据
			每次读取数据，阻塞一次，子goroutine，写出一个，解除阻塞

	 */
	 ch1 := make(chan int)
	 go sendData(ch1)

	 //读取通道的数据
	 for{
	 	time.Sleep(1*time.Second)
	 	v, ok := <- ch1  // 最后一次读取
	 	if !ok{
	 		fmt.Println("已经读取了所有的数据。。",ok,v)
	 		break
		}
		fmt.Println("读取的数据：",v,ok)
	 }

	 fmt.Println("main..over...")

}

func sendData(ch1 chan int){
	//发送方：10条数据
	for i:=0;i<10;i++{
		ch1 <- i //将i写入到通道中
		// 0 1  .. 9
	}
	close(ch1) //将ch1通道关闭
}

