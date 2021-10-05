package main

import (
	"Mandatory1gRPC/course/Mandatory1gRPC/course"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not conect: %s", err)
	}

	defer conn.Close()

	c := course.NewCourseServiceClient(conn)

	getcourse := course.GetCourse{Body: "Hello from client!"}

	response, err := c.GiveCourse(context.Background(), &getcourse)
	if err != nil {
		log.Fatalf("Error when calling GiveCourse: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
