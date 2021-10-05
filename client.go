package main

import (
	"Mandatory1gRPC/courseMicroService/Mandatory1gRPC/courseMicroService"
	"fmt"
	"log"
	"strconv"

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

	c := courseMicroService.NewCourseServiceClient(conn)

	course := courseMicroService.Course{ID: "5", Name: "Fysik", Ects: 8, CourseResponsibleId: 55, NRatings: 250, AvgRatings: 3.1, ActiveStudents: "1, 2, 3, 4, 5"}
	course2 := courseMicroService.Course{ID: "5", Name: "Fysik", Ects: 10, CourseResponsibleId: 55, NRatings: 255, AvgRatings: 3.7, ActiveStudents: "1, 2, 3, 4"}

	getCsReq := courseMicroService.GetCoursesRequest{Message: "Give me courses"}
	postCReq := courseMicroService.PostCourseRequest{Course: &course}
	getCbyIDReq := courseMicroService.GetCourseByIDRequest{ID: "5"}
	putCbyIDReq := courseMicroService.PutCourseByIDRequest{Course: &course2}
	deleteCbyIDReq := courseMicroService.DeleteCourseByIDRequest{ID: "5"}

	getCourses(c, getCsReq)
	postCourse(c, postCReq)
	getCourseByID(c, getCbyIDReq)
	putCourseByID(c, putCbyIDReq)
	getCourseByID(c, getCbyIDReq)
	deleteCourseByID(c, deleteCbyIDReq)
	getCourses(c, getCsReq)
}

func getCourses(c courseMicroService.CourseServiceClient, req courseMicroService.GetCoursesRequest) {
	response, err := c.GetCourses(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	log.Printf("Response from Server:")
	for i := 0; i < len(response.Course); i++ {
		log.Printf(response.Course[i].ID + " " + response.Course[i].Name + " " + strconv.Itoa(int(response.Course[i].Ects)) + " " + strconv.Itoa(int(response.Course[i].NRatings)) + " " + fmt.Sprintf("%f", response.Course[i].AvgRatings) + " " + response.Course[i].ActiveStudents)
	}
}

func postCourse(c courseMicroService.CourseServiceClient, req courseMicroService.PostCourseRequest) {
	response, err := c.PostCourse(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	log.Printf("Response from Server: %s", response.Message)
}

func getCourseByID(c courseMicroService.CourseServiceClient, req courseMicroService.GetCourseByIDRequest) {
	response, err := c.GetCourseByID(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	log.Printf("Response from Server:")
	log.Printf(response.Course.ID + " " + response.Course.Name + " " + strconv.Itoa(int(response.Course.Ects)) + " " + strconv.Itoa(int(response.Course.NRatings)) + " " + fmt.Sprintf("%f", response.Course.AvgRatings) + " " + response.Course.ActiveStudents)
}

func putCourseByID(c courseMicroService.CourseServiceClient, req courseMicroService.PutCourseByIDRequest) {
	response, err := c.PutCourseByID(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	log.Printf("Response from Server: %s", response.Message)
}

func deleteCourseByID(c courseMicroService.CourseServiceClient, req courseMicroService.DeleteCourseByIDRequest) {
	response, err := c.DeleteCourseByID(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error when calling GetCourse: %s", err)
	}

	log.Printf("Response from Server: %s", response.Message)
}
