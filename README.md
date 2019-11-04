## Golang - 100天从新手到大师

> 作者：韩茹，Davie，Steven
>
> 最近有很多小伙伴都在寻找go语言完整学习资料，但是录制视频和教程都需要不少的时间，平时也在筹备Go语言学科的事情，所以时间上比较紧张，我和Davie老师分别负责一部分Golang内容的产出。从技术文章，到视频，到项目代码。也都公布在各大平台上，但是知乎上就只方便看技术文章，B站上就只方便看视频。所以我们在github上上传了我们所有的学习资料，从最基础的入门到项目设计，希望帮助更多想了解和学习Go语言的伙伴，为方便大家交流学习。我们建了学习讨论群组(Go语言学习营：784190273)可以加群进行学习讨论。
>
> 因为是持续创作，所以也会持续更新。有些章节目录还没有内容，敬请期待。。
>
> 创作不易，感谢大家的支持。如果看后觉得有收获，可以打赏请作者喝咖啡吧，如果有疑问可以进群讨论。
>
> 最后感谢**千锋教育Go语言教学部**的鼎力支持。



![WechatIMG723_meitu_1](img/WechatIMG724_meitu_2.jpg)



### Go语言应用领域和就业分析

Go语言是谷歌2009年发布的第二款开源编程语言（系统开发语言)，它是基于编译、垃圾收集和并发的编程语言。
Go语言专门针对多处理器系统应用程序的编程进行了优化，使用Go编译的程序可以媲美 C / C++代码的速度，而且更加安全、支持并行进程。作为出现在21世纪的语言，其近C的执行性能和近解析型语言的开发效率，以及近乎于完美的编译速度，已经风靡全球。特别是在云项目中，大部分都使用了Golang来开发。不得不说，Golang早已深入人心。而对于一个没有历史负担的新项目，Golang或许就是个不二的选择。



**Golang的哲学理念：“Less is more or less is less”。**

 - 学习曲线容易
 - 效率：快速的编译时间，开发效率和运行效率高
 - "出身名门、血统纯正"
 - 自由高效：组合的思想、无侵入式的接口
 - 强大的标准库
 - 部署方便：二进制文件，Copy部署
 - 并行和异步编程几乎无痛点



