package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	 打印2-100内的素数(只能被1和本身整除)


		2，3，5，7，11，13，17.。。。。

		7，
			2，3，4，5，6

		8，
			2
	 */
	for i := 2; i <= 100; i++ {
		flag := true //记录i是否是素数
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ { //判断到根号i就可以，不需要到i的前一个
			if i%j == 0 {
				flag = false //不是不素了
				break
			}
		}

		if flag { //== true
			fmt.Println(i)
		}
	}

}
