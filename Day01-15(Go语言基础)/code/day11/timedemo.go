package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	/*
	time包：
		1年=365天，day
		1天=24小时，hour
		1小时=60分钟，minute
		1分钟=60秒，second
		1秒钟=1000毫秒，millisecond
		1毫秒=1000微秒，microsecond-->μs
		1微秒=1000纳秒，nanosecond-->ns
		1纳秒=1000皮秒，picosecond-->ps
	 */

	 //1.获取当前的时间
	 t1 := time.Now()
	 fmt.Printf("%T\n",t1)
	 fmt.Println(t1) //2019-06-26 10:46:40.349196 +0800 CST m=+0.000495846

	 //2.获取指定的时间
	 t2 := time.Date(2008,7,15,16,30,28,0,time.Local)
	 fmt.Println(t2) //2008-07-15 16:30:28 +0800 CST

	 //3.time-->string之间的转换
	 /*
	 t1.Format("格式模板")-->string
	 	模板的日期必须是固定：06-1-2-3-4-5
	  */
	 s1 := t1.Format("2006-1-2 15:04:05")
	 fmt.Println(s1)

	 s2 := t1.Format("2006/01/02")
	 fmt.Println(s2)

	 //string-->time
	 /*
	 time.Parse("模板",str)-->time,err
	  */
	 s3 := "1999年10月10日"//string
	 t3 ,err := time.Parse("2006年01月02日",s3)
	 if err != nil{
	 	fmt.Println("err:",err)
	 }
	 fmt.Println(t3)
	 fmt.Printf("%T\n",t3)

	 fmt.Println(t1.String())

	 //4.根据当前时间，获取指定的内容
	year,month,day :=  t1.Date() //年，月，日
	fmt.Println(year,month,day)

	hour,min,sec := t1.Clock()
	fmt.Println(hour,min,sec) //时，分，秒

	year2 := t1.Year()
	fmt.Println("年：",year2)
	fmt.Println(t1.YearDay())

	month2 := t1.Month()
	fmt.Println("月：",month2)
	fmt.Println("日：",t1.Day())
	fmt.Println("时：",t1.Hour())
	fmt.Println("分钟：",t1.Minute())
	fmt.Println("秒：",t1.Second())
	fmt.Println("纳秒：",t1.Nanosecond())

	fmt.Println(t1.Weekday()) //Wednesday

	//5.时间戳：指定的日期，距离1970年1月1日0点0时0分0秒的时间差值：秒，纳秒

	t4 := time.Date(1970,1,1,1,0,0,0,time.UTC)
	timeStamp1:=t4.Unix() //秒的差值
	fmt.Println(timeStamp1)
	timeStamp2 := t1.Unix()
	fmt.Println(timeStamp2)

	timeStamp3 := t4.UnixNano()
	fmt.Println(timeStamp3) //3600 000 000 000
	timeStamp4 := t1.UnixNano()
	fmt.Println(timeStamp4)

	//6.时间间隔
	t5 := t1.Add(time.Minute)
	fmt.Println(t1)
	fmt.Println(t5)
	fmt.Println(t1.Add(24 * time.Hour))

	t6 := t1.AddDate(1,0,0)
	fmt.Println(t6)

	d1 := t5.Sub(t1)
	fmt.Println(d1)

	//7.睡眠
	time.Sleep(3 *time.Second) //让当前的程序进入睡眠状态
	fmt.Println("main。。。over。。。。。")

	//睡眠[1-10]的随机秒数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1 //int
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum)*time.Second)
	fmt.Println("睡醒了。。")
}