**目前几个比较流行的领域，Go都有用武之地。**

 - 云计算基础设施领域

   代表项目：docker、kubernetes、etcd、consul、cloudflare CDN、七牛云存储等。

 - 基础软件

   代表项目[tidb、influxdb、cockroachdb等。

 - 微服务

   代表项目：go-kit、micro、monzo bank的typhon、bilibili等。

 - 互联网基础设施

   代表项目：以太坊、hyperledger等。

 - 分布式系统、数据库代理器、中间件等，例如Etcd

 - DevOps - Go / Python / Shell / Ruby 

   

**作为一名Go语言开发者，主要的就业领域包括：**

- Golang开发工程师 / Golang研发工程师
- Golang服务器后台开发 / 游戏服务器开发 
- 云计算平台(golang)开发工程师
- 区块链开发(golang)工程师
- Golang架构师



**给初学者的几个建议：**

- Make English as your working language.
- Practice makes perfect.
- All experience comes from mistakes.
- Don't be one of the leeches.
- Either stand out or kicked out.



### Day01~15 - [Go语言基础](./Day01-15(Go语言基础))

#### Day01 - [Go语言初识](./Day01-15(Go语言基础)/day01_第8节_第一个程序HelloWorld.md)

- Go语言简介 - Go语言的历史 / Go语言核心特性 / Go语言的logo版本 / Go的应用领域
- Go语言环境搭建 - Windows系统 / Linux系统 / MacOS系统
- HelloWorld程序 - Go语言文件的结构格式 / fmt包 / Print函数
- Go的执行原理 - Go的命令
- 安装IDE - Goland工具 / 其他的IDE
- 注释 - 注释的作用 / 单行注释 / 多行注释

- 编码规范

  

#### Day02 - [基本语法](./Day01-15(Go语言基础)/day02_基本语法.md)

- 进制以及转换 - 计算机原理 / 二进制 / 十进制 / 八进制 / 十六进制 / 进制转换

- 变量 - 变量的命名 / 变量的使用 / 变量的分析 / 变量的注意事项 

- 常量 - 常量的命名 / 常量的使用 / 常量的分析 / 常量的注意事项 / iota关键字

  

#### Day03 - [数据类型&运算符](./Day01-15(Go语言基础)/day03_数据类型&运算符.md)

- 数据类型 - 整数 / 浮点数 / 复数 / 字符串  / 字符编码 

- 数据类型的转换 - 强制转换 / 自动转换

- 运算符 - 算术运算符 / 赋值运算符 / 比较运算符 / 逻辑运算符 / 位运算符 / 位移运算符 /  运算符的优先级

- 表达式 - 表达式的值 / 表达式的类型

- 键盘输入和打印输出 - Scanln() / Scanf() / Print() / Printf() / Println()

- 格式化占位符 - %v / %T / %t / %s / %f / % d / %p / %c 。。。

  

#### Day04 - [分支语句](./Day01-15(Go语言基础)/day04_分支语句.md)

- 分支结构的应用场景 - 条件 / 结构 / 代码块 / 流程图

- if语句 - 简单的if / if-else结构 / if-elif-else结构 / 嵌套的if / if的其他写法

- switch语句 - switch的结构 / case语句 / break语句 / fallthrough语句 / switch的其他写法

  

#### Day05 - [循环语句](./Day01-15(Go语言基础)/day05_循环语句.md)

- 循环结构的应用场景 - 条件 / 结构 / 代码块 / 流程图
- for循环 - 基本结构  / 循环中的分支结构 / 嵌套的循环 / for的其他写法
- 循环控制语句 - break / continue
- goto语句 - goto的结构 / goto的注意点
- 生成随机数 - 随机数的生成
- 应用案例 - 1~100求和 / 判断素数 / 猜数字游戏 / 打印九九表 / 打印三角形图案 / 水仙花数 / 百钱百鸡

#### Day06 - [数组](./Day01-15(Go语言基础)/day06_数组.md)

- 数组 - 数组的概念 / 数组的使用 / 数组的语法 / 数组的长度 / 数组的注意点

- 数组的遍历 - 数组的下标 / 普通for遍历数组 / for…range 遍历

- 数组的排序 - 冒泡排序 / 选择排序 / 插入排序 。。。

- 多维数组 - 二维数组 / 二维数组的使用 / 二维数组的遍历

- 数组的数据类型 - 数组是值类型数据 

#### Day07 - [切片](./Day01-15(Go语言基础)/day07_Slice的使用.md)

- 切片 - 切片的概念 / 切片的使用 / 切片的语法 / 切片的长度和容量 / 切片的注意点
- 切片的原理 - 切片的底层数组 / 已有数组上创建切片 / 截取切片
- 切片的遍历 - 切片的下标 / 普通for遍历切片 / for…range 遍历
- 切片相关函数 - make() / append() / copy() /  len() / cap()
- 切片拷贝 - 深拷贝 / 浅拷贝

- 切片的数据类型 - 切片是引用类型数据 

#### Day08 - [Map](./Day01-15(Go语言基础)/day08_Map的使用.md)

- Map - Map的概念 / Map的使用 / Map的语法 / Map的长度和容量 / Map的注意点
- Map的存储特点 - key-value / Map的key类型 / Map中的键值对无序
- Map的创建 - 空Map / 
- Map的操作 - 添加数据 / 修改数据 / 获取数据 / 删除数据

- Map的数据类型 - Map是引用类型数据 

#### Day09 - [string](./Day01-15(Go语言基础)/day09_string.md)

- 字符串的使用 - 计算长度 / 下标运算 / 切片 / 常用方法
- strings包
- strconv包

#### Day10 - [函数](./Day01-15(Go语言基础)/day10_函数.md)

- 函数 - 函数的概念 / 函数的作用

- 函数的语法 - 定义函数 / 调用函数

- 函数的参数 - 参数的使用 / 可变参数 / 参数传递

- 函数的返回值 - 返回值 / return语句 / 没有返回值 / 返回单个值 / 返回多个值

- 变量作用域 - 局部变量 / 全局变量

- 递归函数 - 递归算法 / 递归函数实现

- 函数高级 - 函数的本质 / 匿名函数 / 高阶函数 / 回调函数 / 闭包结构

- defer函数 - 延迟函数 / 延迟参数 / 堆栈的延迟 / defer注意点

  

#### Day11 - [包的管理](./Day01-15(Go语言基础)/day11_包的管理.md)

- 包管理 - 包的概念 / 包的定义 / 包的导入 / main包

- 包的关键字  - package / import 

- 包的导入逻辑 - init()函数 / 同包的多个init()函数 / 不同包的多个init()函数

- 管理外部包

  

#### Day12 - [指针](./Day01-15(Go语言基础)/day12_指针.md)

- 指针 - 指针的概念 / 获取变量的地址 / 操作指针改变变量

- 指针的语法 - 定义指针 / * / 获取指针数值 /  & / 指针的指针

- 指针的注意点 - 指针的类型 / 指针的地址 / 空指针

- 指针的应用 - 指针作为参数 / 指针作为函数的返回值 / 数组指针和指针数组 / 指针函数和函数指针 

  

#### Day13 - [结构体](./Day01-15(Go语言基础)/day13-结构体.md)

- 结构体 - 结构体的概念 / 结构体的定义 / 结构体的初始化 / 结构体的访问 / make和new
- 结构体的匿名字段
- 结构体的嵌套 - 结构体嵌套 / 匿名嵌套 / 提升字段
- 结构体的使用 - 结构体指针 / 结构体作为函数的参数 / 结构体作为函数的返回值 



#### Day14 - [方法和接口](./Day01-15(Go语言基础)/day14_第1节_方法.md)
- 方法 - 方法的概念 / 方法的使用

- 方法的语法 - 方法接受者 / 方法和函数

- 结构体嵌套中的方法 - 方法的"继承" / 方法的"重写"

- 接口 - 接口的概念 / 接口的使用

- 接口的语法 - interface / 接口类型 

- 接口的应用 - 空接口 / 类型断言 / type关键字

  

#### Day15 - [错误处理](./Day01-15(Go语言基础)/day15_错误处理.md)

- 错误 - 错误的概念 / 错误的用法 /

- 错误的类型 - error接口 / 错误类型表示 / 自定义错误

- 错误的处理 - 返回错误的函数 / 错误的处理

- 相关知识 - panic()函数 / recover()函数 / defer()函数

  

### Day16~20 - [Go语言基础进阶](./Day16-20(Go语言基础进阶))

#### Day16 - [I/O操作](./goon.md)
- I / O  - 什么是I/O / os包的使用

- 文件操作 - 获取文件信息 / 文件创建

- I / O 操作 

- 相关包 - bufio包 / ioutil包

- I / O 应用 - 复制文件 / 断点续传 / 遍历文件夹

  

#### Day17 - [并发编程Goroutine](./goon.md)

#### Day18 - [通道Channel](./goon.md)
#### Day19 - [反射机制](./goon.md)
#### Day20 - [综合练习](./goon.md)

### Day21~22 - [网络编程](./Day21-22(网络编程))

### Day23~24 - [MySQL数据库基础](./Day23-24(MySQL数据库基础))

### Day25 - [Go语言链接MySQL](./Day25(Go链接MySQL))

### Day26~31 - [Web前端](./Day26-31(Web前端))
#### Day26 - [HTML](./goon.md)
#### Day27 - [CSS](./goon.md)
#### Day28~30 - [JavaScript](./goon.md)
#### Day31 - [jQuery](./goon.md)

### Day32~35 - [Go Web开发](./Day32-35(Go Web开发))
#### Day32 - [Web初识](./goon.md)
#### Day33 - [http包详解](./goon.md)
#### Day34 - [session和cookie](./goon.md)
#### Day35 - [文本处理](./goon.md)


### Day36~37 - [beego框架](./Day36-37(beego框架))
#### Day36 - [beego框架介绍和流程分析]()
- beego框架 - beego简介 / beego安装 / beego特性

- bee工具 -  bee简介/ bee安装

- bee的用法 - bee命令

- beego程序流程分析 - beego程序入口 / go语言执行顺

- beego框架功能 - 请求拦截 / 路由分发 

- beego控制器 - 处理逻辑

- beego.Run方法 - 解析配置 / 路由分发 / 监听服务

#### Day37 - [beego框架总结和数据库连接配置]()
- conf配置 - 项目数据配置 / 配置数据读取

- controllers -  控制器介绍 / 控制器功能 / 控制器定义

- models- 数据层作用 / model定义

- routers - 路由层功能 / 路由分类

- 静态资源 - 静态资源目录作用 / 静态资源路径设置 

- 数据库安装及配置 - mysql数据库安装 / mysql数据库基本命令 / 可视化工具 

- 数据库驱动 - 数据库驱动分类 / mysql驱动安装 / 连接配置 / 连接数据库 

### Day38~41 - [项目实战一](./Day38-41(beego框架开发博客系统)/day38_项目搭建、登录注册和Session功能开发.md)
#### Day38 - [项目搭建、登录注册和Session功能开发](./Day38-41(beego框架开发博客系统)/day38_项目搭建、登录注册和Session功能开发.md)
- bee工具使用 - 项目创建 / 项目运行

- mysql数据库操作 -  数据库连接配置 / 读取数据库配置 / 连接数据库

- models- 数据库表设计 / 执行数据库操作方法封装

- 用户注册 - 视图表单数据提交 / 服务器接收post数据 / 操作数据库表添加数据

- 用户登录 - 登录功能控制器 / 路由注册 / 服务器接收Post数据 / 数据库表条件查询  

- Session处理 - session功能启用配置 / 添加session数据 / 获取session数据 / 删除session数据

- BaseController - BaseController作用 / controller方法执行顺序 / Parepare方法作用 
 
#### Day39 - [写文章、项目首页和查看文章详情功能开发](./Day38-41(beego框架开发博客系统)/day39_写文章、项目首页和查看文章详情功能开发.md)
- model层 - 数据库表设计 / 数据库添加操作

- 控制器 -  写文章功能控制器定义 / get方法解析html页面 / Post方法接收form表单数据 

- 路由层 - 注册路由解析

- 视图层 - html功能页面 / js逻辑判断 / js表单提交

- 分页功能 - 分页设计 / 数据库limit操作实现分页查询 

- 首页内容显示 - model转换 / 分页视图功能

- Markdown语法 - 常见第三方库 / markdown格式编程语法 / markdown与html转换

#### Day40 - [写文章、项目首页和查看文章详情功能开发](./Day38-41(beego框架开发博客系统)/day40_修改文章、删除文章和文章标签功能开发.md)
- 修改文章功能 - 注册修改功能路由 / 修改功能controller逻辑实现 / get方法渲染页面 / post方法接收表单数据 / 数据库表数据修改操作

- 删除功能 - 按条件删除数据库表数据 / 删除后视图重定向

- 标签功能 - 数据库查询

#### Day41 - [写文章、项目首页和查看文章详情功能开发](./Day38-41(beego框架开发博客系统)/day41_首页功能扩展、图片上传和关于功能开发.md)
- 功能扩展 - tags查询 / page查询 / 多条件逻辑判断 / 多条件查询&的使用 

- 文件上传 - 数据表设计 / js实现文件提交 / 服务器接收文件数据 / 文件类型判断 / 文件大小判断 / 文件名 / 保存文件 

- 项目总结 - beego框架组成 / bee调试工具 / beego程序执行流程 / 数据库操作 / beego项目架构 / session处理 / 模板文件语法

### Day42~43 - [Gin框架](./Day42-43(Gin框架))

### Day44 - [MySQL数据库高级](./Day44(MySQL数据库高级))

### Day45 - [Git](./Day45(Git))

### Day46~50 - [项目实战二](./Day46-50(项目实战二))

### Day51 - [Node.js](./Day51(Node.js))

### Day52 - [Vue](./Day52(Vue))

### Day53 - [Redis数据库](./Day53(Redis数据库))

### Day54~55 - [Iris框架](./Day54-55(iris框架)/day54_web开发介绍、iris框架安装、HTTP请求和返回、iris路由处理.md)
#### Day54 - [web开发介绍、iris框架安装、HTTP请求和返回、iris路由处理](./Day54-55(iris框架)/day54_web开发介绍、iris框架安装、HTTP请求和返回、iris路由处理.md)
- web开发 - 项目架构 / 开发流程 / 实战项目介绍 / 项目技术栈 

- iris框架 - iris简介 / iris特性 / iris框架安装 / iris参考资料 

- http请求和处理 - 数据请求和分类 / http1.0和http1.1 / iris标准请求处理 / 自定义请求处理 / 请求处理数据格式封装 

- 路由处理 - Context概念 / 正则表达式路由

#### Day55 - [框架设置、mvc包、session使用、项目搭建和资源导入](./Day54-55(iris框架)/day55_框架设置、mvc包、session使用、项目搭建和资源导入.md)
- 路由组 - Party实现路由组 / context.Next()方法 / taml配置文件 / yaml配置文件设置 / 自定义配置文件 / 自定义配置设置

- mvc包 - mvc.Application作用 / mvc特性 / 声明周期 / mvc.Configure配置

- session处理和使用 - session与cookie区别 / session支持数据类型 / session创建 / session使用

- 实战项目 - 创建企业管理平台项目 / 目录说明 / 项目资源集成

### Day56~60 - [项目实战三](./Day56-60(项目实战三))

### Day61 - [Linux](./Day61(Linux))

### Day62~64 - [容器](./Day62-64(容器))
#### Day62 - [虚拟化VS容器化](./goon.md)
#### Day63 - [Docker](./goon.md)
#### Day64 - [Kubernetes(k8s)](./goon.md)

### Day65~75 - [分布式](./Day65-75(分布式))
#### Day65 - [分布式理论](./goon.md)
#### Day66~67 - [分布式文件系统FastDFS](./goon.md)
#### Day68 - [Nginx与反响代理部署](./goon.md)
#### Day69~70 - [Go开发实现高可用性etcd系统](./goon.md)
#### Day60~75 - [项目实战四：分布式项目](./goon.md)


### Day76~95 - [微服务](./Day76(微服务特性)/day76_微服务简介及特性介绍.md)

#### Day76 - [微服务简介、单体应用、微服务方案、微服务实践问题](./Day76(微服务特性)/day76_微服务简介及特性介绍.md)
- 微服务简介 - 构建单体应用 / 互联网公司架构

- 单体应用困境 - 应用发展趋势 / 困难与困境 

- 复杂问题解决方案 - 微服务架构 / 系统扩展性 / 微服务的优点 / 微服务不足 

- 微服务实践  - 微服务的定义 / 微服务实践要解决的问题 / 常见的微服务架构

#### Day77 - [Protobuf介绍、编程实现、Protobuf语法](./Day77(protobuf)/day77_Protobuf介绍.md)
- Protobuf介绍 - Protobuf简介 / Protobuf应用场景 / Protobuf 优点 / Protobuf 缺点
- 编程实现Protobuf传输 - 环境准备 / Protobuf 协议语法 / 使用Protobuf的步骤 
- Protobuf语法 - Protobuf 协议语法 / Protobuf序列化原理

#### Day78~79 - [微服务管理](./Day78-79(微服务管理)/day78_微服务管理(上).md)
##### Day78 - [服务发现原理、Consul搭建及配置、服务的管理、集群搭建](./Day78-79(微服务管理)/day78_微服务管理(上).md)
- 服务发现 / 使用服务发现原因 / 服务发现方案 / Consul介绍 / 分布式与单点故障 / Consul环境配置 / Consul内部原理 / Consul启动 / Consul节点发现

##### Day79 - [定义微服务、服务查询、集群搭建](./Day78-79(微服务管理)/day79_微服务管理(下).md)
- 服务管理 / 服务的定义 / 服务的注册与发现 / 服务查询 / 注册多个服务 / 多节点服务集群 / Docker环境的搭建 

#### Day80 - [ RPC介绍、RPC编程与实现、RPC与Protobuf结合](./Day80(RPC远程调用机制)/day80_RPC远程过程调用.md)
- RPC介绍 - RPC简介及原理介绍 / 本地过程调用 / RPC技术实现 

- RPC编程与实现 - RPC官方库 / net/rpc库实现RPC调用编程 

- RPC与Protobuf结合 - 传输数据格式定义 / Protobuf格式数据与RPC结合 

#### Day81~82 - [gRPC远程过程调用](./Day81-82(gRPC远程调用机制)/day81_gRPC远程调用机制介绍.md)

##### Day81 - [gRPC介绍和安装、gRPC框架的使用](./Day81-82(gRPC远程调用机制)/day81_gRPC远程调用机制介绍.md)

- gRPC介绍和安装 - 什么是gRPC / grpc-go介绍 

- gRPC框架使用 - 定义服务 / 编译.proto文件 / gRPC实现RPC编程 

##### Day82 - [gRPC编程使用、TLS验证和Token认证、拦截器及自定义拦截器的使用](./Day81-82(gRPC远程调用机制)/day82_gRPC使用.md)
- gRPC编程使用 - 服务端流RPC / 编译.protoc文件 / 服务端代码 / 客户端代码生成 / 服务的注册监听处理 / 客户端数据接收 / 客户端流模式 / 双向流模式

- TLS验证和Token认证 - 授权认证 / SSL/TLS工作原理 / Token认证过程 / 自定义Token认证

- 拦截器的使用 - Interceptor介绍 / 自定义UnaryServerInterceptor / 拦截器注册

#### Day83~85 - [go-micro微服务框架](./Day83-85(go-micro微服务框架)/day83_go-micro框架介绍.md)

##### Day83 - [micro框架介绍及使用、micro创建微服务](./Day83-85(go-micro微服务框架)/day83_go-micro框架介绍.md)

- micro框架介绍 - 背景 / micro概述介绍 / micro组成 / micro工具组件：API、Web、Sidecar、Bot / Go-Micro特性 / consul安装及环境准备

- 创建微服务 - 微服务的定义 / 初始化服务实例 / Options可选项配置 / 定义服务接口,实现服务业务逻辑 / 微服务调用 / 自定义指定服务发现组件

##### Day84 - [心跳机制与可选项配置、事件驱动机制](./Day83-85(go-micro微服务框架)/day84_go-micro使用(一).md)

- 心跳机制与可选项配置 - 心跳机制的由来 / consul心跳配置 / TTL和间隔时间 

- 事件驱动机制 - 发布/订阅机制 / Broker组件设计 / go-plugins安装 / MQTT介绍及环境搭建 / 编程实现消息订阅和发布

##### Day85 - [micro框架Selector机制](./Day83-85(go-micro微服务框架)/day85_go-micro使用(二).md)

- Selector机制 - 负载均衡算法 / Mico的Selector / Selector的定义 / DefaultSelector / registrySelector

#### Day86 - [RESTful设计和使用](./Day86(RESTful设计)/day86_RESTful标准设计.md)

- RESTful设计规范及使用 - Go-Micro API网关 / 安装Micro工具 / Micro API工作原理 / 反向代理的API服务启动 / REST 映射 

#### Day87 - [微服务项目设计](./goon.md)

#### Day88 - [RPC远程调用机制](./goon.md5)

#### Day89~95 - [项目实战五：微服务项目](./goon.md)

### Day96~100 - [完美收官](./Day96-100(完美收官))

#### Day96~97 - [项目部署和性能调优](./goon.md)

#### Day98 - [项目总结](./goon.md)

#### Day99 - [面试指导](./goon.md)

#### Day100 - [英文面试](./goon.md)





> 致谢：
>
> ​	感谢的千锋教育以及千锋教育Go语言组的同事：Steven老师，Davie老师等在技术上给与的知道和帮助。# Golang-100-Days
