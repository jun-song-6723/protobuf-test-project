package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "google.golang.org"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect to")
	path    = flag.String("file", "./files/test.txt", "file to get content")
	mode    = flag.String("mode", "r", "read/write mode")
	content = flag.String("content", "", "for write mode; content of the file")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFileServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *mode {
	case "r":
		file_to_request := pb.FileReadRequest{Path: *path}
		r, err := c.RequestFileRead(ctx, &file_to_request)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
		log.Printf("\nFile Size: %d\nFile Content: %s\n", r.GetSize(), r.GetBuffer())
	case "w":
		file_to_edit := pb.FileWriteRequest{Path: *path, Buffer: *content}
		r, err := c.RequestFileWrite(ctx, &file_to_edit)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}

		if r.Ok {
			log.Printf("File updated!")
		} else {
			log.Printf("File write failed")
		}
	default:
		log.Printf("invalid mode; enter 'r' for read and 'w' for write")
	}
}
