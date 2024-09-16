package main

import (
	"context"
	"fmt"
	"github.com/fmo/grpc/protos/golang/discounts"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type DiscountServiceServer struct {
	discounts.UnimplementedDiscountServiceServer
}

func (s *DiscountServiceServer) CheckDiscount(ctx context.Context, req *discounts.CheckDiscountRequest) (*discounts.CheckDiscountResponse, error) {
	log.Printf("Checking discount for the coupon code %s", req.CouponCode)

	time.Sleep(1 * time.Second)

	return &discounts.CheckDiscountResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	discounts.RegisterDiscountServiceServer(grpcServer, &DiscountServiceServer{})

	fmt.Println("Starting gRPC server on :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
