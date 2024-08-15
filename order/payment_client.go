package order

import (
	"github.com/fmo/grpc/protos/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type PaymentClient struct {
	client payment.PaymentServiceClient
}

func NewPaymentClient(address string) (*PaymentClient, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), // For testing purposes
	}

	clientConn, err := grpc.NewClient("localhost:50051", dialOpts...)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer clientConn.Close()

	client := pb.NewPaymentServiceClient(clientConn)

	// Create a payment request
	paymentReq := &pb.PaymentRequest{
		UserId:        "user123",
		Amount:        100.0,
		Currency:      "USD",
		PaymentMethod: "credit_card",
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call the MakePayment method
	res, err := client.MakePayment(ctx, paymentReq)
	if err != nil {
		log.Fatalf("Payment failed: %v", err)
	}

	// Process the response
	log.Printf("Payment success: %v, Transaction ID: %s, Message: %s", res.Success, res.TransactionId, res.Message)
}
