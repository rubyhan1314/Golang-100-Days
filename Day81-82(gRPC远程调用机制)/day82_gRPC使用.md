# gRPC编程使用
**@author：Davie**
**版权所有：北京千锋互联科技有限公司**

## 一、gRPC调用
在上节课内容中，我们学习了使用gRPC框架实现服务的调用编程。在gRPC框架中，诸如上节课我们学习的在客户端与服务端之间通过消息结构体定义的方式来传递数据，我们称之为“单项RPC”，也称之为简单模式。除此之外，gRPC中还有数据流模式的RPC调用实现，这正是我们本节课要学习的内容。

### 1.1、服务端流RPC
在服务端流模式的RPC实现中，服务端得到客户端请求后，处理结束返回一个数据应答流。在发送完所有的客户端请求的应答数据后，服务端的状态详情和可选的跟踪元数据发送给客户端。服务端流RPC实现案例如下：
#### 1.1.1、服务接口定义
在.proto文件中定义服务接口,使用服务端流模式定义服务接口,如下所示：
```proto
...
//订单服务service定义
service OrderService {
    rpc GetOrderInfos (OrderRequest) returns (stream OrderInfo) {}; //服务端流模式
}
```
我们可以看到与之前简单模式下的数据作为服务接口的参数和返回值不同的是，此处服务接口的返回值使用了stream进行修饰。通过stream修饰的方式表示该接口调用时，服务端会以数据流的形式将数据返回给客户端。

#### 1.1.2 编译.proto文件，生成pb.go文件
使用gRPC插件编译命令编译.proto文件，编译命令如下:
```go
protoc --go_out=plugins=grpc:. message.proto
```

#### 1.1.3 自动生成文件的变化
与数据结构体发送携带数据实现不同的时，流模式下的数据发送和接收使用新的功能方法完成。在自动生成的go代码程序当中，每一个流模式对应的服务接口，都会自动生成对应的单独的client和server程序，以及对应的结构体实现。具体编程如下图所示：
##### 1.1.3.1 服务端自动生成
```go
type OrderService_GetOrderInfosServer interface {
	Send(*OrderInfo) error
	grpc.ServerStream
}

type orderServiceGetOrderInfosServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrderInfosServer) Send(m *OrderInfo) error {
	return x.ServerStream.SendMsg(m)
}
```
流模式下，服务接口的服务端提供Send方法，将数据以流的形式进行发送
##### 1.1.3.2 客户端自动生成
```go
type OrderService_GetOrderInfosClient interface {
	Recv() (*OrderInfo, error)
	grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrderInfosClient) Recv() (*OrderInfo, error) {
	m := new(OrderInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```
流模式下，服务接口的客户端提供Recv()方法接收服务端发送的流数据。

#### 1.1.4 服务编码实现
定义好服务接口并编译生成代码文件后，即可根据规则对定义的服务进行编码实现。具体的服务编码实现如下所示：
```go
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
```
GetOrderInfos方法就是服务接口的具体实现，因为是流模式开发，服务端将数据以流的形式进行发送,因此，该方法的第二个参数类型为OrderService_GetOrderInfosServer，该参数类型是一个接口，其中包含Send方法，允许发送流数据。Send方法的具体实现在编译好的pb.go文件中，进一步调用grpc.SeverStream.SendMsg方法。

#### 1.1.5 服务的注册和监听的处理
服务的监听与处理与前文所学内容没有区别，依然是相同的步骤:
```go
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
```

#### 1.1.6 客户端数据接收
服务端使用Send方法将数据以流的形式进行发送，客户端可以使用Recv()方法接收流数据,因为数据流失源源不断的，因此使用for无限循环实现数据流的读取，当读取到io.EOF时，表示流数据结束。客户端数据读取实现如下：
```go
...
for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("读取到的信息：", orderInfo)
	}
...
```

#### 1.1.7 运行结果
按照先后顺序，依次运行server.go文件和client.go文件，可以得到运行结果。
##### 1.1.7.1 服务端运行结果
```go
 服务端流 RPC 模式
订单序列号ID： 201907300001
订单详情： {201907300001 衣服 已付款 {} [] 0}
订单序列号ID： 201907310001
订单详情： {201907310001 零食 已付款 {} [] 0}
订单序列号ID： 201907310002
订单详情： {201907310002 食品 未付款 {} [] 0}
```

