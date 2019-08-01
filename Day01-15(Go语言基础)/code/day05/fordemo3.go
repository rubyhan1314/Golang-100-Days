package main

import "fmt"

func main() {
	/*
	循环嵌套：多层循环嵌套在一起

	题目一：
	*****
	*****
	*****
	*****
	*****

	Print()
	Printf()
	Println()

	题目二：
	1X1=1
	2X1=2 2X2=4
	3X1=3 3X2=6 3X3=9
	...
	9X1=9 9X2=18 9X3=27...9X9=81
	 */
	//第一行
	//1.打印*
	//for j := 0; j < 5; j++ {
	//
	//	fmt.Print("*") // j:0,1,2,3,4
	//}
	////2.换行
	//fmt.Println()

	////第二行
	//for j := 0; j < 5; j++ {
	//
	//	fmt.Print("*") // j:0,1,2,3,4
	//}
	////2.换行
	//fmt.Println()

	for i := 1; i <= 5; i++ { //1,2,3,4,5
		for j := 0; j < 5; j++ {

			fmt.Print("*") // j:0,1,2,3,4
		}
		//2.换行
		fmt.Println()
	}
}
