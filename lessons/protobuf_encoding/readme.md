## to run the test

```
protoc --go_out=. --go_opt=paths=source_relative encoding.proto
```

## What is the "Wire Format"?
The wire format is the binary encoding format that Protobuf uses to serialize the data defined in your .proto file before it's 
sent across the network (the "wire") or stored on disk.

<img width="773" alt="Screenshot 2024-08-18 at 13 48 44" src="https://github.com/user-attachments/assets/eec261a0-f80b-4a4a-af5f-9120940efca5">

<img width="821" alt="Screenshot 2024-08-18 at 14 31 21" src="https://github.com/user-attachments/assets/56883424-a0d0-4b62-96b1-04fe45c8081e">
<img width="831" alt="Screenshot 2024-08-18 at 14 31 13" src="https://github.com/user-attachments/assets/a3681aa9-20c2-4348-aa0f-362dcb167211">

The term "Base 128" in Base 128 Varints refers to the fact that each byte in the encoding can represent 128 different values (0-127) for the actual data, with the most significant bit (MSB) used as a continuation flag. This allows for efficient encoding of integers using a variable number of bytes.
