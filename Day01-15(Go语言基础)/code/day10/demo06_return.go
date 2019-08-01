package main

import "fmt"

func main() {
	/*
	return语句：词义"返回"
		A：一个函数有返回值，那么使用return将返回值返回给调用处
		B：同时意味着结束了函数的执行

	注意点：
		1.一个函数定义了返回值，必须使用return语句将结果返回给调用处。return后的数据必须和函数定义的一致：个数，类型，顺序。
		2.可以使用_,来舍弃多余的返回值
		3.如果一个函数定义了有返回值，那么函数中有分支，循环，那么要保证，无论执行了哪个分支，都要有return语句被执行到
		4.如果一个函数没有定义返回值，那么函数中也可以使用return，专门用于结束函数的执行。。
	 */
	a, b, c := fun1()
	fmt.Println(a, b, c)
	_, _, d := fun1()
	fmt.Println(d)
	fmt.Println(fun2(-30))
	fun3()
}

func fun1() (int, float64, string) {
	return 0, 0, "hello"
}

func fun2(age int) int {
	if age >= 0 {
		return age

	} else {
		fmt.Println("年龄不能为负。。")
		return 0
	}
}

func fun3(){
	for i:=0;i<10;i++{
		if i == 5{
			//break //?强制结束循环
			return
		}
		fmt.Println(i)
	}
	fmt.Println("fun3()...over....")
}
