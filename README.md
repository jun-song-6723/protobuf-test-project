# protobuf-test-project

Local Project to test my knowledge on Protobuf, gRPC, and Plugins

## gRPC
Allows a simple server-client interaction where client can read or write content of the file in designated path (relative to server) by setting the mode flag; read by default ('r' for read and 'w' for write)

### Usage
To update the content of test_write.txt and read immediately:

```
go run main.go -file ./files/test_write.txt -mode w -content 'hello from client' \\
go run main.go -file ./files/test_write.txt
```
## Plugin
protoc-gen-message - writes a list of messages and services from proto in text files

### Usage
```
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       --plugin=protoc-gen-msg=protoc-plugin/protoc-gen-msg \
       --msg_out=. --msg_opt=paths=source_relative \
       file.proto
```
