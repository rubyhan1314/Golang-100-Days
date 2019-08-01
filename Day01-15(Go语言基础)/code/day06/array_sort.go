package main

import "fmt"

func main() {
	/*
	数组的排序：
		让数组中的元素具有一定的顺序。

		arr :=[5]int{15,23,8,10,7}
			升序：[7,8,10,15,23]
			将序：[23,15,10,8,7]

	排序算法：
		冒泡排序，插入排序，选择排序，希尔排序，堆排序，快速排序。。。。

	冒泡排序：（Bubble Sort）
		依次比较两个相邻的元素，如果他们的顺序（如从大到小）就把他们交换过来。
	 */
	arr := [5]int{15, 23, 8, 10, 7}
	////第一轮排序
	//for j := 0; j < 4; j++ {
	//	if arr[j] > arr[j+1] {
	//		arr[j], arr[j+1] = arr[j+1], arr[j]
	//	}
	//}
	//fmt.Println(arr)
	//
	////第二轮排序
	//for j:=0;j<3;j++{
	//	if arr[j] > arr[j+1] {
	//		arr[j], arr[j+1] = arr[j+1], arr[j]
	//	}
	//}
	//fmt.Println(arr)

	for i:=1;i<len(arr);i++{
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}

}
