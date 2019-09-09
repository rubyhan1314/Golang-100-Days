#反射reflect

## 一、引入

先看官方Doc中Rob Pike给出的关于反射的定义：

```
Reflection in computing is the ability of a program to examine its own structure, particularly through types; it’s a form of metaprogramming. It’s also a great source of confusion.
(在计算机领域，反射是一种让程序——主要是通过类型——理解其自身结构的一种能力。它是元编程的组成之一，同时它也是一大引人困惑的难题。)

```

维基百科中的定义：

```
在计算机科学中，反射是指计算机程序在运行时（Run time）可以访问、检测和修改它本身状态或行为的一种能力。用比喻来说，反射就是程序在运行的时候能够“观察”并且修改自己的行为。
```



不同语言的反射模型不尽相同，有些语言还不支持反射。《Go 语言圣经》中是这样定义反射的：

> Go 语言提供了一种机制在运行时更新变量和检查它们的值、调用它们的方法，但是在编译时并不知道这些变量的具体类型，这称为反射机制。



为什么要用反射

需要反射的 2 个常见场景：

1. 有时你需要编写一个函数，但是并不知道传给你的参数类型是什么，可能是没约定好；也可能是传入的类型很多，这些类型并不能统一表示。这时反射就会用的上了。
2. 有时候需要根据某些条件决定调用哪个函数，比如根据用户的输入来决定。这时就需要对函数和函数的参数进行反射，在运行期间动态地执行函数。

但是对于反射，还是有几点不太建议使用反射的理由：

1. 与反射相关的代码，经常是难以阅读的。在软件工程中，代码可读性也是一个非常重要的指标。
2. Go 语言作为一门静态语言，编码过程中，编译器能提前发现一些类型错误，但是对于反射代码是无能为力的。所以包含反射相关的代码，很可能会运行很久，才会出错，这时候经常是直接 panic，可能会造成严重的后果。
3. 反射对性能影响还是比较大的，比正常代码运行速度慢一到两个数量级。所以，对于一个项目中处于运行效率关键位置的代码，尽量避免使用反射特性。



## 二、相关基础

反射是如何实现的？我们以前学习过 interface，它是 Go 语言实现抽象的一个非常强大的工具。当向接口变量赋予一个实体类型的时候，接口会存储实体的类型信息，反射就是通过接口的类型信息实现的，反射建立在类型的基础上。

Go 语言在 reflect 包里定义了各种类型，实现了反射的各种函数，通过它们可以在运行时检测类型的信息、改变类型的值。在进行更加详细的了解之前，我们需要重新温习一下Go语言相关的一些特性，所谓温故知新，从这些特性中了解其反射机制是如何使用的。

| 特点                   | 说明                                                         |
| ---------------------- | ------------------------------------------------------------ |
| go语言是静态类型语言。 | 编译时类型已经确定，比如对已基本数据类型的再定义后的类型，反射时候需要确认返回的是何种类型。 |
|空接口interface{}|go的反射机制是要通过接口来进行的，而类似于Java的Object的空接口可以和任何类型进行交互，因此对基本数据类型等的反射也直接利用了这一特点|
Go语言的类型：

- 变量包括（type, value）两部分

  ​	理解这一点就知道为什么nil != nil了

- type 包括 static type和concrete type. 简单来说 static type是你在编码是看见的类型(如int、string)，concrete type是runtime系统看见的类型

- 类型断言能否成功，取决于变量的concrete type，而不是static type。因此，一个 reader变量如果它的concrete type也实现了write方法的话，它也可以被类型断言为writer。



Go是静态类型语言。每个变量都拥有一个静态类型，这意味着每个变量的类型在编译时都是确定的：int，float32, *AutoType, []byte,  chan []int 诸如此类。

在反射的概念中， 编译时就知道变量类型的是静态类型；运行时才知道一个变量类型的叫做动态类型。

- 静态类型
  静态类型就是变量声明时的赋予的类型。比如：

 ```go
type MyInt int // int 就是静态类型

type A struct{
    Name string  // string就是静态
}
var i *int  // *int就是静态类型

 ```



- 动态类型
  动态类型：运行时给这个变量赋值时，这个值的类型(如果值为nil的时候没有动态类型)。一个变量的动态类型在运行时可能改变，这主要依赖于它的赋值（前提是这个变量是接口类型）。

