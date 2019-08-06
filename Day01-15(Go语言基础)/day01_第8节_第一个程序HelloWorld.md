# 第一个程序：HelloWorld

> @author：韩茹
>
> 版权所有：北京千锋互联科技有限公司

##  一、go项目工程结构

配置好工作目录后，就可以编码开发了，在这之前，我们看下go的通用项目结构，这里的结构主要是源代码相应地资源文件存放目录结构。

### 1.1 gopath目录

gopath目录就是我们存储我们所编写源代码的目录。该目录下往往要有3个子目录：src，bin，pkg。

> src ---- 里面每一个子目录，就是一个包。包内是Go的源码文件
>
> pkg ---- 编译后生成的，包的目标文件
>
> bin ---- 生成的可执行文件。



### 1.2 编写第一个程序

每个编程语言的学习，都是从一个"Hello, World."程序开始的，这个例子首次出现于1978年出版的C语言圣经《The C Programming Language》。关于"Hello, World."还有一个很美好的故事，那就是所有的程序员期待着计算机有一天能拥有真正的智能，然后对创造他的人们"发自内心"的说一句，Hello, World。



1.在HOME/go的目录下，(就是GOPATH目录里)，创建一个目录叫src，然后再该目录下创建一个文件夹叫hello，在该目录下创建一个文件叫helloworld.go，并双击打开，输入以下内容：

```go
package main

import "fmt"

func main() {
   fmt.Println("Hello, World!")
}
```

2.执行go程序

执行go程序由几种方式

方式一：使用go run命令

​	step1：打开终端：

​			window下使用快捷键win+R，输入cmd打开命令行提示符

​			linux下可以使用快捷键：ctrl+alt+T

​			mac下command+空格，输入termainl

​	step2：进入helloworld.go所在的目录

​	step3：输入go run helloworld.go命令并观察运行结果。

方式二：使用go build命令

​	step1：打开终端：在任意文件路径下，运行:
​			go install hello 

​		也可以进入项目(应用包)的路径，然后运行：
​			go install 

注意，在编译生成go程序的时，go实际上会去两个地方找程序包：
GOROOT下的src文件夹下，以及GOPATH下的src文件夹下。

在程序包里，自动找main包的main函数作为程序入口，然后进行编译。

​	step2：运行go程序
​		在/home/go/bin/下(如果之前没有bin目录则会自动创建)，会发现出现了一个hello的可执行文件，用如下命令运行:
​		./hello

​	

### 1.3 第一个程序的解释说明

#### 3.2.1 package

- 在同一个包下面的文件属于同一个工程文件，不用`import`包，可以直接使用
- 在同一个包下面的所有文件的package名，都是一样的
- 在同一个包下面的文件`package`名都建议设为是该目录名，但也可以不是

#### 3.2.2  import

import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包的函数，fmt 包实现了格式化 IO（输入/输出）的函数

可以是相对路径也可以是绝对路径，推荐使用绝对路径（起始于工程根目录）

1. 点操作
   我们有时候会看到如下的方式导入包

   ```go
   import(
   	. "fmt"
   ) 
   ```

   这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调

   用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`

2. 别名操作
   别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字

   ```go
   import(
   	f "fmt"
   ) 
   ```

   别名操作的话调用包函数时前缀变成了我们的前缀，即`f.Println("hello world")`

3. _操作
   这个操作经常是让很多人费解的一个操作符，请看下面这个import

   ```go
   import (
     "database/sql"
     _ "github.com/ziutek/mymysql/godrv"
   ) 
   ```

   _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数

#### 3.3.3 main

main(),是程序运行的入口。





### 1.4 包的说明

我们知道源代码都是存放在GOPATH的src目录下，那么多个多个项目的时候，怎么区分呢？答案是通过包，使用包来组织我们的项目目录结构。有过java开发的都知道，使用包进行组织代码，包以网站域名开头就不会有重复，比如千锋的网站是`http://www.mobiletrain.org`，我们就可以以`mobiletrain.org`的名字创建一个文件夹，我自己的go项目都放在这个文件夹里，这样就不会和其他人的项目冲突，包名也是唯一的。

如果有自己的域名，那也可以使用自己的域名。如果没有个人域名，现在流行的做法是使用你个人的github名，因为每个人的是唯一的，所以也不会有重复。

![package1](http://7xtcwd.com1.z0.glb.clouddn.com/package1.png)



如上，src目录下跟着一个个域名命名的文件夹。再以github.com文件夹为例，它里面又是以github用户名命名的文件夹，用于存储属于这个github用户编写的go源代码。





千锋Go语言的学习群：784190273

作者B站：

https://space.bilibili.com/353694001

对应视频地址：

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

源代码：

https://github.com/rubyhan1314/go_foundation