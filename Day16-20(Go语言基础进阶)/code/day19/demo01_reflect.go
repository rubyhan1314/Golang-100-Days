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
