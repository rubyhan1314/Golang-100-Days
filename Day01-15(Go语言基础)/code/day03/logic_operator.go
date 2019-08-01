package main

import "fmt"

func main() {
	/*
	逻辑运算符：操作数必须是bool，运算结果也是bool
	逻辑与：&&
		运算规则：所有的操作数都是真，结果才为真，有一个为假，结果就为假
			"一假则假，全真才真"
	逻辑或：||
		运算规则：偶有的操作数都是假，结果才为假，有一个为真，结果就为真
			"一真则真，全假才假"
	逻辑非：！
		!T-->false
		!F-->true
	 */
	f1 := true
	f2 := false
	f3 := true
	res1 := f1 && f2
	fmt.Printf("res1: %t\n", res1)

	res2 := f1 && f2 && f3
	fmt.Printf("res2: %t\n", res2)

	res3 := f1 || f2
	fmt.Printf("res3: %t\n", res3)
	res4 := f1 || f2 || f3
	fmt.Printf("res4: %t\n", res4)
	fmt.Println(false || false || false)

	fmt.Printf("f1:%t,!f1:%t\n", f1, !f1)
	fmt.Printf("f2:%t,!f2:%t\n", f2, !f2)

	a := 3
	b := 2
	c := 5
	res5 := a > b && c%a == b && a < (c/b)
	fmt.Println(res5)

	res6 := b*2 < c || a/b != 0 || c/a > b
	fmt.Println(res6)
	res7:=!(c/a==b)
	fmt.Println(res7)

}
