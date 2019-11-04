package main

import (
	"google.golang.org/grpc"
	"ServerStreamCode/message"
	"io"
	"fmt"
	"context"
	"time"
)

func main() {
	//1、Dail连接
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	request := message.OrderRequest{TimeStamp: time.Now().Unix()}

	orderInfosClient, err := orderServiceClient.GetOrderInfos(context.TODO(), &request)

	for {
		orderInfo, err := orderInfosClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("读取到的信息：", orderInfo)
	}

}
