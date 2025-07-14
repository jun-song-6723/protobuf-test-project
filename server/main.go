package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "google.golang.org"
	"google.golang.org/grpc"
)

type file_server struct {
	pb.UnimplementedFileServiceServer
}

func (s *file_server) RequestFileRead(_ context.Context, req *pb.FileReadRequest) (*pb.FileReadResponse, error) {
	path_to_file := req.Path
	content, err := os.ReadFile(path_to_file)

	if err != nil {
		return nil, err
	}

	file_info, err := os.Stat(path_to_file)

	if err != nil {
		return nil, err
	}

	size := file_info.Size()
	return &pb.FileReadResponse{Size: size, Buffer: string(content)}, nil
}

func (s *file_server) RequestFileWrite(_ context.Context, req *pb.FileWriteRequest) (*pb.FileWriteResponse, error) {
	path_to_file := req.Path
	content := req.Buffer

	err := os.WriteFile(path_to_file, []byte(content), 0644)

	if err != nil {
		return &pb.FileWriteResponse{Ok: false}, err
	}

	return &pb.FileWriteResponse{Ok: true}, nil
}

var (
	port = flag.Int("port", 50051, "server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &file_server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
