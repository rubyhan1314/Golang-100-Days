package main

import "fmt"

func main() {
	/*
	位运算符：
		将数值，转为二进制后，按位操作
	按位&：
		对应位的值如果都为1才为1，有一个为0就为0
	按位|：
		对应位的值如果都是0才为0，有一个为1就为1
	异或^：
		二元：a^b
			对应位的值不同为1，相同为0
		一元：^a
			按位取反：
				1--->0
				0--->1
	位清空：&^
			对于 a &^ b
				对于b上的每个数值
				如果为0，则取a对应位上的数值
				如果为1，则结果位就取0

	位移运算符：
	<<：按位左移，将a转为二进制，向左移动b位
		a << b
	>>: 按位右移，将a 转为二进制，向右移动b位
		a >> b
	 */

	a := 60
	b := 13
	/*
	a: 60 0011 1100
	b: 13 0000 1101
	&     0000 1100
	|     0011 1101
	^	  0011 0001
	&^    0011 0000


	a : 0000 0000 ... 0011 1100
	^   1111 1111 ... 1100 0011
	 */
	fmt.Printf("a:%d, %b\n",a,a)
	fmt.Printf("b:%d, %b\n",b,b)

	res1 := a & b
	fmt.Println(res1) // 12

	res2 := a | b
	fmt.Println(res2) // 61

	res3 := a ^ b
	fmt.Println(res3) // 49

	res4 := a &^ b
	fmt.Println(res4) // 48

	res5 := ^a
	fmt.Println(res5)

	c:=8
	/*
	c : ... 0000 1000
	      0000 100000
	>>        0000 10
	 */
	res6 := c << 2
	fmt.Println(res6)

	res7 := c >> 2
	fmt.Println(res7)
}