package main

import (
	"time"
	"ServerStreamCode/message"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//订单服务实现
type OrderServiceImpl struct {
}

//获取订单信息s
func (os *OrderServiceImpl) GetOrderInfos(request *message.OrderRequest, stream message.OrderService_GetOrderInfosServer) error {
	fmt.Println(" 服务端流 RPC 模式")

	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	for id, info := range orderMap {
		if (time.Now().Unix() >= request.TimeStamp) {
			fmt.Println("订单序列号ID：", id)
			fmt.Println("订单详情：", info)
			//通过流模式发送给客户端
			stream.Send(&info)
		}
	}
	return nil
}


func main() {
	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}