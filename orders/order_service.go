package main

import (
	"context"
	"fmt"
	"github.com/fmo/grpc/protos/golang/orders"
	"github.com/fmo/grpc/protos/golang/payments"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type OrderServiceServer struct {
	orders.UnimplementedOrderServiceServer
}

func (s *OrderServiceServer) PlaceOrder(ctx context.Context, req *orders.OrderRequest) (*orders.OrderResponse, error) {
	log.Printf("Creating order for user: %s, items: %s", req.UserId, req.Items)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to payment service: %v", err)
	}
	defer conn.Close()

	paymentClient := payments.NewPaymentServiceClient(conn)

	var totalAmount float64
	for _, item := range req.Items {
		totalAmount += float64(item.Quantity) * 10.0
	}

	paymentReq := &payments.PaymentRequest{
		UserId: req.UserId,
		Amount: totalAmount,
	}

	paymentCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	paymentRes, err := paymentClient.MakePayment(paymentCtx, paymentReq)
	if err != nil {
		return nil, fmt.Errorf("failed to make payment: %v", err)
	}

	if !paymentRes.Success {
		return &orders.OrderResponse{
			Success: false,
			OrderId: "",
			Message: "Payment failed, order not created",
		}, nil
	}

	return &orders.OrderResponse{
		Success: true,
		OrderId: "order_12345",
		Message: "Order created successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	orders.RegisterOrderServiceServer(grpcServer, &OrderServiceServer{})

	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