```go
var A interface{} // 静态类型interface{}
A = 10            // 静态类型为interface{}  动态为int
A = "String"      // 静态类型为interface{}  动态为string
var M *int
A = M             // A的值可以改变

```



Go语言的反射就是建立在类型之上的，Golang的指定类型的变量的类型是静态的（也就是指定int、string这些的变量，它的type是static type），在创建变量的时候就已经确定，反射主要与Golang的interface类型相关（它的type是concrete type），只有interface类型才有反射一说。



在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型:

```go
(value, type)
```

value是实际变量值，type是实际变量的类型。一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。

例如，创建类型为*os.File的变量，然后将其赋给一个接口变量r：

```go
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

var r io.Reader
r = tty
```

接口变量r的pair中将记录如下信息：(tty, *os.File)，这个pair在接口变量的连续赋值过程中是不变的，将接口变量r赋给另一个接口变量w:

```go
var w io.Writer
w = r.(io.Writer)
```

接口变量w的pair与r的pair相同，都是:(tty, *os.File)，即使w是空接口类型，pair也是不变的。

interface及其pair的存在，是Golang中实现反射的前提，理解了pair，就更容易理解反射。反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。

所以我们要理解两个基本概念 Type 和 Value，它们也是 Go语言包中 reflect 空间里最重要的两个类型。

## 三、Type和Value

我们一般用到的包是reflect包。

既然反射就是用来检测存储在接口变量内部(值value；类型concrete type) pair对的一种机制。那么在Golang的reflect反射包中有什么样的方式可以让我们直接获取到变量内部的信息呢？ 它提供了两种类型（或者说两个方法）让我们可以很容易的访问接口变量内容，分别是reflect.ValueOf() 和 reflect.TypeOf()，看看官方的解释

```go
// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i.  ValueOf(nil) returns the zero 
func ValueOf(i interface{}) Value {...}

翻译一下：ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0


// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type {...}

翻译一下：TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil

```



reflect.TypeOf()是获取pair中的type，reflect.ValueOf()获取pair中的value。

首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数。

```go
t := reflect.TypeOf(i) //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
v := reflect.ValueOf(i) //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
```



示例代码：

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//反射操作：通过反射，可以获取一个接口类型变量的 类型和数值
	var x float64 =3.4

	fmt.Println("type:",reflect.TypeOf(x)) //type: float64
	fmt.Println("value:",reflect.ValueOf(x)) //value: 3.4

	fmt.Println("-------------------")
	//根据反射的值，来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64: ",v.Kind() == reflect.Float64)
	fmt.Println("type : ",v.Type())
	fmt.Println("value : ",v.Float())
}

```

运行结果：

```
type: float64
value: 3.4
-------------------
kind is float64:  true
type :  float64
value :  3.4

```

说明

1. reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
2. reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
3. 也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种

Type 和 Value 都包含了大量的方法，其中第一个有用的方法应该是 Kind，这个方法返回该类型的具体信息：Uint、Float64 等。Value 类型还包含了一系列类型方法，比如 Int()，用于返回对应的值。以下是Kind的种类：

```go

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

```



## 四、反射的规则

其实反射的操作步骤非常的简单，就是通过实体对象获取反射对象(Value、Type)，然后操作相应的方法即可。

下图描述了实例、Value、Type 三者之间的转换关系：

![WX20190827-170219](img/WX20190827-170219.png)

反射 API 的分类总结如下：

**1) 从实例到 Value**

通过实例获取 Value 对象，直接使用 reflect.ValueOf() 函数。例如：

```go
func ValueOf(i interface {}) Value
```

**2) 从实例到 Type**

通过实例获取反射对象的 Type，直接使用 reflect.TypeOf() 函数。例如：

```go
func TypeOf(i interface{}) Type
```

**3) 从 Type 到 Value**

Type 里面只有类型信息，所以直接从一个 Type 接口变量里面是无法获得实例的 Value 的，但可以通过该 Type 构建一个新实例的 Value。reflect 包提供了两种方法，示例如下：

```go
//New 返回的是一个 Value，该 Value 的 type 为 PtrTo(typ)，即 Value 的 Type 是指定 typ 的指针类型
func New(typ Type) Value
//Zero 返回的是一个 typ 类型的零佳，注意返回的 Value 不能寻址，位不可改变
func Zero(typ Type) Value
```

如果知道一个类型值的底层存放地址，则还有一个函数是可以依据 type 和该地址值恢复出 Value 的。例如：

```go
func NewAt(typ Type, p unsafe.Pointer) Value
```

**4) 从 Value 到 Type**

从反射对象 Value 到 Type 可以直接调用 Value 的方法，因为 Value 内部存放着到 Type 类型的指针。例如：

```go
func (v Value) Type() Type
```



**5) 从 Value 到实例**

Value 本身就包含类型和值信息，reflect 提供了丰富的方法来实现从 Value 到实例的转换。例如：

```go
//该方法最通用，用来将 Value 转换为空接口，该空接口内部存放具体类型实例
//可以使用接口类型查询去还原为具体的类型
func (v Value) Interface() （i interface{})

