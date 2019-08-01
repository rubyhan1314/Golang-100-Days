package main

import "fmt"

func main() {
	/*
		递归函数(recursion)：一个函数自己调用自己，就叫做递归函数。
			递归函数要有一个出口，逐渐的向出口靠近
 */
	//1.求1-5的和
	sum := getSum(5)
	fmt.Println(sum)

	//2.fibonacci数列：
	/*
	1	2	3	4	5	6	7	8	9	10	11	12		。。。
	1	1	2	3	5	8	13	21	34	55	89	144


	 */
	 res := getFibonacci(12)
	 fmt.Println(res)
}
func getFibonacci(n int)int{
	if n== 1 || n == 2{
		return 1
	}
	return getFibonacci(n-1)+getFibonacci(n-2)
}

func getSum(n int)int{
	fmt.Println("**********")
	if n == 1{
		return 1
	}
	return getSum(n-1) + n
}
/*
求1-5的和
getSum(5)

	getSum(4) + 5

		getSum(3) + 4

			getSum(2) + 3

				getSum(1) + 2

				1
 */