##### 1.1.7.2 客户端运行结果
```go
客户端请求RPC调用：服务端流模式
读取到的信息： OrderId:"201907310001" OrderName:"\351\233\266\351\243\237" OrderStatus:"\345\267\262\344\273\230\346\254\276" 
读取到的信息： OrderId:"201907310002" OrderName:"\351\243\237\345\223\201" OrderStatus:"\346\234\252\344\273\230\346\254\276" 
读取到的信息： OrderId:"201907300001" OrderName:"\350\241\243\346\234\215" OrderStatus:"\345\267\262\344\273\230\346\254\276" 
读取结束
```

### 1.2、客户端流模式
上文演示的是服务端以数据流的形式返回数据的形式。对应的，也存在客户端以流的形式发送请求数据的形式。

#### 1.2.1 服务接口的定义
与服务端同理,客户端流模式的RPC服务声明格式，就是使用stream修饰服务接口的接收参数,具体如下所示:
```proto
...
//订单服务service定义
service OrderService {
    rpc AddOrderList (stream OrderRequest) returns (OrderInfo) {}; //客户端流模式
}
```

#### 1.2.2  编译.proto文件
使用编译命令编译.protow文件。客户端流模式中也会自动生成服务接口的接口。
##### 1.2.2.1 自动生成的服务流接口实现
```go
type OrderService_AddOrderListServer interface {
	SendAndClose(*OrderInfo) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderServiceAddOrderListServer struct {
	grpc.ServerStream
}

func (x *orderServiceAddOrderListServer) SendAndClose(m *OrderInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceAddOrderListServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```
SendAndClose和Recv方法是客户端流模式下的服务端对象所拥有的方法。

##### 1.2.2.2 自动生成的客户端流接口实现
```go
type OrderService_AddOrderListClient interface {
	Send(*OrderRequest) error
	CloseAndRecv() (*OrderInfo, error)
	grpc.ClientStream
}

type orderServiceAddOrderListClient struct {
	grpc.ClientStream
}

func (x *orderServiceAddOrderListClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceAddOrderListClient) CloseAndRecv() (*OrderInfo, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OrderInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```
Send和CloseAndRecv是客户端流模式下的客户端对象所拥有的方法。

#### 1.2.3 服务的实现
客户端流模式的服务接口具体实现如下：
```go
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
		fmt.Println(orderRequest)
	}
}
```

#### 1.2.4 服务的注册和监听处理
依然是采用相同的服务注册和监听处理方式对服务进行注册和监听处理。
```go
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
```

#### 1.2.5 客户端实现
客户端调用send方法流数据到服务端，具体实现如下：
```go
...
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
```

#### 1.2.6 程序运行
##### 1.2.6.1 服务端
运行案例，程序输出如下：
```go
 客户端流 RPC 模式
201907300001 衣服 已付款
201907310001 零食 已付款
201907310002 食品 未付款
 读取数据结束 
 客户端流 RPC 模式
201907300001 衣服 已付款
201907310001 零食 已付款
201907310002 食品 未付款
 读取数据结束 
```

##### 1.2.6.2 客户端
客户端运行程序输出如下：
```go
客户端请求RPC调用：客户端流模式
 读取数据结束 
 读取数据结束了 
```

### 1.3、双向流模式
上文已经讲过了服务端流模式和客户端流模式。如果将客户端和服务端两种流模式结合起来,就是第三种模式，双向流模式。即客户端发送数据的时候以流数据发送，服务端返回数据也以流的形式进行发送，因此称之为双向流模式。
#### 1.3.1 双向流服务的定义
```go
//订单服务service定义
service OrderService {
    rpc GetOrderInfos (stream OrderRequest) returns (stream OrderInfo) {}; //双向流模式
}
```

#### 1.3.2 编译.proto文件
##### 1.3.2.1 服务端接口实现
```go
type OrderService_GetOrderInfosServer interface {
	Send(*OrderInfo) error
	Recv() (*OrderRequest, error)
	grpc.ServerStream
}

type orderServiceGetOrderInfosServer struct {
	grpc.ServerStream
}

func (x *orderServiceGetOrderInfosServer) Send(m *OrderInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceGetOrderInfosServer) Recv() (*OrderRequest, error) {
	m := new(OrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```

##### 1.3.2.2 客户端接口实现
```go
type OrderService_GetOrderInfosClient interface {
	Send(*OrderRequest) error
	Recv() (*OrderInfo, error)
	grpc.ClientStream
}

type orderServiceGetOrderInfosClient struct {
	grpc.ClientStream
}

func (x *orderServiceGetOrderInfosClient) Send(m *OrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceGetOrderInfosClient) Recv() (*OrderInfo, error) {
	m := new(OrderInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
```

