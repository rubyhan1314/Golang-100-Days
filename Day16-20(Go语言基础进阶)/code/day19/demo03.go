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

func (p Person) Say(msg string){
	fmt.Println("hello,",msg)
}

func (p Person) PrintInfo(){
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n",p.Name,p.Age,p.Sex)
}
func main() {
	p1 := Person{"王二狗",30,"男"}
	GetMessage(p1)
}

//获取input的信息
func GetMessage(input interface{}){
	getType := reflect.TypeOf(input)//先获取input的类型
	fmt.Println("get Type is :",getType.Name()) //Person
	fmt.Println("get Kind is :" ,getType.Kind()) //struct

	getValue :=reflect.ValueOf(input)
	fmt.Println("get all Fields is :" ,getValue) //{王二狗 30 男}

	//获取字段
	/*
	step1：先获取Type对象：reflect.Type，
		NumField()
		Field(index)
	step2：通过Filed()获取每一个Filed字段
	step3：Interface()，得到对应的Value
	 */
	 for i:=0;i<getType.NumField();i++{
	 	filed := getType.Field(i)
	 	value := getValue.Field(i).Interface() //获取第一个数值
	 	fmt.Printf("字段名称：%s，字段类型：%s，字段数值：%v\n",filed.Name,filed.Type,value)
	 }


	 //获取方法
	 for i:=0;i<getType.NumMethod();i++{
	 	method := getType.Method(i)
	 	fmt.Printf("方法名称：%s，方法类型：%v\n",method.Name,method.Type)
	 }
}
