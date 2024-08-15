package payments

import (
	"github.com/fmo/grpc/protos/golang/payments"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	payments.RegisterPaymentServiceServer()
}
