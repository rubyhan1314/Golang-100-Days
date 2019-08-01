package main

import "fmt"

func main() {
	/*
	panic：词义"恐慌"，
	recover："恢复"
	go语言利用panic()，recover()，实现程序中的极特殊的异常的处理
		panic(),让当前的程序进入恐慌，中断程序的执行
		recover(),让程序恢复，必须在defer函数中执行
	 */
	defer func(){
		if msg := recover();msg != nil{
			fmt.Println(msg,"程序回复啦。。。")
		}
	}()
	funA()
	defer myprint("defer main：3.....")
	funB()
	defer myprint("defer main：4.....")

	fmt.Println("main..over。。。。")

}
func myprint(s string){
	fmt.Println(s)
}

func funA(){
	fmt.Println("我是一个函数funA()....")
}

func funB(){//外围函数

	fmt.Println("我是函数funB()...")
	defer myprint("defer funB()：1.....")

	for i:= 1;i<=10;i++{
		fmt.Println("i:",i)
		if i == 5{
			//让程序中断
			panic("funB函数，恐慌了")
		}
	}//当外围函数的代码中发生了运行恐慌，只有其中所有的已经defer的函数全部都执行完毕后，该运行恐慌才会真正被扩展至调用处。
	defer myprint("defer funB()：2.....")
}



