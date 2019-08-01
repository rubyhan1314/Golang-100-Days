package main

import "fmt"

var a = 1000 //全局变量
var b int = 2000
//c := 3000 //语法错误
func main() {
	/*
	注意点：
	1.变量必须先定义才能使用
	2.变量的类型和赋值必须一致
	3.同一个作用域内，变量名不能冲突
	4.简短定义方式，左边的变量至少有一个是新的
	5.简短定义方式，不能定义全局变量
	6.变量的零值，就是默认值
		整型：默认值是0
		浮点类型：默认是0
		字符串：默认值""
	 */
	var num int
	num = 100
	fmt.Printf("num的数值是:%d,地址是:%p\n",num, &num)

	num = 200
	fmt.Printf("num的数值是:%d,地址是:%p\n",num, &num)

	// fmt.Println(num2) //undefined: num2

	var name string
	//name = 100
	//fmt.Println(name) //cannot use 100 (type int) as type string in assignment
	name = "王二狗"
	fmt.Println(name)

	//var name string  = "李小花"
	//fmt.Println(name)

	num ,name ,sex := 1000,"李小花" , "男" //no new variables on left side of :=
	fmt.Println(num,name,sex)

	fmt.Println(a)

	fmt.Println("----------------------")
	var m int //整数，默认值是0
	fmt.Println(m)
	var n float64 //0.0-->0
	fmt.Println(n)
	var s string //""
	fmt.Println(s)
	var s2 []int //切片[]
	fmt.Println(s2)
	fmt.Println(s2 == nil)

	//var sum = 100 //sum declared and not used

}
