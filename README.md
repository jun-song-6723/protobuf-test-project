# protobuf-test-project

Local Project to test my knowledge on Protobuf, gRPC, and Plugins

Client can read or write content of the file in designated path by setting the mode flag; read by default ('r' for read and 'w' for write)

### Usage
go run main.go -file ./files/test_write.txt -mode w -content 'hello from client'
go run main.go -file ./files/test_write.txt

Also supports custom plugin:
protoc-gen-message - Writes a list of messages and services in text files

### Usage
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       --plugin=protoc-gen-msg=protoc-plugin/protoc-gen-msg \
       --msg_out=. --msg_opt=paths=source_relative \
       file.proto
