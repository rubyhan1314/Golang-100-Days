# 面向对象(OOP)

> @author：韩茹
> 版权所有：北京千锋互联科技有限公司



go并不是一个纯面向对象的编程语言。在go中的面向对象，结构体替换了类。

Go并没有提供类class，但是它提供了结构体struct，方法method，可以在结构体上添加。提供了捆绑数据和方法的行为，这些数据和方法与类类似。

## 1.1 定义结构体和方法

通过以下代码来更好的理解，首先在src目录下创建一个package命名为oop，在oop目录下，再创建一个子目录命名为employee，在该目录下创建一个go文件命名为employee.go。

目录结构：oop -> employee -> employee.go

在employee.go文件中保存以下代码：

```go
package employee

import (  
    "fmt"
)

type Employee struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
}

func (e Employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
```

然后在oop目录下，创建文件并命名为main.go，并保存以下内容

```go
package main

import "oop/employee"

func main() {  
    e := employee.Employee {
        FirstName: "Sam",
        LastName: "Adolf",
        TotalLeaves: 30,
        LeavesTaken: 20,
    }
    e.LeavesRemaining()
}
```

运行结果：

```
Sam Adolf has 10 leaves remaining 
```

## 1.2 New()函数替代了构造函数

我们上面写的程序看起来不错，但是里面有一个微妙的问题。让我们看看当我们用0值定义employee struct时会发生什么。更改main的内容。转到下面的代码，

```go
package main

import "oop/employee"

func main() {  
    var e employee.Employee
    e.LeavesRemaining()
}
```

运行结果：

```
has 0 leaves remaining
```

通过运行结果可以知道，使用Employee的零值创建的变量是不可用的。它没有有效的名、姓，也没有有效的保留细节。在其他的OOP语言中，比如java，这个问题可以通过使用构造函数来解决。使用参数化构造函数可以创建一个有效的对象。



go不支持构造函数。如果某个类型的零值不可用，则程序员的任务是不导出该类型以防止其他包的访问，并提供一个名为NewT(parameters)的函数，该函数初始化类型T和所需的值。在go中，它是一个命名一个函数的约定，它创建了一个T类型的值给NewT(parameters)。这就像一个构造函数。如果包只定义了一个类型，那么它的一个约定就是将这个函数命名为New(parameters)而不是NewT(parameters)。



更改employee.go的代码：

首先修改employee结构体为非导出，并创建一个函数New()，它将创建一个新Employee。代码如下：

```go
package employee

import (  
    "fmt"
)

type employee struct {  
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}

func (e employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
```

我们在这里做了一些重要的改变。我们已经将Employee struct的起始字母e设置为小写，即我们已经将类型Employee struct更改为type Employee struct。通过这样做，我们成功地导出了employee结构并阻止了其他包的访问。将未导出的结构的所有字段都导出为未导出的方法是很好的做法，除非有特定的需要导出它们。由于我们不需要在包之外的任何地方使用employee struct的字段，所以我们也没有导出所有字段。

由于employee是未导出的，所以不可能从其他包中创建类型employee的值。因此，我们提供了一个输出的新函数。将所需的参数作为输入并返回新创建的employee。

这个程序还需要做一些修改，让它能够工作，但是让我们运行这个程序来了解到目前为止变化的效果。如果这个程序运行，它将会失败，有以下编译错误，

```
go/src/constructor/main.go:6: undefined: employee.Employee  
```

这是因为我们有未导出的Employee，因此编译器抛出错误，该类型在main中没有定义。完美的。正是我们想要的。现在没有其他的包能够创建一个零值的员工。我们成功地防止了一个无法使用的员工结构价值被创建。现在创建员工的唯一方法是使用新功能。

修改main.go代码

```go
package main  

import "oop/employee"

func main() {  
    e := employee.New("Sam", "Adolf", 30, 20)
    e.LeavesRemaining()
}
```

运行结果：

