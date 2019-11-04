package main

import (
	"google.golang.org/grpc"
	"context"
	"fmt"
	"io"
	"ClientStreamCode/message"
)

func main() {

	//1、Dail连接
	conn, err := grpc.Dial("localhost:8091", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	fmt.Println("客户端请求RPC调用：客户端流模式")
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}

	//调用服务方法
	addOrderListClient, err := orderServiceClient.AddOrderList(context.Background())
	if err != nil {
		panic(err.Error())
	}
	//调用方法发送流数据
	for _, info := range orderMap {
		err = addOrderListClient.Send(&info)
		if err != nil {
			panic(err.Error())
		}
	}

	for {
		orderInfo, err := addOrderListClient.CloseAndRecv()
		if err == io.EOF {
			fmt.Println(" 读取数据结束了 ")
			return
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(orderInfo.GetOrderStatus())
	}
}
