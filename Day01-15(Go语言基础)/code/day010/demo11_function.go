package main

import "fmt"

func main() {
	/*
	Go语言的数据类型：

		数值类型：整数，浮点
			进行运算操作，加减乘除，打印
		字符串：
			可以获取单个字符，截取子串，遍历，strings包下的函数操作。。
		数组，切片，map。。
			存储数据，修改数据，获取数据，遍历数据。。。
		函数：
			加(),进行调用

		注意点：
			函数作为一种复合数据类型，可以看做是一种特殊的变量。
				函数名()：将函数进行调用，函数中的代码会全部执行，然后将return的结果返回给调用处
				函数名：指向函数体的内存地址

	 */
	 //1.整型
	 a := 10
	 //运算：
	 a += 5
	 fmt.Println("a:",a)
	 //2.数组，切片，map。。容器
	 b := [4]int{1,2,3,4}
	 b[0] =100
	 for i:=0;i<len(b);i++{
	 	fmt.Printf("%d\t",b[i])
	 }
	 fmt.Println()

	 //3.函数做一个变量
	 fmt.Printf("%T\n",fun1) //func(int, int)
	 fmt.Println(fun1) // fun1() fun1  0x1094db0,看做函数名对应的函数体的地址

	 //4.直接定义一个函数类型的变量
	 var c func(int,int)
	 fmt.Println(c) //<nil> 空

	 //var d  string
	 //d = "hello"
	 c = fun1 //将fun1的值(函数体的地址)赋值给c
	 fmt.Println(c)

	 fun1(10,20)
	 c(100,200) //c也是函数类型的，加小括号也可以被调用

	 res1 := fun2 //将fun2的值(函数的地址)赋值给res1，res1和fun2指向同一个函数体
	 res2 := fun2(1,2) //将fun2函数进行调用，将函数的执行结果赋值给res2，相当于：a+b
	 fmt.Println(res1)
	 fmt.Println(res2)

	 fmt.Println(res1(10,20))//也可以被调用
	 //res2() //cannot call non-function res2 (type int)
}
func fun2(a,b int)int{
	return a + b
}
func fun1(a, b int){
	fmt.Printf("a:%d,b:%d\n",a,b)
}
