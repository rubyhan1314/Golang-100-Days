package main

import (
	"net/rpc"
	"RpcAndProto/message"
	"time"
	"fmt"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}
	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "201907310003", TimeStamp: timeStamp}
	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*response)
}
