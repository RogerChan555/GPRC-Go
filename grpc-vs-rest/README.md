## How gRPC solves lack of verbs problems of REST

gRPC avoids the challenge of multiple update operations being difficult in REST because it is based on Remote Procedure Call (RPC), which allows defining precise methods for every operation. 
Unlike REST, which relies on a few HTTP verbs (GET, POST, PUT, DELETE, etc.), gRPC lets you define as many operations as needed using Protocol Buffers (.proto files).

Here's why gRPC solves this problem:

1. Explicit Methods for Each Update Operation
   * Instead of trying to fit updates into HTTP verbs like PUT or PATCH, you can define separate RPC methods.
     
     ```
       service AccountService {
          rpc CreditAccount (CreditRequest) returns (AccountResponse);
          rpc DebitAccount (DebitRequest) returns (AccountResponse);
          rpc TransferFunds (TransferRequest) returns (TransferResponse);
        }
     ```

   * Each method handles a specific update operation, making the API design more intuitive.

2. Streaming Support for Batch Updates
    * gRPC supports bidirectional streaming, allowing multiple updates to be sent efficiently in a single connection.
    * Example: A client can send multiple update requests in a stream instead of making multiple REST API calls.

3. Strongly Typed Messages
    * REST often uses JSON, which lacks strict schema enforcement. In gRPC, Protocol Buffers enforce structured data, reducing ambiguity.

4. Efficient Binary Communication
    * Instead of sending updates as multiple HTTP requests with JSON payloads, gRPC sends them in a compact binary format, reducing overhead.

5. Built-in Concurrency Handling
    * gRPC servers can efficiently handle multiple concurrent requests using HTTP/2 multiplexing, whereas REST APIs typically require separate HTTP requests for each update.
