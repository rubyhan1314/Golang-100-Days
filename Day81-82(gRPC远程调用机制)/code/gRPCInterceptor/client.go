package main

import (
	"google.golang.org/grpc"
	"gRPCInterceptor/message"
	"context"
	"google.golang.org/grpc/grpclog"
	"fmt"
	"google.golang.org/grpc/credentials"
)

//token认证
type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

//组织token认证的metadata信息
func (ta *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":    ta.AppKey,
		"appkey": ta.AppSecret,
	}, nil
}

//是否基于TLS认证进行安全传输
func (a *TokenAuthentication) RequireTransportSecurity() bool {
	return true
}

func main() {

	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "go-grpc-example")
	if err != nil {
		panic(err.Error())
	}

	auth := TokenAuthentication{
		AppKey:    "hello1",
		AppSecret: "20190812",
	}

	//1、Dail连接
	conn, err := grpc.Dial("localhost:8093", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
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
