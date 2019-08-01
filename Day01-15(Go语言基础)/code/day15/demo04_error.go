package main

import (
	"path/filepath"
	"fmt"
)

func main() {
	files,_:=filepath.Glob("[")
	//if err != nil && err == filepath.ErrBadPattern{
	//	fmt.Println(err) //syntax error in pattern
	//	return
	//}
	fmt.Println("files:",files)
}
