package main

import (
	"time"
	"fmt"
)

func main() {
	/*
	 2. func After(d Duration) <-chan Time
			返回一个通道：chan，存储的是d时间间隔之后的当前时间

		相当于：return NewTimer(d).C
	 */
	ch :=  time.After(3 *time.Second)
	fmt.Printf("%T\n",ch) //<-chan time.Time
	fmt.Println(time.Now()) //2019-08-15 11:43:33.941039 +0800 CST m=+0.000537462


	time2 := <-ch
	fmt.Println(time2) //2019-08-15 11:43:36.945775 +0800 CST m=+3.005338577
}
