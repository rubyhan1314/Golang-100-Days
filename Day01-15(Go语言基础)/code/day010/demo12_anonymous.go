package main

import "fmt"

func main() {
	/*
	匿名：没有名字
		匿名函数：没有名字的函数。

	定义一个匿名函数，直接进行调用。通常只能使用一次。也可以使用匿名函数赋值给某个函数变量，那么就可以调用多次了。

	匿名函数：
		Go语言是支持函数式编程：
		1.将匿名函数作为另一个函数的参数，回调函数
		2.将匿名函数作为另一个函数的返回值，可以形成闭包结构。
	 */
	 fun1()
	 fun1()
	 fun2 := fun1
	 fun2()

	 //匿名函数
	 func (){
	 	fmt.Println("我是一个匿名函数。。")
	 }()

	 fun3:=func(){
	 	fmt.Println("我也是一个匿名函数。。")
	 }
	 fun3()
	 fun3()

	 //定义带参数的匿名函数
	 func (a,b int){
	 	fmt.Println(a,b)
	 }(1,2)

	 //定义带返回值的匿名函数
	 res1 := func (a, b int)int{
	 	return a + b
	 }(10,20) //匿名函数调用了，将执行结果给res1
	 fmt.Println(res1)

	 res2 := func (a,b int)int{
	 	return a + b
	 } //将匿名函数的值，赋值给res2
	 fmt.Println(res2)

	 fmt.Println(res2(100,200))
}

func fun1(){
	fmt.Println("我是fun1()函数。。")
}
