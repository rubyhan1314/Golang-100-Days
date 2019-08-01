package main

import "fmt"

func main() {
	/*
	if...else语句
		if 条件{
			//条件成立，执行此处的代码。。
			A段
		}else{
			//条件不成立，执行此处的代码。。
			B段
		}

	注意点：
		1.if后的{，一定是和if条件写在同一行的
		2.else一定是if语句}之后，不能自己另起一行
		3.if和else中的内容，二者必选其一的执行
	 */
	//给定一个成绩，如果大于等于60，就是及格，否则呢就是不及格
	/*
	score := 0
	fmt.Println("请输入您的成绩：")
	fmt.Scanln(&score)

	if score >= 60{
		fmt.Println(score,"及格。。")
	}else{
		fmt.Println(score,"不及格。。")
	}
*/
	//给定性别，如果是男，就去男厕所，否则去女厕所
	sex := "女"//bool, int, string

	if sex == "男"{
		fmt.Println("可以去男厕所啦。。。")
	} else{
		fmt.Println("你去女厕所吧。。")
	}

	fmt.Println("main...over....")
}
