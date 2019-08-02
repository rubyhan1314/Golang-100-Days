package main

import (
	"os"
	"io"
	"fmt"
	"io/ioutil"
)

func main() {
	/*
	拷贝文件：
	 */
	srcFile := "/Users/ruby/Documents/pro/a/guliang.jpeg"
	destFile := "guliang3.jpeg"
	//total,err := CopyFile1(srcFile,destFile)
	//total,err := CopyFile2(srcFile,destFile)
	total,err := CopyFile3(srcFile,destFile)
	fmt.Println(total,err)
}

func CopyFile3(srcFile,destFile string)(int,error){
	bs,err := ioutil.ReadFile(srcFile)
	if err != nil{
		return 0,err
	}
	err = ioutil.WriteFile(destFile,bs,0777)
	if err != nil{
		return 0,err
	}
	return len(bs),nil
}

func CopyFile2(srcFile,destFile string)(int64,error){
	file1,err := os.Open(srcFile)
	if err != nil{
		return 0,err
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0,err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2,file1)
}


//该函数：用于通过io操作实现文件的拷贝，返回值是拷贝的总数量(字节),错误
func CopyFile1(srcFile,destFile string)(int,error){
	file1,err :=os.Open(srcFile)
	if err != nil{
		return 0,err
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0,err
	}
	defer file1.Close()
	defer file2.Close()

	//读写
	bs := make([]byte,1024,1024)
	n := -1 //读取的数据量
	total := 0
	for{
		n,err = file1.Read(bs)
		if err == io.EOF || n==0{
			fmt.Println("拷贝完毕。。")
			break
		}else if err != nil{
			fmt.Println("报错了。。")
			return total,err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total,nil
}

