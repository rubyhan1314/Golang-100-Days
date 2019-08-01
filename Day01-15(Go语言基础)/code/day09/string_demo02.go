package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	strings包下的关于字符串的函数

	 */
	s1 := "helloworld"
	//1.是否包含指定的内容-->bool
	fmt.Println(strings.Contains(s1, "abc"))
	//2.是否包含chars中任意的一个字符即可
	fmt.Println(strings.ContainsAny(s1, "abcd"))
	//3.统计substr在s中出现的次数
	fmt.Println(strings.Count(s1, "lloo"))

	//4.以xxx前缀开头，以xxx后缀结尾
	s2 := "20190525课堂笔记.txt"
	if strings.HasPrefix(s2, "201905") {
		fmt.Println("19年5月的文件。。")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档。。")
	}

	//索引
	//helloworld
	fmt.Println(strings.Index(s1, "lloo"))     //查找substr在s中的位置，如果不存在就返回-1
	fmt.Println(strings.IndexAny(s1, "abcdh")) //查找chars中任意的一个字符，出现在s中的位置
	fmt.Println(strings.LastIndex(s1, "l"))    //查找substr在s中最后一次出现的位置

	//字符串的拼接
	ss1 := []string{"abc","world","hello","ruby"}
	s3 := strings.Join(ss1,"-")
	fmt.Println(s3)

	//切割
	s4 := "123,4563,aaa,49595,45"
	ss2 := strings.Split(s4,",")
	//fmt.Println(ss2)
	for i := 0;i<len(ss2);i++{
		fmt.Println(ss2[i])
	}

	//重复，自己拼接自己count次
	s5 :=strings.Repeat("hello",5)
	fmt.Println(s5)

	//替换
	//helloworld
	s6 := strings.Replace(s1,"l","*",-1)
	fmt.Println(s6)
	//fmt.Println(strings.Repeat("hello",5))

	s7:="heLLo WOrlD**123.."
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))

	/*
	截取子串：
	substring(start,end)-->substr
	str[start:end]-->substr
		包含start，不包含end下标
	 */
	 fmt.Println(s1)
	 s8 := s1[:5]
	 fmt.Println(s8)
	 fmt.Println(s1[5:])
}
