package courseMicroService

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"golang.org/x/net/context"
)

type Server struct {
	AllCourses []*Course
}

func (s *Server) GetCourses(ctx context.Context, req *GetCoursesRequest) (*GetCoursesResponce, error) {
	log.Printf("From client: %s", req.Message)
	return &GetCoursesResponce{Course: s.AllCourses}, nil
}

func (s *Server) PostCourse(ctx context.Context, req *PostCourseRequest) (*PostCourseResponce, error) {
	log.Printf("From client: A Course that have been added")
	log.Printf(req.Course.ID + " " + req.Course.Name + " " + strconv.Itoa(int(req.Course.Ects)) + " " + strconv.Itoa(int(req.Course.CourseResponsibleId)) + strconv.Itoa(int(req.Course.NRatings)) + " " + fmt.Sprintf("%f", req.Course.AvgRatings) + " " + req.Course.ActiveStudents)
	s.AllCourses = append(s.AllCourses, req.Course)
	return &PostCourseResponce{Message: "Course was succesfully added."}, nil
}

func (s *Server) GetCourseByID(ctx context.Context, req *GetCourseByIDRequest) (*GetCourseByIDResponce, error) {
	log.Printf("From client: ID to find %s", req.ID)
	for _, a := range s.AllCourses {
		if a.ID == req.ID {
			return &GetCourseByIDResponce{Course: a}, nil
		}
	}
	return nil, errors.New("couldn't find ID")
}

func (s *Server) DeleteCourseByID(ctx context.Context, req *DeleteCourseByIDRequest) (*DeleteCourseByIDResponce, error) {
	log.Printf("From client: ID to delete %s", req.ID)
	lenBefore := len(s.AllCourses)
	var newAllCourses []*Course
	for _, a := range s.AllCourses {
		if a.ID != req.ID {
			newAllCourses = append(newAllCourses, a)
		}
	}
	s.AllCourses = newAllCourses
	if lenBefore == len(s.AllCourses) {
		return nil, errors.New("couldn't find ID")
	} else {
		return &DeleteCourseByIDResponce{Message: "Course deleted"}, nil
	}

}

func (s *Server) PutCourseByID(ctx context.Context, req *PutCourseByIDRequest) (*PutCourseByIDResponse, error) {
	log.Printf("From client: ID to put %s", req.Course.ID)
	log.Printf(req.Course.ID + " " + req.Course.Name + " " + strconv.Itoa(int(req.Course.Ects)) + " " + strconv.Itoa(int(req.Course.CourseResponsibleId)) + strconv.Itoa(int(req.Course.NRatings)) + " " + fmt.Sprintf("%f", req.Course.AvgRatings) + " " + req.Course.ActiveStudents)
	var newAllCourses []*Course
	var edited bool
	for _, a := range s.AllCourses {
		if a.ID != req.Course.ID {
			newAllCourses = append(newAllCourses, a)
		} else {
			newAllCourses = append(newAllCourses, req.Course)
			edited = true
		}
	}
	s.AllCourses = newAllCourses
	if edited {
		return &PutCourseByIDResponse{Message: "Course updated"}, nil
	} else {
		return nil, errors.New("couldn't find ID")
	}
}
