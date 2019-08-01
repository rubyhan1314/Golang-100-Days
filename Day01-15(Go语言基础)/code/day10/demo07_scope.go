package main

import "fmt"

//全局变量的定义
//num3 := 1000//不支持简短定义的写法
var num3 = 1000
func main() {
	/*
	作用域：变量可以使用的范围。
		局部变量：函数内部定义的变量，就叫做局部变量。
					变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。

		全局变量：函数外部定义的变量，就叫做全局变量。
					所有的函数都可以使用，而且共享这一份数据

	 */
	//定义在main函数中，所以n的作用域就是main函数的范围内
	 n:= 10
	 fmt.Println(n)

	 if a := 1;a <= 10{
	 	fmt.Println(a) // 1
	 	fmt.Println(n) // 10
	 }
	 //fmt.Println(a) //不能访问a，出了作用域
	 fmt.Println(n)

	 if b := 1;b <= 10{
	 	n := 20
	 	fmt.Println(b) // 1
	 	fmt.Println(n) // 20
	 }

	 fun1()
	 fun2()
	 fmt.Println("main中访问全局变量：",num3) //2000

}

func fun1(){
	//fmt.Println(n)
	num1 := 100
	fmt.Println("fun1()函数中：num1：",num1)
	num3 = 2000
	fmt.Println("fun1()函数，访问全局变量：",num3) // 2000
}

func fun2()  {
	num1 := 200
	fmt.Println(num1)
	fmt.Println("fun2()函数，访问全局变量：",num3) //2000
}