//Value 自身也提供丰富的方法，直接将 Value 转换为简单类型实例，如果类型不匹配，则直接引起 panic
func (v Value) Bool () bool
func (v Value) Float() float64
func (v Value) Int() int64
func (v Value) Uint() uint64
```



**6) 从 Value 的指针到值**

从一个指针类型的 Value 获得值类型 Value 有两种方法，示例如下。

```go
//如果 v 类型是接口，则 Elem() 返回接口绑定的实例的 Value，如采 v 类型是指针，则返回指针值的 Value，否则引起 panic
func (v Value) Elem() Value
//如果 v 是指针，则返回指针值的 Value，否则返回 v 自身，该函数不会引起 panic
func Indirect(v Value) Value
```



**7) Type 指针和值的相互转换**

指针类型 Type 到值类型 Type。例如：

```go
//t 必须是 Array、Chan、Map、Ptr、Slice，否则会引起 panic
//Elem 返回的是其内部元素的 Type
t.Elem() Type
```



值类型 Type 到指针类型 Type。例如：

```go
//PtrTo 返回的是指向 t 的指针型 Type
func PtrTo(t Type) Type
```



**8) Value 值的可修改性**

Value 值的修改涉及如下两个方法：

```go
//通过 CanSet 判断是否能修改
func (v Value ) CanSet() bool
//通过 Set 进行修改
func (v Value ) Set(x Value)
```



Value 值在什么情况下可以修改？我们知道实例对象传递给接口的是一个完全的值拷贝，如果调用反射的方法 reflect.ValueOf() 传进去的是一个值类型变量， 则获得的 Value 实际上是原对象的一个副本，这个 Value 是无论如何也不能被修改的。

**根据 Go 官方关于反射的博客，反射有三大定律：**

> 1. Reflection goes from interface value to reflection object.
> 2. Reflection goes from reflection object to interface value.
> 3. To modify a reflection object, the value must be settable.

第一条是最基本的：反射可以从接口值得到反射对象。

​		反射是一种检测存储在 interface中的类型和值机制。这可以通过 TypeOf函数和 ValueOf函数得到。

第二条实际上和第一条是相反的机制，反射可以从反射对象获得接口值。

​		它将 ValueOf的返回值通过 Interface()函数反向转变成 interface变量。

前两条就是说 接口型变量和 反射类型对象可以相互转化，反射类型对象实际上就是指的前面说的 reflect.Type和 reflect.Value。

第三条不太好懂：如果需要操作一个反射变量，则其值必须可以修改。

​		反射变量可设置的本质是它存储了原变量本身，这样对反射变量的操作，就会反映到原变量本身；反之，如果反射变量不能代表原变量，那么操作了反射变量，不会对原变量产生任何影响，这会给使用者带来疑惑。所以第二种情况在语言层面是不被允许的。



## 五、反射的使用

### 5.1 从relfect.Value中获取接口interface的信息

当执行reflect.ValueOf(interface)之后，就得到了一个类型为”relfect.Value”变量，可以通过它本身的Interface()方法获得接口变量的真实内容，然后可以通过类型判断进行转换，转换为原有真实类型。不过，我们可能是已知原有类型，也有可能是未知原有类型，因此，下面分两种情况进行说明。

#### 已知原有类型

已知类型后转换为其对应的类型的做法如下，直接通过Interface方法然后强制转换，如下：

```go
realValue := value.Interface().(已知的类型)
```

示例代码：

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}
```

运行结果：

```
0xc000098000
1.2345
```

说明

1. 转换的时候，如果转换的类型不完全符合，则直接panic，类型要求非常严格！
2. 转换的时候，要区分是指针还是指
3. 也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”



#### 未知原有类型

很多情况下，我们可能并不知道其具体类型，那么这个时候，该如何做呢？需要我们进行遍历探测其Filed来得知，示例如下:

