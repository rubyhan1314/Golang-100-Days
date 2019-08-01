package main

import "fmt"

func main() {
	/*
	函数的参数：
		形式参数：也叫形参。函数定义的时候，用于接收外部传入的数据的变量。
			函数中，某些变量的数值无法确定，需要由外部传入数据。

		实际参数：也叫实参。函数调用的时候，给形参赋值的实际的数据


	函数调用：
		1.函数名：声明的函数名和调用的函数名要统一
		2.实参必须严格匹配形参：顺序，个数，类型，一一对应的。
	 */

	//1.求1-10的和
	getSum(10)
	//2.求1-20的和
	getSum(20)
	fmt.Println()

	//3.1-100的和
	getSum(100)

	//4.求2个整数的和
	getAdd(20,10)
	getAdd2(1,2)

	fun1(1.3,2.4,"hello")
}

//定义一个函数：用于求1-10的和
func getSum(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("1-%d的和是：%d\n",n, sum)
}

func getAdd(a int,b int){
	sum := a + b
	fmt.Printf("%d + %d = %d \n",a,b,sum)
}

func getAdd2(a, b int){//参数的类型一致，可以简写在一起
	fmt.Printf("a:%d,b:%d\n",a,b)
}

func fun1(a,b float64,c string){
	fmt.Printf("a:%.2f,b:%.2f,c:%s\n",a,b,c)
}

