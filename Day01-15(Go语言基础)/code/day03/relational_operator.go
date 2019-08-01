package main

import "fmt"

func main() {
	/*
	关系运算符：>,<,>=,<=,==,!=
		结果总是bool类型的：true，false
		==：表示比较两个数值是相等的
		!=：表示比较两个数值是不相等的
	 */
	 a := 3
	 b := 5
	 c := 3
	 res1 := a > b
	 res2 := b > c
	 fmt.Printf("%T,%t\n",res1,res1)
	 fmt.Printf("%T,%t\n",res2,res2)

	 res3 := a == b
	 fmt.Println(res3)

	 res4 := a == c
	 fmt.Println(res4)

	 fmt.Println(a != b, a != c)
}
