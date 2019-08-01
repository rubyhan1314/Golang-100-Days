package main

import "fmt"

func main() {
	/*

	一维数组：存储的多个数据是数值本身
		a1 :=[3]int{1,2,3}

	二维数组：存储的是一维的一维
		a2 := [3][4]int{{},{},{}}

		该二维数组的长度，就是3。
		存储的元素是一维数组，一维数组的元素是数值，每个一维数组长度为4。

	多维数组：。。。


	 */
	a2 := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(a2)
	fmt.Printf("二维数组的地址：%p\n", &a2)
	fmt.Printf("二维数组的长度：%d\n", len(a2))

	fmt.Printf("一维数组的长度：%d\n", len(a2[0]))
	fmt.Println(a2[0][3]) // 4
	fmt.Println(a2[1][2]) //7
	fmt.Println(a2[2][1]) // 10

	//遍历二维数组
	for i:=0;i<len(a2);i++{
		for j:=0;j<len(a2[i]);j++{
			fmt.Print(a2[i][j],"\t")
		}
		fmt.Println()
	}
	fmt.Println("---------------------")
	//for range 遍历二维数组
	for _,arr := range a2{
		for _,val := range arr{
			fmt.Print(val,"\t")
		}
		fmt.Println()
	}
}
