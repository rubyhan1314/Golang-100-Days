package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	/*
	ioutil包：
		ReadFile()
		WriteFile()
		ReadDir()
		..
	 */

	 //1.读取文件中的所有的数据
	// fileName := "/Users/ruby/Documents/pro/a/aa.txt"
	//data,err := ioutil.ReadFile(fileName)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//2.写出数据
	//fileName := "/Users/ruby/Documents/pro/a/bbb.txt"
	//s1 := "床前明月光，地上鞋三双"
	//err := ioutil.WriteFile(fileName,[]byte(s1),os.ModePerm)
	//fmt.Println(err)

	//3.ReadAll()
	//s2 := "王二狗和李小花是两个好朋友，Ruby就是我，也是他们的朋友"
	//r1:=strings.NewReader(s2)
	//data,err :=ioutil.ReadAll(r1)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//4.ReadDir()，读取一个目录下的子内容：子文件和子目录，但是只能读取一层
	//dirName := "/Users/ruby/Documents/pro"
	//fileInfos,err := ioutil.ReadDir(dirName)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(len(fileInfos))
	//for i:=0;i<len(fileInfos);i++{
	//	//fmt.Printf("%T\n",fileInfos[i])
	//	fmt.Printf("第 %d 个：名称：%s，是否是目录：%t\n",i,fileInfos[i].Name(),fileInfos[i].IsDir())
	//}

	//5.临时目录和临时文件
	dir,err := ioutil.TempDir("/Users/ruby/Documents/pro/a","Test")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	fmt.Println(dir)

	file,err := ioutil.TempFile(dir,"Test")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())

}
