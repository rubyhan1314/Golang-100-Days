package main

import (
	"google.golang.org/grpc"
	"grpcSSLCode/message"
	"context"
	"google.golang.org/grpc/grpclog"
	"fmt"
	"google.golang.org/grpc/credentials"
)

func main() {

	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "go-grpc-example")
	if err != nil {
		panic(err.Error())
	}

	grpc.WithInsecure()

	//1、Dail连接
	conn, err := grpc.Dial("localhost:8092", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	serviceClient := message.NewMathServiceClient(conn)

	addArgs := message.RequestArgs{Args1: 3, Args2: 5}

	response, err := serviceClient.AddMethod(context.Background(), &addArgs)
	if err != nil {
		grpclog.Fatal(err.Error())
	}

	fmt.Println(response.GetCode(), response.GetMessage())
}
