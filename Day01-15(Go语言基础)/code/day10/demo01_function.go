package main

import "fmt"

func main() { //程序的入口，是一个特殊的函数
	/*
	函数：function
	一、概念：
		具有特定功能的代码，可以被多次调用执行。
	二、意义：
		1.可以避免重复的代码
		2.增强程序的扩展性
	三、使用
		step1：函数的定义，也叫声明
		step2：函数的调用，就是执行函数中代码的过程
	四、语法
		1.定义函数的语法
			func funcName(parametername type1, parametername type2) (output1 type1, output2 type2) {
				//这里是处理逻辑代码
				//返回多个值
				return value1, value2
			}
			A:func，定义函数的关键字
			B:funcName，函数的名字
			C:()，函数的标志
			D:参数列表：形式参数用于接收外部传入函数中的数据
			E:返回值列表：函数执行后返回给调用处的结果

		2、调用函数的语法
			函数名(实际参数)

		函数的调用处，就是函数调用的位置

		3、注意事项
			A：函数必须先定义，再调用，如果不定义：undefined: getSum
				定义了函数，没有调用，那么函数就失去了意义

			B：函数名不能冲突
			C：main()，是一个特殊的函数，作为程序的入口，由系统自动调用
				而其他函数，程序中通过函数名来调用。
	 */
	//求1-10的和
	getSum() //函数的调用处

	fmt.Println("hello world....")

	//求1-10的和
	getSum()

	fmt.Println("aaa...")

	//求1-10的和
	getSum()
}

//定义一个函数：用于求1-10的和
func getSum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1-100的和是：%d\n", sum)
}



