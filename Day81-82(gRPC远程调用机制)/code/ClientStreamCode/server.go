package main

import (
	"google.golang.org/grpc"
	"net"
	"fmt"
	"ClientStreamCode/message"
	"io"
)

//订单服务实现
type OrderServiceImpl struct {
}

//添加订单信息服务实现
func (os *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
	fmt.Println(" 客户端流 RPC 模式")

	for {
		//从流中读取数据信息
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(" 读取数据结束 ")
			result := message.OrderInfo{OrderStatus: " 读取数据结束 "}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//打印接收到的数据
		fmt.Println(orderRequest.GetOrderId(),orderRequest.GetOrderName(),orderRequest.GetOrderStatus())
	}
}

func main() {

	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8091")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
