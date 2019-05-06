# Server streaming RPC --- 1 request／many responses

## compile .proto file
```
protoc -I ./proto --go_out=plugins=grpc:./proto/ ./proto/file.proto
```
## run gRPC server
```
$ go run ./server
```
## execute gRPC client
```
$ go run ./client $FILE_NAME
```

