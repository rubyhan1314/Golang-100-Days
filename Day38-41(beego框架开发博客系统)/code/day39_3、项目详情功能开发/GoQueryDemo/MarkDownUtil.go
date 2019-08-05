package main

import (
	"io/ioutil"
	"github.com/russross/blackfriday"
	"fmt"
)

func main() {

	fileread, _ := ioutil.ReadFile("./test.md")
	//转换Markdown语法，如将"#"转换为"<h1></h1>"
	subHtml := blackfriday.MarkdownCommon(fileread)
	subHtmlStr := string(subHtml)
	fmt.Println(subHtmlStr)
}
