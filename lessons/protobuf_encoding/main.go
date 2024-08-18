package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	pb "protobuf_encoding/protobufs"
)

func main() {
	// Create a Test1 message and set 'a' to 150
	msg := &pb.Test1{
		A: 150,
	}

	// Serialize the message to binary
	serializedMsg, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("Failed to serialize message: %v", err)
	}

	// Print the serialized output as hexadecimal
	fmt.Printf("Serialized message: %X\n", serializedMsg)
}
