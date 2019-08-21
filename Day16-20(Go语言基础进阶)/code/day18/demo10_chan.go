package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go func() {
		fmt.Println("子goroutine开始执行。。")
		//time.Sleep(3 * time.Second)
		data := <-ch1 //从ch1中读取数据
		fmt.Println("data：",data)
	}()

	time.Sleep(5*time.Second)
	ch1 <- 10
	fmt.Println("main..over...")


	ch := make(chan int)
	ch <- 100 //阻塞
}