```
Sam Adolf has 10 leaves remaining 
```

因此，我们可以明白，虽然Go不支持类，但是结构体可以有效地使用，在使用构造函数的位置，使用New(parameters)的方法即可。

## 1.3组成(Composition )替代了继承(Inheritance)

Go不支持继承，但它支持组合。组合的一般定义是“放在一起”。构图的一个例子就是汽车。汽车是由轮子、发动机和其他各种部件组成的。

博客文章就是一个完美的组合例子。每个博客都有标题、内容和作者信息。这可以用组合完美地表示出来。

### 1.3.1 通过嵌入结构体实现组成

可以通过将一个struct类型嵌入到另一个结构中实现。

示例代码：

```go
package main

import (  
    "fmt"
)

/*
我们创建了一个author struct，它包含字段名、lastName和bio。我们还添加了一个方法fullName()，将作者作为接收者类型，这将返回作者的全名。
*/
type author struct {  
    firstName string
    lastName  string
    bio       string
}

func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}
/*
post struct有字段标题、内容。它还有一个嵌入式匿名字段作者。这个字段表示post struct是由author组成的。现在post struct可以访问作者结构的所有字段和方法。我们还在post struct中添加了details()方法，它打印出作者的标题、内容、全名和bio。
*/
type post struct {  
    title     string
    content   string
    author
}

func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.author.fullName())
    fmt.Println("Bio: ", p.author.bio)
}

func main() {  
    author1 := author{
        "Naveen",
        "Ramanathan",
        "Golang Enthusiast",
    }
    post1 := post{
        "Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,
    }
    post1.details()
}

```

运行结果：

```
Title:  Inheritance in Go  
Content:  Go supports composition instead of inheritance  
Author:  Naveen Ramanathan  
Bio:  Golang Enthusiast  
```

**嵌入结构体的切片**

在以上程序的main函数下增加以下代码，并运行

```go
type website struct {  
        []post
}
func (w website) contents() {  
    fmt.Println("Contents of Website\n")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}
```

运行报错：

```
main.go:31:9: syntax error: unexpected [, expecting field name or embedded type 
```

这个错误指向structs []post的嵌入部分。原因是不可能匿名嵌入一片。需要一个字段名。我们来修正这个错误，让编译器通过。

```go
type website struct {  
        posts []post
}
```

现在让我们修改的main函数,为我们的新的website创建几个posts。修改完完整代码如下：

```go
package main

import (  
    "fmt"
)

type author struct {  
    firstName string
    lastName  string
    bio       string
}

func (a author) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {  
    title   string
    content string
    author
}
func (p post) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.fullName())
    fmt.Println("Bio: ", p.bio)
}

type website struct {  
 posts []post
}
func (w website) contents() {  
    fmt.Println("Contents of Website\n")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}
func main() {  
    author1 := author{
        "Naveen",
        "Ramanathan",
        "Golang Enthusiast",
    }
    post1 := post{
        "Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,
    }
    post2 := post{
        "Struct instead of Classes in Go",
        "Go does not support classes but methods can be added to structs",
        author1,
    }
    post3 := post{
        "Concurrency",
        "Go is a concurrent language and not a parallel one",
        author1,
    }
    w := website{
        posts: []post{post1, post2, post3},
    }
    w.contents()
}   
```

运行结果：

```
Contents of Website

Title:  Inheritance in Go  
Content:  Go supports composition instead of inheritance  
Author:  Naveen Ramanathan  
Bio:  Golang Enthusiast

Title:  Struct instead of Classes in Go  
Content:  Go does not support classes but methods can be added to structs  
Author:  Naveen Ramanathan  
Bio:  Golang Enthusiast

Title:  Concurrency  
Content:  Go is a concurrent language and not a parallel one  
Author:  Naveen Ramanathan  
Bio:  Golang Enthusiast  
```

## 1.4 多态性(Polymorphism)

