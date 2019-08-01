package main

import "fmt"

func main() {
	/*
	字符串：
	1.概念：多个byte的集合，理解为一个字符序列
	2.语法：使用双引号
		"abc","hello","A"
			也可以使用``
	3.编码问题
			计算机本质只识别0和1
			A：65，B:66,C:67...
			a:97,b:98...
		ASCII(美国标准信息交换码)

		中国的编码表：gbk，兼容ASCII
			中
			家
		Unicode编码表：号称统一了全世界
			UTF-8，UTF-16,UTF-32...

	4.转义字符：\
		A：有一些字符，有特殊的作用，可以转义为普通的字符
			\',\'
		B：有一些字符，就是一个普通的字符，转义后有特殊的作用
			\n,换行
			\t,制表符
	 */
	 //1.定义字符串
	 var s1 string
	 s1 = "王二狗"
	 fmt.Printf("%T,%s\n",s1,s1)

	 s2 := `Hello World`
	 fmt.Printf("%T,%s\n",s2,s2)

	 //2.区别：'A',"A"
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%d\n",v1,v1) //int32
	fmt.Printf("%T,%s\n",v2,v2)

	v3 := '中'
	fmt.Printf("%T,%d,%c,%q\n",v3,v3,v3,v3)

	//3.转义字符
	fmt.Println("\"HelloWorld\"")
	fmt.Println("Hello\nWor\tld")

	fmt.Println(`He"lloWor"ld`)
	fmt.Println("Hello`Wor`ld")
}
