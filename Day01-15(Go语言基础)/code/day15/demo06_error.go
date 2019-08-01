package main

import "fmt"

func main() {
	length,width := 6.7,-9.1
	area, err := rectArea(length,width)
	if err != nil{
		fmt.Println(err)
		if err,ok := err.(*areaError);ok{
			if err.legnthNegative(){
				fmt.Printf("error：长度，%.2f，小于零\n",err.lenght)
			}
			if err.widthNegative(){
				fmt.Printf("error：宽度，%.2f，小于零\n",err.width)
			}
		}
		return
	}
	fmt.Println("矩形的面积是：",area)
}

type areaError struct {
	msg string //错误的描述
	lenght float64 //发生错误的时候，矩形的长度
	width float64 //发生错误的时候，矩形的宽度
}

func (e *areaError) Error() string{
	return e.msg
}

func (e *areaError) legnthNegative() bool{
	return e.lenght <0
}

func (e *areaError) widthNegative()bool{
	return e.width <0
}

func rectArea(length ,width float64)(float64,error){
	msg := ""
	if length <0 {
		msg = "长度小于零"
	}
	if width <0 {
		if msg == ""{
			msg = "宽度小于零"
		}else{
			msg += "，宽度也小于零"
		}
	}

	if msg != ""{
		return 0,&areaError{msg,length,width}
	}
	return length * width ,nil
}
