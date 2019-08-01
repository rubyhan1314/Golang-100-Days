package main

import "fmt"

func main() {
	/*
	1. 标准写法：
		for 表达式1;表达式2;表达式3{
			循环体
		}
	2.同时省略表达式1和表达式3
		for 表达式2{

		}
	相当于while(条件)
	3.同时省略3个表达式
		for{

		}
	相当于while(true)
	注意点：当for循环中，省略了表达式2，就相当于直接作用在了true上

	4.其他的写法：for循环中同时省略几个表达式都可以。。
		省略表达式1：
		省略表达式2：循环永远成立-->死循环
		省略表达式3：
	 */
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("-->",i)

	fmt.Println("------------")
	for{
		fmt.Println("i--->",i)
		i++
	}
}
