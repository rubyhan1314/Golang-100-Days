package main

import (
	"strings"
	"fmt"
	"os"
	"log"
	"io"
	"strconv"
)

func main() {
	/*
	断点续传：
		文件传递：文件复制

		"/Users/ruby/Documents/pro/a/guliang.jpeg"

		复制到当前的工程下：

	思路：边复制，边记录复制的总量

	 */
	 srcFile := "/Users/ruby/Documents/pro/a/guliang.jpeg"
	 destFile := srcFile[strings.LastIndex(srcFile,"/")+1:]
	 fmt.Println(destFile)
	 tempFile := destFile+"temp.txt"
	 fmt.Println(tempFile)

	 file1,err := os.Open(srcFile)
	 HandleErr(err)
	 file2,err := os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	 HandleErr(err)
	 file3,err := os.OpenFile(tempFile,os.O_CREATE|os.O_RDWR,os.ModePerm)
	 HandleErr(err)

	 defer file1.Close()
	 defer file2.Close()

	 //step1：先读取临时文件中的数据，再seek
	 file3.Seek(0,io.SeekStart)
	 bs := make([]byte,100,100)
	 n1,err := file3.Read(bs)
	 //HandleErr(err)
	 countStr := string(bs[:n1])
	 count,err := strconv.ParseInt(countStr,10,64)
	 //HandleErr(err)
	 fmt.Println(count)


	 //step2：设置读，写的位置：
	 file1.Seek(count,io.SeekStart)
	 file2.Seek(count,io.SeekStart)
	 data := make([]byte,1024,1024)
	 n2 := -1 //读取的数据量
	 n3 := -1 //写出的数据量
	 total := int(count) //读取的总量

	 //step3：复制文件
	 for{
	 	n2,err =file1.Read(data)
	 	if err == io.EOF || n2 == 0{
	 		fmt.Println("文件复制完毕。。",total)
	 		file3.Close()
	 		os.Remove(tempFile)
	 		break
		}
	 	n3,err = file2.Write(data[:n2])
	 	total += n3

	 	//将复制的总量，存储到临时文件中
	 	file3.Seek(0,io.SeekStart)
	 	file3.WriteString(strconv.Itoa(total))

	 	fmt.Printf("total:%d\n",total)

	 	//假装断电
	 	//if total > 8000{
	 	//	panic("假装断电了。。。")
		//}


	 }



}

func HandleErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}