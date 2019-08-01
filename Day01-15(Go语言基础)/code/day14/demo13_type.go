package main

import "fmt"

type Person struct {
	name string
}

func (p Person) show(){
	fmt.Println("Person--->",p.name)
}

//类型别名
type People = Person

func (p People) show2(){
	fmt.Println("People--->",p.name)
}

type Student struct {
	//嵌入两个结构体
	Person
	People
}

func main() {
	var s Student
	//s.name = "王二狗" //ambiguous selector s.name
	s.Person.name = "王二狗"
	//s.show() //ambiguous selector s.show
	s.Person.show()

	fmt.Printf("%T,%T\n",s.Person,s.People) //main.Person,main.Person

	s.People.name = "李小花"
	s.People.show2()

	s.Person.show()
}