package main

import (
	"os"
	"fmt"
)

func main() {
	f,err := os.Open("test.txt")
	if err != nil{
		//log.Fatal(err)
		fmt.Println(err) //open test.txt: no such file or directory
		if ins ,ok := err.(*os.PathError);ok{
			fmt.Println("1.Op:",ins.Op)
			fmt.Println("2.Path:",ins.Path)
			fmt.Println("3.Err:",ins.Err)
		}
		return
	}
	fmt.Println(f.Name(),"打开文件成功。。")

}