```go
package main

import (
	"fmt"
	"reflect"
)




type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person)Say(msg string)  {
	fmt.Println("hello，",msg)
}
func (p Person)PrintInfo()  {
	fmt.Printf("姓名：%s,年龄：%d，性别：%s\n",p.Name,p.Age,p.Sex)
}



func main() {
	p1 := Person{"王二狗",30,"男"}

	DoFiledAndMethod(p1)

}

// 通过接口来获取任意参数
func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input) //先获取input的类型
	fmt.Println("get Type is :", getType.Name()) // Person
	fmt.Println("get Kind is : ", getType.Kind()) // struct

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue) //{王二狗 30 男}

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface() //获取第i个值
		fmt.Printf("字段名称:%s, 字段类型:%s, 字段数值:%v \n", field.Name, field.Type, value)
	}

	// 通过反射，操作方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	// 2. 再公国reflect.Type的Method获取其Method
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称:%s, 方法类型:%v \n", method.Name, method.Type)
	}
}

```



运行结果：

```
get Type is : Person
get Kind is :  struct
get all Fields is: {王二狗 30 男}
字段名称:Name, 字段类型:string, 字段数值:王二狗 
字段名称:Age, 字段类型:int, 字段数值:30 
字段名称:Sex, 字段类型:string, 字段数值:男 
方法名称:PrintInfo, 方法类型:func(main.Person) 
方法名称:Say, 方法类型:func(main.Person, string) 
```



说明

通过运行结果可以得知获取未知类型的interface的具体变量及其类型的步骤为：

1. 先获取interface的reflect.Type，然后通过NumField进行遍历
2. 再通过reflect.Type的Field获取其Field
3. 最后通过Field的Interface()得到对应的value

通过运行结果可以得知获取未知类型的interface的所属方法（函数）的步骤为：

1. 先获取interface的reflect.Type，然后通过NumMethod进行遍历
2. 再分别通过reflect.Type的Method获取对应的真实的方法（函数）
3. 最后对结果取其Name和Type得知具体的方法名
4. 也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
5. struct 或者 struct 的嵌套都是一样的判断处理方式



如果是struct的话，可以使用Elem()

```go
tag := t.Elem().Field(0).Tag //获取定义在struct里面的Tag属性
name := v.Elem().Field(0).String() //获取存储在第一个字段里面的值
```



### 5.2 通过reflect.Value设置实际变量的值

reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，即：要修改反射类型的对象就一定要保证其值是“addressable”的。

这里需要一个方法：

![WX20190826-143547](img/WX20190826-143547.png)



解释起来就是：Elem返回接口v包含的值或指针v指向的值。如果v的类型不是interface或ptr，它会恐慌。如果v为零，则返回零值。



```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of pointer:", newValue.Type())
	fmt.Println("settability of pointer:", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	////////////////////
	// 如果reflect.ValueOf的参数不是指针，会如何？
	//pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}

```

运行结果：

```
old value of pointer: 1.2345
type of pointer: float64
settability of pointer: true
new value of pointer: 77
```

说明

1. 需要传入的参数是* float64这个指针，然后可以通过pointer.Elem()去获取所指向的Value，**注意一定要是指针**。
2. 如果传入的参数不是指针，而是变量，那么
   - 通过Elem获取原始值对应的对象则直接panic
   - 通过CanSet方法查询是否可以设置返回false
3. newValue.CantSet()表示是否可以重新设置其值，如果输出的是true则可修改，否则不能修改，修改完之后再进行打印发现真的已经修改了。
4. reflect.Value.Elem() 表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的
5. 也就是说如果要修改反射类型对象，其值必须是“addressable”【对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象】
6. struct 或者 struct 的嵌套都是一样的判断处理方式



### 5.3 通过reflect.Value来进行方法的调用

这算是一个高级用法了，前面我们只说到对类型、变量的几种反射的用法，包括如何获取其值、其类型、以及如何重新设置新值。但是在项目应用中，另外一个常用并且属于高级的用法，就是通过reflect来进行方法【函数】的调用。比如我们要做框架工程的时候，需要可以随意扩展方法，或者说用户可以自定义方法，那么我们通过什么手段来扩展让用户能够自定义呢？关键点在于用户的自定义方法是未可知的，因此我们可以通过reflect来搞定。

Call()方法：

![WX20190902-144001](img/WX20190902-144001.png)



**通过反射，调用方法。**

