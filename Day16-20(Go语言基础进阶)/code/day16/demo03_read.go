package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	/*
	读取数据：
		Reader接口：
			Read(p []byte)(n int, error)
	 */
	 //读取本地aa.txt文件中的数据
	 //step1：打开文件
	 fileName := "/Users/ruby/Documents/pro/a/aa.txt"
	 file,err := os.Open(fileName)
	 if err != nil{
	 	fmt.Println("err:",err)
	 	return
	 }
	 //step3：关闭文件
	 defer file.Close()

	 //step2：读取数据
	 bs := make([]byte,4,4)
	 /*
	 //第一次读取
	 n,err :=file.Read(bs)
	 fmt.Println(err) //<nil>
	 fmt.Println(n) //4
	 fmt.Println(bs) //[97 98 99 100]
	fmt.Println(string(bs)) //abcd

	//第二次读取
	n,err = file.Read(bs)
	fmt.Println(err)//<nil>
	fmt.Println(n)//4
	fmt.Println(bs) //[101 102 103 104]
	fmt.Println(string(bs)) //efgh

	//第三次读取
	n,err = file.Read(bs)
	fmt.Println(err) //<nil>
	fmt.Println(n) //2
	fmt.Println(bs) //[105 106 103 104]
	fmt.Println(string(bs)) //ijgh

	//第四次读取
	n,err = file.Read(bs)
	fmt.Println(err) //EOF
	fmt.Println(n) //0
	 */
	 n := -1
	 for{
	 	n,err = file.Read(bs)
	 	if n == 0 || err == io.EOF{
	 		fmt.Println("读取到了文件的末尾，结束读取操作。。")
	 		break
		}
	 	fmt.Println(string(bs[:n]))
	 }
}
