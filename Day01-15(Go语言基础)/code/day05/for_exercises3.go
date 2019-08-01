package main

import (
	"math"
	"fmt"
)

func main() {
	/*
	 水仙花数：三位数：[100,999]
		每个位上的数字的立方和，刚好等于该数字本身。那么就叫水仙花数，4个
		比如：153
			1*1*1 + 5*5*5 + 3*3*3 = 1+125+27=153

		268
			268 /100 = 2
			268 % 10 = 8

			268 -->26  %10 = 6
			268 -->68  /10 =6
	 */

	for i := 100; i < 1000; i++ {
		x := i / 100     //百位
		y := i / 10 % 10 //十位
		z := i % 10      //个位

		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			fmt.Println(i)
		}
	}
	fmt.Println("------------------------")
	/*
	百位：1-9
	十位：0-9
	个位：0-9
	 */
	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Println(n)

				}
			}
		}
	}

}
