package academic

import "time"

//StudentInterface interface
type StudentInterface interface {
	AddCourses(course *Course) (added bool)
}

//Student struct
type Student struct {
	ID         int64     `json:"student_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"student_email"`
	JoinedDate time.Time `json:"joined_date"`
	Courses    []*Course `json:"student_courses"`
}

//NewStudent function
func NewStudent(id int64, fname, lname, email string) (student *Student) {
	student = newStudent(id, fname, lname, email).(*Student)
	return
}

func newStudent(id int64, fname, lname, email string) StudentInterface {
	return &Student{
		ID:         id,
		FirstName:  fname,
		LastName:   lname,
		Email:      email,
		JoinedDate: time.Now().UTC(),
	}
}

//AddCourses method
func (s *Student) AddCourses(course *Course) (added bool) {
	initialCount := len(s.Courses)
	s.Courses = append(s.Courses, course)
	if initialCount < len(s.Courses) {
		added = true
	}
	return
}
