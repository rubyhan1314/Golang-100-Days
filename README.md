## Golang - 100天从新手到大师

> 作者：Steven，韩茹，Davie
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
#### Day02 - [基本语法](./Day01-15(Go语言基础)/day02_基本语法.md)
#### Day03 - [数据类型&运算符](./Day01-15(Go语言基础)/day03_数据类型&运算符.md)
#### Day04 - [分支语句](./Day01-15(Go语言基础)/day04_分支语句.md)
#### Day05 - [循环语句](./Day01-15(Go语言基础)/day05_循环语句.md)
#### Day06 - [数组](./Day01-15(Go语言基础)/day06_数组.md)
#### Day07 - [切片](./Day01-15(Go语言基础)/day07_Slice的使用.md)
#### Day08 - [Map](./Day01-15(Go语言基础)/day08_Map的使用.md)
#### Day09 - [string](./Day01-15(Go语言基础)/day09_string.md)
#### Day10 - [函数](./Day01-15(Go语言基础)/day10_函数.md)
#### Day11 - [包的管理](./Day01-15(Go语言基础)/day11_包的管理.md)
#### Day12 - [指针](./Day01-15(Go语言基础)/day12_指针.md)
#### Day13 - [结构体](./Day01-15(Go语言基础)/day13-结构体.md)
#### Day14 - [方法和接口](./Day01-15(Go语言基础)/day14_第1节_方法.md)
#### Day15 - [错误处理](./Day01-15(Go语言基础)/day15_错误处理.md)

### Day16~20 - [Go语言基础进阶](./Day16-20(Go语言基础进阶))
#### Day16 - [I/O操作](./goon.md)
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

### Day38~41 - [项目实战一](./Day38-41(项目实战一))

### Day42~43 - [Gin框架](./Day42-43(Gin框架))

### Day44 - [MySQL数据库高级](./Day44(MySQL数据库高级))

### Day45 - [Git](./Day45(Git))

### Day46~50 - [项目实战二](./Day46-50(项目实战二))

### Day51 - [Node.js](./Day51(Node.js))

### Day52 - [Vue](./Day52(Vue))

### Day53 - [Redis数据库](./Day53(Redis数据库))

### Day54~55 - [Iris框架](./Day54-55(Iris框架))

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


### Day76~95 - [微服务](./Day76-90(微服务))

#### Day76 - [微服务特性](./goon.md)

#### Day77 - [protobuf](./goon.md)

#### Day78~79 - [微服务管理](./goon.md)

#### Day80 - [RPC远程调用机制](./goon.md)

#### Day81~82 - [gRPC远程调用机制](./goon.md)

#### Day83~85 - [go-micro微服务框架](./goon.md)

#### Day86 - [RESTful](./goon.md)

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
