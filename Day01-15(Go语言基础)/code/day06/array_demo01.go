package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：整数，浮点，布尔，字符串
			复合类型：array，slice，map，struct，pointer，function，channel。。。

		数组：
			1.概念：存储一组相同数据类型的数据结构
					理解为容器，存储一组数据
			2.语法：
					var 数组名 [长度] 数据类型
					var 数组名 = [长度] 数据类型{元素1，元素2.。。}
					数组名 := [...]数据类型{元素。。。}

			3.通过下标访问
				下标，也叫索引：index，
				默认从0开始的整数，直到长度减1
				数组名[index]
					赋值
					取值

				不能越界：[0,长度-1]

			4.长度和容量：go语言的内置函数
				len(array/map/slice/string)，长度
				cap()，容量
	 */

	 var num1 int
	 num1 = 100

	 num1 = 200
	 fmt.Println(num1)
	 fmt.Printf("%p\n",&num1)

	 //step1：创建数组
	 var arr1 [4] int
	 fmt.Printf("%p\n",&arr1)
	 //step2：数组的访问
	 arr1[0] = 1
	 arr1[1] = 2
	 arr1[2] = 3
	 arr1[3] = 4
	 fmt.Println(arr1[0])//打印第一个数值
	 fmt.Println(arr1[2])//打印第三个数值
	 //fmt.Println(arr1[4]) //invalid array index 4 (out of bounds for 4-element array)

	 fmt.Println("数组的长度：",len(arr1)) //容器中实际存储的数据量
	 fmt.Println("数组的容量：",cap(arr1)) //容器中能够存储的最大的数量
	 //因为数组定长，长度和容量相同
	 arr1[0] = 100
	 fmt.Println(arr1[0])

	//数组的其他创建方式
	var a [4] int //同 var a= [4] int
	fmt.Println(a)

	var b = [4]int{1,2,3,4}
	fmt.Println(b)

	var c = [5]int{1,2,4}
	fmt.Println(c)

	var d = [5]int{1:1,3:2}
	fmt.Println(d)

	var e = [5]string{"rose","王二狗","ruby"}
	fmt.Println(e)

	f := [...]int{1,2,3,4,5}
	fmt.Println(f)
	fmt.Println(len(f))
	g:=[...]int{1:3,6:5}
	fmt.Println(g)
	fmt.Println(len(g))
}
