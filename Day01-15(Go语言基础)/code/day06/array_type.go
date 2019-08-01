package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：int，float，string，bool。。
			复合类型：array，slice，map，function，pointer，channel。。

		数组的数据类型：
			[size]type

		值类型：理解为存储的数值本身
			将数据传递给其他的变量，传递的是数据的副本(备份)
				int,float,string,bool,array
		引用类型：理解为存储的数据的内存地址
				slice，map。。
	 */

	//1.数据类型
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [3]float64{2.15, 3.18, 6.19}
	arr3 := [4]int{5, 6, 7, 8}
	arr4 := [2]string{"hello", "world"}
	fmt.Printf("%T\n", arr1) //[4]int
	fmt.Printf("%T\n", arr2) //[3]float64
	fmt.Printf("%T\n", arr3) //[4]int
	fmt.Printf("%T\n", arr4) //[2]string

	//2.赋值
	num2 := num //值传递
	fmt.Println(num, num2) //10 10
	num2 = 20
	fmt.Println(num, num2) //10 20

	//数组呢
	arr5 := arr1 //值传递
	fmt.Println(arr1)
	fmt.Println(arr5)

	arr5[0] =1
	fmt.Println(arr1)
	fmt.Println(arr5)

	a := 3
	b := 4
	fmt.Println(a==b) //比较a和b的数值是否相等
	fmt.Println(arr5 == arr1) //比较数组的对应下标位置的数值是否相等
	//fmt.Println(arr1 == arr2) //invalid operation: arr1 == arr2 (mismatched types [4]int and [3]float64)

}
