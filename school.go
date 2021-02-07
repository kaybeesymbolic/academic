package academic

import (
	"strconv"
	"time"
)

//SchoolInterface interface
type SchoolInterface interface {
	AddStudent(student *Student) (added bool)
	AddTutor(tutor *Tutor) (added bool)
	AddCourse(course *Tutor) (added bool)
}

//School struct
type School struct {
	ID       int64      `json:"school_id"`
	Students []*Student `json:school_students"`
	Tutors   []*Tutor   `json:"school_tutors"`
	Courses  []*Course  `json:"school_courses"`
}

//NewSchool function
func NewSchool(id int64) *School {
	return newSchool(id).(*School)
}
func newSchool(id int64) SchoolInterface {
	return &School{ID: id}
}

//AddTutor method
func (s *School) AddTutor(tutor *Tutor) (added bool) {
	count := len(s.Tutors)
	s.Tutors = append(s.Tutors, tutor)
	if count < len(s.Tutors) {
		added = true
	}
	return
}

//AddStudent method
func (s *School) AddStudent(student *Student) (added bool) {
	count := len(s.Students)
	s.Students = append(s.Students, student)
	if count < len(s.Students) {
		added = true
	}
	return
}

//AddCourse method
func (s *School) AddCourse(tutor *Tutor) (added bool) {
	var id int64 = 848
	count := len(s.Courses)
	subject := &Subject{
		ID:        332,
		TutorID:   tutor.ID,
		CreatedAt: time.Now().UTC(),
		Heading:   map[string]interface{}{"Title": "Title of subject"},
		Body:      map[string]interface{}{"Body": "Body of subject"},
	}
	course := &Course{
		ID:       id,
		Name:     "Course" + strconv.Itoa(int(id)),
		SchoolID: s.ID,
		Content:  subject,
	}
	s.Courses = append(s.Courses, course)
	if count < len(s.Courses) {
		added = true
	}
	return
}