先获取结构体对象，然后

示例代码：

```go
package main

import (
	"fmt"
	"reflect"
)




type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person)Say(msg string)  {
	fmt.Println("hello，",msg)
}
func (p Person)PrintInfo()  {
	fmt.Printf("姓名：%s,年龄：%d，性别：%s\n",p.Name,p.Age,p.Sex)
}

func (p Person) Test(i,j int,s string){
	fmt.Println(i,j,s)
}


// 如何通过反射来进行方法的调用？
// 本来可以用结构体对象.方法名称()直接调用的，
// 但是如果要通过反射，
// 那么首先要将方法注册，也就是MethodByName，然后通过反射调动mv.Call

func main() {
	p2 := Person{"Ruby",30,"男"}
	// 1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，
	// 得到“反射类型对象”后才能做下一步处理
	getValue := reflect.ValueOf(p2)

	// 2.一定要指定参数为正确的方法名
	// 先看看没有参数的调用方法

	methodValue1 := getValue.MethodByName("PrintInfo")
	fmt.Printf("Kind : %s, Type : %s\n",methodValue1.Kind(),methodValue1.Type())
	methodValue1.Call(nil) //没有参数，直接写nil

	args1 := make([]reflect.Value, 0) //或者创建一个空的切片也可以
	methodValue1.Call(args1)

	// 有参数的方法调用
	methodValue2 := getValue.MethodByName("Say")
	fmt.Printf("Kind : %s, Type : %s\n",methodValue2.Kind(),methodValue2.Type())
	args2 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)

	methodValue3 := getValue.MethodByName("Test")
	fmt.Printf("Kind : %s, Type : %s\n",methodValue3.Kind(),methodValue3.Type())
	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200),reflect.ValueOf("Hello")}

	methodValue3.Call(args3)
}

```

运行结果：

```
Kind : func, Type : func()
姓名：Ruby,年龄：30，性别：男
姓名：Ruby,年龄：30，性别：男
Kind : func, Type : func(string)
hello， 反射机制
Kind : func, Type : func(int, int, string)
100 200 Hello
```



**通过反射，调用函数。**

首先我们要先确认一点，函数像普通的变量一样，之前的章节中我们在讲到函数的本质的时候，是可以把函数作为一种变量类型的，而且是引用类型。如果说Fun()是一个函数，那么f1 := Fun也是可以的，那么f1也是一个函数，如果直接调用f1()，那么运行的就是Fun()函数。

那么我们就先通过ValueOf()来获取函数的反射对象，可以判断它的Kind，是一个func，那么就可以执行Call()进行函数的调用。

示例代码：

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//函数的反射
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("Kind : %s , Type : %s\n",value.Kind(),value.Type()) //Kind : func , Type : func()

	value2 := reflect.ValueOf(fun2)
	fmt.Printf("Kind : %s , Type : %s\n",value2.Kind(),value2.Type()) //Kind : func , Type : func(int, string)


	//通过反射调用函数
	value.Call(nil)

	value2.Call([]reflect.Value{reflect.ValueOf(100),reflect.ValueOf("hello")})

}

func fun1(){
	fmt.Println("我是函数fun1()，无参的。。")
}

func fun2(i int, s string){
	fmt.Println("我是函数fun2()，有参数。。",i,s)
}

```



**说明**

1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理
2. reflect.Value.MethodByName这个MethodByName，需要指定准确真实的方法名字，如果错误将直接panic，MethodByName返回一个函数值对应的reflect.Value方法的名字。
3. []reflect.Value，这个是最终需要调用的方法的参数，可以没有或者一个或者多个，根据实际参数来定。
4. reflect.Value的 Call 这个方法，这个方法将最终调用真实的方法，参数务必保持一致，如果reflect.Value.Kind不是一个方法，那么将直接panic。
5. 本来可以用对象访问方法直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName，然后通过反射调用methodValue.Call

## 





本文参照：

http://www.sohu.com/a/313420275_657921

https://studygolang.com/articles/12348?fr=sidebar

http://c.biancheng.net/golang/









千锋Go语言的学习群：784190273

github知识库：

https://github.com/rubyhan1314

Golang网址：

https://www.qfgolang.com/

作者B站：

https://space.bilibili.com/353694001

对应视频地址：

https://www.bilibili.com/video/av56018934

https://www.bilibili.com/video/av47467197

源代码：

https://github.com/rubyhan1314/go_reflect



