package academic

import "time"

//TutorInterface interface
type TutorInterface interface {
	AddContent(content *Subject) (added bool, c *Subject)
}

//Tutor struct
type Tutor struct {
	ID         int64     `json:"tutor_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"tutor_email"`
	JoinedDate time.Time `json:"joined_date"`
	Courses    []Course  `json:"tutor_courses"`
}

//NewTutor function
func NewTutor(id int64, fname, lname, email string) (tutor *Tutor) {
	tutor = newTutor(id, fname, lname, email).(*Tutor)
	return
}

func newTutor(id int64, fname, lname, email string) TutorInterface {
	return &Tutor{
		ID:         id,
		FirstName:  fname,
		LastName:   lname,
		Email:      email,
		JoinedDate: time.Now().UTC(),
	}
}

//AddContent method
func (t *Tutor) AddContent(content *Subject) (added bool, c *Subject) {
	content.TutorID = t.ID
	if t.ID == content.TutorID {
		added = true
		c = content
	}
	return
}
