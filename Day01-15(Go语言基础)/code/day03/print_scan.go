package main

import (
	"fmt"

	"os"
	"bufio"
)

func main() {
	/*
	输入和输出：
		fmt包：输入，输出

		输出：
			Print() //打印
			Printf() //格式化打印
			Println() //打印之后换行

		格式化打印占位符：
			%v,原样输出
			%T，打印类型
			%t,bool类型
			%s，字符串
			%f，浮点
			%d，10进制的整数
			%b，2进制的整数
			%o，8进制
			%x，%X，16进制
				%x：0-9，a-f
				%X：0-9，A-F
			%c，打印字符
			%p，打印地址
			。。。

		输入：
			Scanln()
				Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
			Scanf()

		bufio包

	 */
	//a := 100           //int
	//b := 3.14          //float64
	//c := true          // bool
	//d := "Hello World" //string
	//e := `Ruby`        //string
	//f := 'A'
	//fmt.Printf("%T,%b\n", a, a)
	//fmt.Printf("%T,%f\n", b, b)
	//fmt.Printf("%T,%t\n", c, c)
	//fmt.Printf("%T,%s\n", d, d)
	//fmt.Printf("%T,%s\n", e, e)
	//fmt.Printf("%T,%d,%c\n", f, f, f)
	//fmt.Println("-----------------------")
	//fmt.Printf("%v\n", a)
	//fmt.Printf("%v\n", b)
	//fmt.Printf("%v\n", c)
	//fmt.Printf("%v\n", d)
	//fmt.Printf("%v\n", e)
	//fmt.Printf("%v\n", f)
	//
	//
	//fmt.Println("-----------------------")


	//var x int
	//var y float64
	//fmt.Println("请输入一个整数，一个浮点类型：")
	//fmt.Scanln(&x,&y)//读取键盘的输入，通过操作地址，赋值给x和y   阻塞式
	//fmt.Printf("x的数值：%d，y的数值：%f\n",x,y)
	//
	//fmt.Scanf("%d,%f",&x,&y)
	//fmt.Printf("x:%d,y:%f\n",x,y)

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)

}
