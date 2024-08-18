## to run the test

```
protoc --go_out=. --go_opt=paths=source_relative encoding.proto
```

## What is the "Wire Format"?
The wire format is the binary encoding format that Protobuf uses to serialize the data defined in your .proto file before it's 
sent across the network (the "wire") or stored on disk.
