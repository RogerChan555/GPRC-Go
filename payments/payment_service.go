package main

import (
	"context"
	"fmt"
	"github.com/fmo/grpc/protos/golang/payments"
	"google.golang.org/grpc"
	"log"
	"net"
)

type PaymentServiceServer struct {
	payments.UnimplementedPaymentServiceServer
}

func (s *PaymentServiceServer) MakePayment(ctx context.Context, req *payments.PaymentRequest) (*payments.PaymentResponse, error) {
	log.Printf("Processing payment for user: %s, amount: %f", req.UserId, req.Amount)

	return &payments.PaymentResponse{
		Success:       true,
		TransactionId: "txn_12345",
		Message:       "Payment processed successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	payments.RegisterPaymentServiceServer(grpcServer, &PaymentServiceServer{})

	fmt.Println("Starting gRPC server on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
