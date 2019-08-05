package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

func main() {

	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		//Find(".title a")与Find(".title").Find("a")一样
		fmt.Println("第", i+1, "个帖子的标题：", title)
		//ret,_ := contentSelection.Html()
		//fmt.Printf("\n\n\n%v", ret)
		//fmt.Println(contentSelection.Text())
	})
	//最终输出为 html 文档：
	//new, err := doc.Html()
	//fmt.Println(new)
}
