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

	// Hex literal
	hexLiteral := []byte{0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x0a}
	fmt.Printf("Hex literal: %x\n", hexLiteral)

	// UTF-8 string
	utf8String := "Hello, Protobuf!"
	fmt.Printf("UTF-8 string: %x\n", utf8String)
}
