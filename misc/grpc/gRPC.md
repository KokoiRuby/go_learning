## RPC

Remote Procedure Call 远程过程调用，是一种协议/规范，对应的实现 [Dubbo](https://dubbo.apache.org/)、[Thrift](https://thrift.apache.org/)、[gRPC](https://grpc.io/)、[Hetty](https://hetty.xyz/)。

定义：客户端请求远程主机服务/进程，而无需了解底层网络技术 = **透明**。

- 既然 RPC 的客户端认为自己是在调用本地对象，那么传输层使用什么协议就无需 care = 透明。
- 同理，消息格式的构成，客户端也无需关心。
- 同理，服务器无论使用什么语言，调用都应该成功。

### Why RPC？

传统单体 → 云原生微服务，需要一种高效透明应用程序之间的通讯手段。

### RPC Mechanism

既然要求**透明**，那么需要对通信细节进行**封装**。

1. 客户端发起调用 Call。

2. 客户端 stub 封装消息体 Marshalling。

3. 调用 PRC 运行时发送消息。

4. 服务端接收解封 Unmarshalling 处理返回，调用服务器 stub 封装消息体 Marshalling。

5. 调用 PRC 运行时发送消息。

6. 服务端接收解封 Unmarshalling 处理返回。

   

![What is RPC in Operating System](https://static.javatpoint.com/operating-system/images/what-is-rpc-in-operating-system4.png)

## [gRPC](https://grpc.io/docs/what-is-grpc/introduction/)

> HTTP/2 + protobuf

gRPC 一开始由 google 开发，是一款语言中立、平台中立、开源的远程过程调用 RPC 系统，基于 HTTP/2。

流程：

1. 定义一个服务，指定其能够被远程调用的方法（包含参数和返回类型）。
2. 服务端实现该接口，并运行 gRPC 服务器。
3. 客户端拥有一个 stub 提供同服务端一样的方法。

**Features**:

- LB
- Security
- Observability
- SD
- Codec/Compression

![图1](http://www.grpc.io/img/grpc_concept_diagram_00.png)

### [ProtoBuf](https://protobuf.dev/)

Protocol Buffers：Google 开源的成熟的**结构数据序列化机制** JSON/XMLlike but simpler, smaller & faster。

ProtoBuf 是一种与语言和平台无关的**接口定义语言 IDL，定义服务暴露的方法以及请求/响应的格式**。



### Service Definition

服务定义，即声明一个方法（参数列表+返回值）能够让远程进行调用。

```go
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string greeting = 1;
}

message HelloResponse {
  string reply = 1;
}
```

### Using the API

**Protobuf file `.proto` → Codegen → Generated Source Code**

- Server：需实现声明方法 + 启动 gRPC server
- Client：本地对象 stub 实现了 Server 相同的方法，客户端直接调用。

### Life Cycle

- **Unary**: single req & single res

  1. 一旦客户端调用了 stub，服务器便会收到通知，知道该 RPC 已被调用，并接收到该调用的客户端元数据、方法名称和（如果适用的话）指定的截止时间。

  2. 服务器可以立即发送其初始元数据（这必须在发送任何响应之前进行），或者等待客户端的请求消息。哪个先，取决于具体实现。

  3. 一旦服务器收到客户端的请求消息，它会进行必要的工作来创建和填充响应。响应创建完成后（如果成功），会连同状态详情（状态码和可选的状态消息）以及可选的尾随元数据一起返回给客户端。

  4. 如果响应状态是 OK 的，那么客户端将收到响应，从而在客户端一侧完成调用。

     

- **Server Streaming**: single req & stream res

  - 服务器返回 stream 消息来响应客户端的请求。

  - 在发送完所有消息之后，服务器的状态详情（状态码和可选的状态消息）以及可选的尾随元数据会发送给客户端。

    

- **Client Streaming**: stream req & single res

  - 客户端发送 steram 消息给服务器，而不是一条消息
  - 。服务器会返回一条消息（以及其状态详情和可选的尾随元数据），通常是在收到客户端所有消息之后

  

- **Bidirectional streamin**: stream req & stream res

  - 两个流是独立的，客户端和服务器可以以任何顺序读取和写入消息, wait-all or ping-pong.

```go
// .proto
rpc SayHello(HelloRequest) returns (HelloResponse);
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```

![image-20240808112744146](./gRPC.assets/image-20240808112744146.png)

**Deadlines/Timeouts**

客户端可指定超时实践 + `DEADLINE_EXCEEDED` error。

服务端可以查询是否超时以及剩余时间（以完成 RPC）。

**RPC termination**

客户端和服务端独立判断调用的成功与否，他们的判断结果可能不一致。

**Cancelling an RPC**

客户端和服务端可以在任何时刻取消一个 RPC，取消会直接中止调用。

**Channels**

通道提供了与指定主机和端口上的 gRPC 服务器的连接；创建客户端 stub 时使用。客户端可以指定通道参数，以修改 gRPC 的默认行为。通道有状态：`connected` or `idle`。

### [Go](https://grpc.io/docs/languages/go/)

