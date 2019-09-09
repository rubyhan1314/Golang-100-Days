package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
	School string
}
func main() {
	s1 := Student{"王二狗",19,"千锋教育"}
	//通过反射，更改对象的数值：前提也是数据可以被更改
	fmt.Printf("%T\n",s1) //main.Student
	p1 := &s1
	fmt.Printf("%T\n",p1)  //*main.Student
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name,p1.Name)

	//改变数值
	value:=reflect.ValueOf(&s1)
	if value.Kind() == reflect.Ptr{
		newValue:=value.Elem()
		fmt.Println(newValue.CanSet())  //true

		f1 :=newValue.FieldByName("Name")
		f1.SetString("韩茹")
		f3:=newValue.FieldByName("School")
		f3.SetString("幼儿园")
		fmt.Println(s1)
	}
}
