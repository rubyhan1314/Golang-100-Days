# 一、字符串(string)

> @author：韩茹
> 
> 版权所有：北京千锋互联科技有限公司

## 1.1 什么是string

Go中的字符串是一个字节的切片。可以通过将其内容封装在“”中来创建字符串。Go中的字符串是Unicode兼容的，并且是UTF-8编码的。

示例代码：

```go
package main

import (  
    "fmt"
)

func main() {  
    name := "Hello World"
    fmt.Println(name)
}
```



## 1.2  string的使用

### 1.2.1 访问字符串中的单个字节

```go
package main

import (  
    "fmt"
)

func main() {  
    name := "Hello World"
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%d ", s[i])
    }
    fmt.Printf("\n")
    for i:= 0; i < len(s); i++ {
        fmt.Printf("%c ",s[i])
    }
}
```

运行结果：

72 101 108 108 111 32 87 111 114 108 100 
H e l l o   W o r l d 

## 1.3 strings包

访问strings包，可以有很多操作string的函数。



## 1.4 strconv包

访问strconv包，可以实现string和其他数值类型之间的转换。



千锋Go语言的学习群：784190273

作者B站：

https://space.bilibili.com/353694001

对应视频地址：

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

源代码：

https://github.com/rubyhan1314/go_foundation


