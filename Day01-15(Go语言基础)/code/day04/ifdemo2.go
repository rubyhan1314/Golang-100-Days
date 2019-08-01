package main

import "fmt"

func main() {
	/*
	if语句的其他写法：
	if 初始化语句; 条件{
		//注意变量的作用域问题
	}
	 */

	 if num := 4; num > 0{
	 	fmt.Println("正数。。",num)
	 }else if num < 0{
	 	fmt.Println("负数。。",num)
	 }
	 //fmt.Println(num) // undefined: num

	 num2 := 5
	 if num2 > 0{
	 	fmt.Println("num2，是正数。。",num2)
	 }
	 fmt.Println(num2)

}
