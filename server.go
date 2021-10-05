package main

import (
	"Mandatory1gRPC/course/Mandatory1gRPC/course"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	s := course.Server{}

	grpcServer := grpc.NewServer()

	course.RegisterCourseServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8080: %v", err)
	}

}
