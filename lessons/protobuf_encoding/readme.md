## to run the test

```
protoc --go_out=. --go_opt=paths=source_relative encoding.proto
```

## What is the "Wire Format"?
The wire format is the binary encoding format that Protobuf uses to serialize the data defined in your .proto file before it's 
sent across the network (the "wire") or stored on disk.

<img width="773" alt="Screenshot 2024-08-18 at 13 48 44" src="https://github.com/user-attachments/assets/eec261a0-f80b-4a4a-af5f-9120940efca5">
