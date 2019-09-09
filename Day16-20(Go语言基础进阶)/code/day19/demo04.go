package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.23
	fmt.Println("num的数值是：",num)

	// 需要操作指针
	//通过reflect.ValueOf() 获取num的Value对象
	pointer:=reflect.ValueOf(&num) //注意参数必须是指针才能修改其值
	newValue:=pointer.Elem()

	fmt.Println("类型：",newValue.Type()) //float64
	fmt.Println("是否可以修改数据：",newValue.CanSet()) //true

	//重新赋值
	newValue.SetFloat(3.14)
	fmt.Println(num)


	//如果reflect.ValueOf的参数不是指针？
	value := reflect.ValueOf(num)
	//value.SetFloat(6.28) //panic: reflect: reflect.Value.SetFloat using unaddressable value
	//fmt.Println(value.CanSet()) //false

	//value.Elem() //如果非指针，会直接panic，panic: reflect: call of reflect.Value.Elem on float64 Value
}