Go中的多态性是在接口的帮助下实现的。正如我们已经讨论过的，接口可以在Go中隐式地实现。如果类型为接口中声明的所有方法提供了定义，则实现一个接口。让我们看看在接口的帮助下如何实现多态。

任何定义接口所有方法的类型都被称为隐式地实现该接口。

类型接口的变量可以保存实现接口的任何值。接口的这个属性用于实现Go中的多态性。

举个例子，一个虚构的组织有两种项目的收入:固定的账单和时间和材料。组织的净收入是由这些项目的收入之和计算出来的。为了保持本教程的简单，我们假设货币是美元，我们不会处理美分。它将使用整数来表示。

首先我们定义一个接口：Income

```go
type Income interface {  
    calculate() int
    source() string
}
```

接下来，定义两个结构体：FixedBilling和TimeAndMaterial

```go
type FixedBilling struct {  
    projectName string
    biddedAmount int
}
```

```go
type TimeAndMaterial struct {  
    projectName string
    noOfHours  int
    hourlyRate int
}
```

下一步是定义这些结构体类型的方法，计算并返回实际收入和收入来源。

```go
func (fb FixedBilling) calculate() int {  
    return fb.biddedAmount
}

func (fb FixedBilling) source() string {  
    return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {  
    return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {  
    return tm.projectName
}
```

接下来，我们来声明一下计算和打印总收入的calculateNetIncome函数。

```go
func calculateNetIncome(ic []Income) {  
    var netincome int = 0
    for _, income := range ic {
        fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
        netincome += income.calculate()
    }
    fmt.Printf("Net income of organisation = $%d", netincome)
}
```

上面的calculateNetIncome函数接受一部分Income接口作为参数。它通过遍历切片和调用calculate()方法来计算总收入。它还通过调用source()方法来显示收入来源。根据收入接口的具体类型，将调用不同的calculate()和source()方法。因此，我们在calculateNetIncome函数中实现了多态。

在未来，如果组织增加了一种新的收入来源，这个函数仍然可以正确地计算总收入，而没有一行代码更改。

最后我们写以下主函数：

```go
func main() {  
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
    incomeStreams := []Income{project1, project2, project3}
    calculateNetIncome(incomeStreams)
}
```

运行结果：

```
Income From Project 1 = $5000  
Income From Project 2 = $10000  
Income From Project 3 = $4000  
Net income of organisation = $19000  
```



假设该组织通过广告找到了新的收入来源。让我们看看如何简单地添加新的收入方式和计算总收入，而不用对calculateNetIncome函数做任何更改。由于多态性，这样是可行的。

首先让我们定义Advertisement类型和calculate()和source()方法。

```go
type Advertisement struct {  
    adName     string
    CPC        int
    noOfClicks int
}

func (a Advertisement) calculate() int {  
    return a.CPC * a.noOfClicks
}

func (a Advertisement) source() string {  
    return a.adName
}
```

广告类型有三个字段adName, CPC(cost per click)和noof点击数(cost per click)。广告的总收入是CPC和noOfClicks的产品。



修改主函数：

```go
func main() {  
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
    bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
    popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
    incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
    calculateNetIncome(incomeStreams)
}
```

运行结果：

```
Income From Project 1 = $5000  
Income From Project 2 = $10000  
Income From Project 3 = $4000  
Income From Banner Ad = $1000  
Income From Popup Ad = $3750  
Net income of organisation = $23750 
```

综上，我们没有对calculateNetIncome函数做任何更改，尽管我们添加了新的收入方式。它只是因为多态性而起作用。由于新的Advertisement类型也实现了Income接口，我们可以将它添加到incomeStreams切片中。calculateNetIncome函数也在没有任何更改的情况下工作，因为它可以调用Advertisement类型的calculate()和source()方法。



千锋Go语言的学习群：784190273

对应视频地址：

https://www.bilibili.com/video/av47467197

源代码：

https://github.com/rubyhan1314/go_foundation