#### 1.3.3 服务实现
```go
//实现grpc双向流模式
func (os *OrderServiceImpl) GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error {

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(" 数据读取结束 ")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}

		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]message.OrderInfo{
			"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
			"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
			"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
		}

		result := orderMap[orderRequest.GetOrderId()]
		//发送数据
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
```

#### 1.3.4 服务端及客户端的编程实现
##### 1.3.4.1 服务端实现
```go
func main() {
	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
```

##### 1.3.4.2 客户端实现
```go
func main() {

	//1、Dail连接
	conn, err := grpc.Dial("localhost:8092", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)

	fmt.Println("客户端请求RPC调用：双向流模式")
	orderIDs := []string{"201907300001", "201907310001", "201907310002"}

	orderInfoClient, err := orderServiceClient.GetOrderInfos(context.Background())
	for _, orderID := range orderIDs {
		orderRequest := message.OrderRequest{OrderId: orderID}
		err := orderInfoClient.Send(&orderRequest)
		if err != nil {
			panic(err.Error())
		}
	}

	//关闭
	orderInfoClient.CloseSend()

	for {
		orderInfo, err := orderInfoClient.Recv()
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		if err != nil {
			return
		}
		fmt.Println("读取到的信息：", orderInfo)
	}
}
```

## 二、TLS验证和Token认证
上节课我们学习掌握了grpc-go框架的四种流模式。在实际的生产环境中，一个功能完整的服务，不仅包含基本的方法调用和数据交互的功能，还包括授权认证，数据追踪，负载均衡等方面。本节课，我们来看一下除了在gRPC调用过程中，如何实现授权认证，以及如何进行拦截处理。

### 2.1 授权认证
gRPC中默认支持两种授权方式,分别是：SSL/TLS认证方式、基于Token的认证方式。

#### 2.1.1 SSL/TLS认证方式
SSL全称是Secure Sockets Layer，又被称之为安全套接字层，是一种标准安全协议，用于在通信过程中建立客户端与服务器之间的加密链接。
TLS的全称是Transport Layer Security，TLS是SSL的升级版。在使用的过程中，往往习惯于将SSL和TLS组合在一起写作SSL/TLS。
简而言之，SSL/TLS是一种用于网络通信中加密的安全协议。
##### 2.1.1.1 SSL/TLS工作原理
使用SSL/TLS协议对通信连接进行安全加密，是通过非对称加密的方式来实现的。所谓非对称加密方式又称之为公钥加密，密钥对由公钥和私钥两种密钥组成。私钥和公钥成对存在，先生成私钥，通过私钥生成对应的公钥。公钥可以公开，私钥进行妥善保存。

在加密过程中：客户端想要向服务器发起链接，首先会先向服务端请求要加密的公钥。获取到公钥后客户端使用公钥将信息进行加密，服务端接收到加密信息，使用私钥对信息进行解密并进行其他后续处理，完成整个信道加密并实现数据传输的过程。
##### 2.1.1.2 制作证书
可以自己在本机计算机上安装openssl，并生成相应的证书。
```openssl
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

##### 2.1.1.3 编程实现服务端

```go
type MathManager struct {
}

func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
	fmt.Println(" 服务端 Add方法 ")
	result := request.Args1 + request.Args2
	fmt.Println(" 计算结果是：", result)
	response = new(message.Response)
	response.Code = 1;
	response.Message = "执行成功"
	return response, nil
}

func main() {

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		grpclog.Fatal("加载在证书文件失败", err)
	}

	//实例化grpc server, 开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds))

	message.RegisterMathServiceServer(server, new(MathManager))

	lis, err := net.Listen("tcp", ":8092")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}
