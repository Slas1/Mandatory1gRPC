package course

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) GiveCourse(ctx context.Context, course *GetCourse) (*GetCourse, error) {
	log.Printf("Recevied message body from client: %s", course.Body)
	return &GetCourse{Body: "Hello From the Server!"}, nil
}
