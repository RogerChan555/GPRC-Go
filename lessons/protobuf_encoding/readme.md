## To run the test

```
protoc --go_out=. --go_opt=paths=source_relative encoding.proto
```

## Resource for Encoding
https://protobuf.dev/programming-guides/encoding

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

## Base 128 Variants

The term "Base 128" in Base 128 Varints refers to the fact that each byte in the encoding can represent 128 different values (0-127) for the actual data, with the most significant bit (MSB) used as a continuation flag. 
This allows for efficient encoding of integers using a variable number of bytes.

### Why 128 Different Values?

A byte has 8 bits:

A byte is made up of 8 bits, which means it can represent 
2<sup>8</sup> = 256 different values, ranging from 0 to 255.

Using 7 bits for data:

In Base 128 Varint encoding, we only use 7 of these 8 bits to represent the actual data.
The remaining 1 bit (the most significant bit, or MSB) is used as a continuation flag.
7 bits can represent 128 values:

With 7 bits, the number of possible combinations is 
2<sup>7</sup> = 128
These combinations allow you to represent values from 0 to 127.

Example:
Consider a single byte in binary. Normally, 8 bits can range from:

00000000 (which is 0 in decimal)
to 11111111 (which is 255 in decimal).
But in Base 128 Varints, since only 7 bits are used for the data, the range becomes:

0000000 (which is 0 in decimal)
to 1111111 (which is 127 in decimal).
Visual Breakdown:
8-bit binary representation: b7 b6 b5 b4 b3 b2 b1 b0
7 bits for data: b6 b5 b4 b3 b2 b1 b0
1 bit for continuation flag: b7

By limiting the data representation to 7 bits, you get 128 possible values (from 0000000 to 1111111), which corresponds to the decimal range of 0 to 127.

This is why each byte in Base 128 Varint encoding can represent 128 different values. The 8th bit is reserved for signaling whether the next byte continues the value or if this byte is the last one.

### Continue to Base 128 Variants

The varint encoding of 150 is 0x96 0x01, which is 9601 in hexadecimal.

When it's marshalled to binary, it uses base 128 varints to encode the go struct to binary.

Variable-width integers, or varints, are at the core of the wire format. 

Variable-width integers (also known as variable-length integers or varints) are a way to encode integers using a variable number of bytes, rather than a fixed number. 
This encoding is commonly used in data formats where space efficiency is critical.

### More about Variable-with integers

Variable-length encoding format, often used in protocols like Protocol Buffers (Protobuf) to optimize the storage of data. Specifically, it's describing a method where small unsigned 64-bit integers 
can be encoded in fewer bytes, while larger numbers may take more bytes.

Here's how it works:
Variable-Length Encoding: The idea is to save space when encoding numbers by using fewer bytes for smaller values. Instead of always using 8 bytes to store a 64-bit integer, variable-length encoding adapts to the size of the number.

Small Values, Fewer Bytes: If the number is small, it can be represented using fewer bytes. For example, a number like 5 can be encoded in just 1 byte. On the other hand, a larger number like 1,000,000 might require more bytes, possibly 3 or 4.

Up to Ten Bytes for the Largest Values: For very large numbers, such as the maximum value of an unsigned 64-bit integer (18,446,744,073,709,551,615), the encoding may use up to 10 bytes.

This is efficient because most real-world numbers are small, so using fewer bytes reduces storage size and speeds up transmission, while still allowing large numbers to be represented when needed.

### Why unsigned? 

Range: By specifying "unsigned," it's clear that the integer can take values from 0 up to the maximum, utilizing the full 64 bits for the value itself. If the integer were signed, it would halve the range of positive values to accommodate negative numbers.

Application: Unsigned integers are often used in scenarios where negative numbers are not needed, such as counting, indexing, or representing large values where maximizing the range is important.

## Message Structure

A protocol buffer message is a series of key-value pairs. The binary version of a message just uses the field’s number as the key – the name and declared type for each field can only be determined on the decoding end by referencing the message type’s definition (i.e. the .proto file). Protoscope does not have access to this information, so it can only provide the field numbers.

When a message is encoded, each key-value pair is turned into a record consisting of the field number, a wire type and a payload. The wire type tells the parser how big the payload after it is. This allows old parsers to skip over new fields they don’t understand. This type of scheme is sometimes called Tag-Length-Value, or TLV.

There are six wire types: VARINT, I64, LEN, SGROUP, EGROUP, and I32

<img width="615" alt="Screenshot 2024-08-18 at 16 08 16" src="https://github.com/user-attachments/assets/e70ce79b-2626-40a1-a769-8d64643bdbb2">

![Screenshot 2024-08-18 at 18 00 17](https://github.com/user-attachments/assets/10624395-562a-4fe1-86b3-dc45a2e63e67)

![Screenshot 2024-08-18 at 18 02 00](https://github.com/user-attachments/assets/c8b53181-d0ed-4308-8c9a-3e3571e552ef)
