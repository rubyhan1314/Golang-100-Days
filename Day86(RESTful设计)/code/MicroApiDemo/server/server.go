package main

import (
	"MicroApiDemo/proto"
	"context"
	"github.com/micro/go-micro"
	"log"
	"errors"
	"fmt"
)

type StudentServiceImpl struct {
}

//服务实现
func (ss *StudentServiceImpl) GetStudent(ctx context.Context, request *proto.Request, resp *proto.Student) error {

	//tom
	studentMap := map[string]proto.Student{
		"davie":  proto.Student{Name: "davie", Classes: "软件工程专业", Grade: 80},
		"steven": proto.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   proto.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   proto.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	if request.Name == "" {
		return errors.New(" 请求参数错误,请重新请求。")
	}

	//获取对应的student
	student := studentMap[request.Name]
	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*resp = student
		return nil
	}
	return errors.New(" 未查询当相关学生信息 ")
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.student"),
	)

	service.Init()

	//注册
	proto.RegisterStudentServiceHandler(service.Server(), new(StudentServiceImpl))

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
