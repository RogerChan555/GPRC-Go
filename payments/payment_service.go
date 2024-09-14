package main

import (
	"context"
	"fmt"
	"github.com/fmo/grpc/protos/golang/discounts"
	"github.com/fmo/grpc/protos/golang/payments"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"time"
)

type PaymentServiceServer struct {
	payments.UnimplementedPaymentServiceServer
}

func (s *PaymentServiceServer) MakePayment(ctx context.Context, req *payments.PaymentRequest) (*payments.PaymentResponse, error) {
	log.Printf("Processing payment for user: %s, amount: %f", req.UserId, req.Amount)

	time.Sleep(3 * time.Second)

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient("localhost:50053", opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to discount service: %v", err)
	}
	defer conn.Close()

	discountClient := discounts.NewDiscountServiceClient(conn)

	discountReq := &discounts.CheckDiscountRequest{
		CouponCode: "some-code",
	}

	discountRes, err := discountClient.CheckDiscount(ctx, discountReq)
	if err != nil {
		return nil, fmt.Errorf("failed to check discount: %v", err)
	}

	if discountRes.Success {
		// apply discount
	}

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
