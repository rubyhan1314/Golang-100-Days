package main

import "fmt"

func main() {
	/*
	题目二：								i	j
	1X1=1								1	1
	2X1=2 2X2=4							2	1,2
	3X1=3 3X2=6 3X3=9					3	1,2,3
	...									...
	9X1=9 9X2=18 9X3=27...9X9=81		9	1,2,3,...9


	i X j = ...
	 */
	////第一行
	//for j := 1; j <= 1; j++ {
	//	fmt.Printf("%d X %d = %d\t", 1, j, 1)
	//}
	//fmt.Println()
	////第二行
	//for j := 1; j <= 2; j++ {
	//	fmt.Printf("%d X %d = %d\t", 2, j, 2*j)
	//}
	//fmt.Println()
	////第三行
	//for j := 1; j <= 3; j++ {
	//	fmt.Printf("%d X %d = %d\t", 3, j, 2*j)
	//}
	//fmt.Println()

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

}
