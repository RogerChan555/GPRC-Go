## To run the test

```
protoc --go_out=. --go_opt=paths=source_relative encoding.proto
```

## Resource for Encoding
https://protobuf.dev/programming-guides/encoding/#varints

## What is the "Wire Format"?
The wire format is the binary encoding format that Protobuf uses to serialize the data defined in your .proto file before it's 
sent across the network (the "wire") or stored on disk.

## Protoscope to Describe the low-level Wire Format

Protoscope is a very simple language for describing snippets of the low-level wire format, which we’ll use to provide a visual reference for the encoding of various messages. Protoscope’s syntax consists of a sequence of tokens that each encode down to a specific byte sequence.

## A Simple Example

Let's use a simple example to illustrate how much space a serialized Protobuf message might
consume on disk.

```
message Test {
  optional int32 a = 1;
}
```

if you set `a = 150`, the wire format would look like this:

```
08 96 01
```

Here's the breakdown:

```
08 (1 byte): This encodes the field number (`1`) and the wire type (`variant`)
96 01 (2 bytes): This is the variable-length encoding of the integer `150`
```

So, this message consumes 3 bytes in total.

<img width="831" alt="Screenshot 2024-08-18 at 14 31 13" src="https://github.com/user-attachments/assets/a3681aa9-20c2-4348-aa0f-362dcb167211">
<img width="821" alt="Screenshot 2024-08-18 at 14 31 21" src="https://github.com/user-attachments/assets/56883424-a0d0-4b62-96b1-04fe45c8081e">

The term "Base 128" in Base 128 Varints refers to the fact that each byte in the encoding can represent 128 different values (0-127) for the actual data, with the most significant bit (MSB) used as a continuation flag. This allows for efficient encoding of integers using a variable number of bytes.

The varint encoding of 150 is 0x96 0x01, which is 9601 in hexadecimal.

When its marshalled to binary, it uses base 128 varints to encode the go struct to binary.

## Message Structure

A protocol buffer message is a series of key-value pairs. The binary version of a message just uses the field’s number as the key – the name and declared type for each field can only be determined on the decoding end by referencing the message type’s definition (i.e. the .proto file). Protoscope does not have access to this information, so it can only provide the field numbers.

When a message is encoded, each key-value pair is turned into a record consisting of the field number, a wire type and a payload. The wire type tells the parser how big the payload after it is. This allows old parsers to skip over new fields they don’t understand. This type of scheme is sometimes called Tag-Length-Value, or TLV.

There are six wire types: VARINT, I64, LEN, SGROUP, EGROUP, and I32

<img width="615" alt="Screenshot 2024-08-18 at 16 08 16" src="https://github.com/user-attachments/assets/e70ce79b-2626-40a1-a769-8d64643bdbb2">

![Screenshot 2024-08-18 at 18 00 17](https://github.com/user-attachments/assets/10624395-562a-4fe1-86b3-dc45a2e63e67)

![Screenshot 2024-08-18 at 18 02 00](https://github.com/user-attachments/assets/c8b53181-d0ed-4308-8c9a-3e3571e552ef)
