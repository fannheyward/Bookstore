gRPC demo with grpc-gateway.

```
protoc -I $GOPATH/src/github.com/googleapis/googleapis -I . pb/book.proto --go_out=plugins=grpc:.
protoc -I $GOPATH/src/github.com/googleapis/googleapis -I. pb/book.proto --grpc-gateway_out=logtostderr=true:.
protoc -I $GOPATH/src/github.com/googleapis/googleapis/ -I. pb/book.proto --swagger_out=logtostderr=true:.
```
