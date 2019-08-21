package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	通过range访问通道
	 */
	 ch1 := make(chan int)
	 go sendData(ch1)
	 //for循环的for range，来访问通道
	 for v := range ch1{ // v <- ch1
	 	fmt.Println("读取数据：",v)
	 }
	 fmt.Println("main..over...")
}
func sendData(ch1 chan int){
	for i:=0;i<10;i++{
		time.Sleep(1* time.Second)
		ch1 <- i // 0 1...9
	}
	close(ch1)//通知对方，通道关闭
}
