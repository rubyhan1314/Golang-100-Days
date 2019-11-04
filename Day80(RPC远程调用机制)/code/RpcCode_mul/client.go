package main

import (
	"net/rpc"
	"RpcCode_mul/param"
	"fmt"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":8082")
	if err != nil {
		panic(err.Error())
	}

	var result *float32
	addParma := &param.AddParma{Args1: 1.2, Args2: 2.3}
	err = client.Call("MathUtil.Add", addParma, &result)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("计算结果:", *result)
}
