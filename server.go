package main

import (
	"Mandatory1gRPC/courseMicroService/Mandatory1gRPC/courseMicroService"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Course Service Started")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}

	s := courseMicroService.Server{AllCourses: allCourses}

	grpcServer := grpc.NewServer()

	courseMicroService.RegisterCourseServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8080: %v", err)
	}

}

var allCourses = []*courseMicroService.Course{
	{ID: "1", Name: "Danish", Ects: 1, CourseResponsibleId: 11, NRatings: 100, AvgRatings: 1.1, ActiveStudents: "0, 1, 2, 3, 4"},
	{ID: "2", Name: "Math", Ects: 15, CourseResponsibleId: 12, NRatings: 10, AvgRatings: 9.9, ActiveStudents: "1, 3, 4"},
	{ID: "3", Name: "History", Ects: 5, CourseResponsibleId: 13, NRatings: 30, AvgRatings: 5.0, ActiveStudents: "3, 4"},
	{ID: "4", Name: "SoftwareEngenering", Ects: 60, CourseResponsibleId: 14, NRatings: 50, AvgRatings: 6.0, ActiveStudents: "0, 1"},
}
