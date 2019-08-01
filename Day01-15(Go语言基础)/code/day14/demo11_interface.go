package main

import "fmt"

func main() {
	/*
	接口的嵌套：
	 */
	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("--------------")
	var a1 A = cat
	a1.test1()

	fmt.Println("--------------")
	var b1 B = cat
	b1.test2()
	fmt.Println("--------------")
	var c1 C = cat
	c1.test1()
	c1.test2()
	c1.test3()

	fmt.Println("----------")
	//var c2 C = a1
	var a2 A = c1
	a2.test1()
}

type A interface {
	test1()
}

type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}

type Cat struct {
	//如果想实现接口C，那不止要实现接口C的方法，还要实现接口A，B中的方法
}

func (c Cat) test1() {
	fmt.Println("test1()....")
}

func (c Cat) test2() {
	fmt.Println("test2()....")
}

func (c Cat) test3() {
	fmt.Println("test3()....")
}
