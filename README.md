# Intro

gRPC is a modern, lightweight communication protocol and high-performance RPC framework introduced by Google. It can efficiently connect services in a microservices environment
with built-in support for load balancing, tracing, health checking, and authentication. gRPC provides easy to use and efficient communication using protocol buffers,
an open source mechanism for serializing structured data.

# Benefits of gRPC Microservices

### Performance

gRPC offers significant performance benefits through efficient binary serialization using Protocol Buffers, reducing message size and transmission time compared to text-based formats like JSON.

### Code generation and interoperability

gRPC enhances code generation and interoperability by automatically generating client and server code from service definitions written in Protocol Buffers, ensuring consistent and compatible communication across different programming languages.

### Fault tolerance

gRPC provides fault tolerance through features like automatic retries, deadline propagation, and load balancing, ensuring resilient and reliable service-to-service communication.

### Security

gRPC ensures robust security by supporting TLS for encrypted communication, providing authentication mechanisms, and allowing fine-grained control over access policies and permissions.

## Protocol Buffer Compiler Installation

https://grpc.io/docs/protoc-installation/

```
brew install protobuf
protoc --version 
```

## Golang Specific Installations

```
go install google.golang.org/protobuf/cmd/protoc-gen-geo@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Streaming

RPC supports efficient streaming by enabling client-side, server-side, and bidirectional streaming, allowing continuous data flow and real-time communication between services.

# 10 Questions about gRPC

1) What’s the difference between RPC and gRPC?
2) What are protocol buffers? 
3) Difference between json and protocol buffers?
4) What are the features of the gRPC framework?
5) How to use gRPC in a Microservice environment?

# To create stubs run it under protos folder

```
protoc \
    --go_out=./golang/orders \
    --go_opt=paths=source_relative \
    --go-grpc_out=./golang/orders \
    --go-grpc_opt=paths=source_relative \
    ./order.proto
```

```
protoc \
    --go_out=./golang/payments \
    --go_opt=paths=source_relative \
    --go-grpc_out=./golang/payments \
    --go-grpc_opt=paths=source_relative \
    ./payment.proto
```

This is mainly called as Unary RPC.

## To call the order service

```
 grpcurl -d '{                                           
  "userId": "user_123",
  "items": [
    {"productId": "product_1", "quantity": 2},
    {"productId": "product_2", "quantity": 1}
  ]
}' -plaintext localhost:50052 orders.OrderService/PlaceOrder
```

## To tag sub golang module

```
git tag protos/golang/orders/v0.0.2
git push origin protos/golang/orders/v0.0.2
```

And under protos/golang/orders go mod should be created. Also the repo requires that 
will do the following 

```
go get github.com/fmo/grpc/protos/golang/orders@v0.0.2
```

## Streaming

Putting stream keyword in front of request or response objects or both (called as bidirectional streaming)
makes it possible to stream the data.

## Protocol Buffers

Protocol buffers allow you to serialize structured data to be transmitted over a wire. You can also define service functions and generate language-specific source code. The definitions of messages and service functions are written in a configuration file called a .proto file.