```

##### 2.1.1.3 编程实现客户端
```
func main() {

	//TLS连接
	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "go-grpc-example")
	if err != nil {
		panic(err.Error())
	}
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
```

#### 2.1.2 基于Token认证方式

##### 2.1.2.1 Token认证介绍
在web应用的开发过程中，我们往往还会使用另外一种认证方式进行身份验证，那就是：Token认证。基于Token的身份验证是无状态，不需要将用户信息服务存在服务器或者session中。

##### 2.1.2.2 Token认证过程
基于Token认证的身份验证主要过程是：客户端在发送请求前，首先向服务器发起请求，服务器返回一个生成的token给客户端。客户端将token保存下来，用于后续每次请求时，携带着token参数。服务端在进行处理请求之前，会首先对token进行验证，只有token验证成功了，才会处理并返回相关的数据。

##### 2.1.2.3 gRPC的自定义Token认证
在gRPC中，允许开发者自定义自己的认证规则，通过

```go
grpc.WithPerRPCCredentials()
```

设置自定义的认证规则。WithPerRPCCredentials方法接收一个PerRPCCredentials类型的参数，进一步查看可以知道PerRPCCredentials是一个接口，定义如下：

```go
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
    RequireTransportSecurity() bool
}
```

因此，开发者可以实现以上接口，来定义自己的token信息。

在本案例中，我们自定义的token认证结构体如下：
```go
//token认证
type TokenAuthentication struct {
	AppKey    string
	AppSecret string
}

//组织token信息
func (ta *TokenAuthentication) RequestMetaData(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":    ta.AppKey,
		"appkey": ta.AppSecret,
	}, nil
}

//是否基于TLS认证进行安全传输
func (a *TokenAuthentication) RequireTransportSecurity() bool {
	return true
}
```

需要注意的是，RequestMetaData方法中的appid和appkey字段均需要保持小写，不能大写。RequireTransportSecurity方法用于设置是否基于tls认证进行安全传输。

在客户端进行连接时，我们将自定义的token认证信息作为参数进行传入。
```go
auth := TokenAuthentication{
		AppKey:    "hello",
		AppSecret: "20190812",
}
conn, err := grpc.Dial("localhost:8093", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
if err != nil {
	panic(err.Error())
}
```

##### 2.1.2.4 服务端
在服务端的调用方法中实现对token请求参数的判断，可以通过metadata获取token认证信息。具体实现如下：
```go
func (mm *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {

	//通过metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var appKey string
	var appSecret string

	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}
	fmt.Println(" 服务端 Add方法 ")
	result := request.Args1 + request.Args2
	fmt.Println(" 计算结果是：", result)
	response = new(message.Response)
	response.Code = 1;
	response.Message = "执行成功"
	return response, nil
}
```

最后，运行项目，token认证成功。客户端修改token信息，再次运行，会发现提示token不合法。

## 三、拦截器的使用
### 3.1、需求
在上节课程中，我们学习使用了gRPC框架中的两种认证方式：TLS验证和Token验证。

但是，在服务端的方法中，每个方法都要进行token的判断。程序效率太低，可以优化一下处理逻辑，在调用服务端的具体方法之前，先进行拦截，并进行token验证判断，这种方式称之为拦截器处理。

除了此处的token验证判断处理以外，还可以进行日志处理等。

### 3.2、Interceptor
使用拦截器，首先需要注册。
在grpc中编程实现中，可以在NewSever时添加拦截器设置，grpc框架中可以通过UnaryInterceptor方法设置自定义的拦截器，并返回ServerOption。具体代码如下：
```go
grpc.UnaryInterceptor()
```
UnaryInterceptor()接收一个UnaryServerInterceptor类型，继续查看源码定义，可以发现UnaryServerInterceptor是一个func，定义如下：
```go
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)
```

通过以上代码，如果开发者需要注册自定义拦截器，需要自定义实现UnaryServerInterceptor的定义。

### 3.3、自定义UnaryServerInterceptor
接下来就自定义实现func,符合UnaryServerInterceptor的标准，在该func的定义中实现对token的验证逻辑。自定义实现func如下：
```go
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	//通过metadata
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var appKey string
	var appSecret string
	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}
	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}
	//通过token验证，继续处理请求
	return handler(ctx, req)
}
```

在自定义的TokenInterceptor方法定义中,和之前在服务的方法调用的验证逻辑一致，从metadata中取出请求头中携带的token认证信息，并进行验证是否正确。如果token验证通过，则继续处理请求后续逻辑，后续继续处理可以由grpc.UnaryHandler进行处理。grpc.UnaryHandler同样是一个方法，其具体的实现就是开发者自定义实现的服务方法。grpc.UnaryHandler接口定义源码定义如下：
```go
type UnaryHandler func(ctx context.Context, req interface{}) (interface{}, error)
```

### 3.4、拦截器注册
在服务端调用grpc.NewServer时进行拦截器的注册。详细如下：
```go
server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(TokenInterceptor))
```

### 3.5、项目运行
依次运行server.go程序和client.go程序，可以得到程序运行的正确结果。修改token携带值，可以验证token非法情况的拦截器效果。