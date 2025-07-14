package main

import (
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	// Read the CodeGeneratorRequest from stdin.
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
		os.Exit(1)
	}

	// Unmarshal the request.
	var req pluginpb.CodeGeneratorRequest
	if err := proto.Unmarshal(input, &req); err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal request: %v\n", err)
		os.Exit(1)
	}

	// Initialize the protogen plugin.
	plugin, err := protogen.Options{}.New(&req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	// Process each file in the request.
	for _, f := range plugin.Files {
		// Skip files that are not meant to be generated.
		if !f.Generate {
			continue
		}

		// Create an output file, e.g., <proto_file>.messages.txt
		outMsgFile := f.GeneratedFilenamePrefix + ".messages.txt"
		outSerFile := f.GeneratedFilenamePrefix + ".services.txt"
		var msgContent string
		var serContent string

		// Iterate over messages in the proto file.
		for _, msg := range f.Messages {
			msgContent += fmt.Sprintf("Message: %s\n", msg.Desc.Name())
		}

		for _, ser := range f.Services {
			serContent += fmt.Sprintf("Service: %s\n", ser.Desc.Name())
		}

		// fmt.Printf("Go import path %s \n", f.GoImportPath.String())

		// Register the output file in the response.
		plugin.NewGeneratedFile(outMsgFile, f.GoImportPath).Write([]byte(msgContent))
		plugin.NewGeneratedFile(outSerFile, f.GoImportPath).Write([]byte(serContent))
	}

	// Marshal and write the CodeGeneratorResponse to stdout.
	resp := plugin.Response()
	output, err := proto.Marshal(resp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to marshal response: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stdout.Write(output); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write output: %v\n", err)
		os.Exit(1)
	}
}
