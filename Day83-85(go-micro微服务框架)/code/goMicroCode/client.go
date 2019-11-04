package main

import (
	"goMicroCode/message"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"time"
)

func main() {

	service := micro.NewService(
		micro.Name("student.client"),
	)

	service.Init()

	studentService := message.NewStudentServiceClient("student_service", service.Client())

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "davie"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
	time.Sleep(50 * time.Second)
}
