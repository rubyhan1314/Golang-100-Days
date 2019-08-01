package main

import "fmt"

func main() {
	/*
	1.switch的标注写法：
	switch 变量{
	case 数值1：分支1
	case 数值2：分支2
	。。。
	default：
		最后一个分支
	}
	2.省略switch后的变量，相当于直接作用在true上
	switch{//true
	case true:
	case false:
	}

	3.case后可以同时跟随多个数值
	switch 变量{
	case 数值1，数值2，数值3：

	case 数值4，数值5：

	}

	4.switch后可以多一条初始化语句
	switch 初始化语句;变量{
	}
	 */
	switch {
	case true:
		fmt.Println("true..")
	case false:
		fmt.Println("false...")
	}
	/*
	成绩：
	[0-59]，不及格
	[60,69],及格
	[70,79],中
	[80,89]，良好
	[90,100],优秀
	 */
	score := 88
	switch {
	case score >= 0 && score < 60:
		fmt.Println(score, "不及格")
	case score >= 60 && score < 70:
		fmt.Println(score, "及格")
	case score >= 70 && score < 80:
		fmt.Println(score, "中等")
	case score >= 80 && score < 90:
		fmt.Println(score, "良好")
	case score >= 90 && score <= 100:
		fmt.Println(score, "优秀")
	default:
		fmt.Println("成绩有误。。。")

	}

	fmt.Println("---------------------")
	letter := ""
	switch letter {
	case "A", "E", "I", "O", "U":
		fmt.Println(letter, "是元音。。")
	case "M", "N":
		fmt.Println("M或N。。")
	default:
		fmt.Println("其他。。")
	}
	/*
	一个月的天数
	1，3，5，7，8，10，12
		31
	4，6，9，11
		30
	2：29/28
	 */
	month := 9
	day := 0
	year := 2019
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		day = 31

	case 4, 6, 9, 11:
		day = 30
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Println("月份有误。。")
	}
	fmt.Printf("%d 年 %d 月 的天数是：%d\n", year, month, day)
	fmt.Println("--------------------------")

	switch language := "golang";language {
	case "golang":
		fmt.Println("Go语言。。")
	case "java":
		fmt.Println("Java语言。。")
	case "python":
		fmt.Println("Python语言。。")
	}
	//fmt.Println(language) //undefined: language
}
