package main

import "fmt"

func main() {
	/*
	接口：interface
		在Go中，接口是一组方法签名。

		当某个类型为这个接口中的所有方法提供了方法的实现，它被称为实现接口。

		Go语言中，接口和类型的实现关系，是非侵入式

		//其他语言中，要显示的定义
		class Mouse implements USB{}

	1.当需要接口类型的对象时，可以使用任意实现类对象代替
	2.接口对象不能访问实现类中的属性


	多态：一个事物的多种形态
		go语言通过接口模拟多态

		就一个接口的实现
			1.看成实现本身的类型，能够访问实现类中的属性和方法
			2.看成是对应的接口类型，那就只能够访问接口中的方法

	接口的用法：
		1.一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数。
		2.定义一个类型为接口类型，实际上可以赋值为任意实现类的对象


	鸭子类型：

	 */
	 //1.创建Mouse类型
	 m1 := Mouse{"罗技小红"}
	 fmt.Println(m1.name)
	 //2.创建FlashDisk
	 f1 := FlashDisk{"闪迪64G"}
	 fmt.Println(f1.name)

	 testInterface(m1)
	 testInterface(f1)

	 var usb USB
	 usb= f1
	 usb.start()
	 usb.end()
	 //fmt.Println(usb.name)

	 f1.deleteData()
	 //usb.de

	 var arr [3]USB
	 arr[0] = m1
	 arr[1] = f1
	 fmt.Println(arr)
}

//1.定义接口
type USB interface {
	start() //USB设备开始工作
	end() //USB设备结束工作
}

//2.实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse)start(){
	fmt.Println(m.name,"鼠标，准备就绪，可以开始工作了，点点点。。")
}
func (m Mouse) end(){
	fmt.Println(m.name,"结束工作，可以安全退出。。")
}

func (f FlashDisk)start(){
	fmt.Println(f.name,"准备开始工作，可以进行数据的存储。。")
}
func (f FlashDisk)end(){
	fmt.Println(f.name,"可以弹出。。")
}

//3.测试方法
func testInterface(usb USB){ //usb = m1  usb = f1
	usb.start()
	usb.end()
}

func (f FlashDisk) deleteData(){
	fmt.Println(f.name,"U盘删除数据。。")
}


