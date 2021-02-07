package academic

import "time"

//Course struct
type Course struct {
	ID       int64    `json:"course_id"`
	Name     string   `json:"course_name"`
	SchoolID int64    `json:"school_id"`
	Content  *Subject `json:"course_content"`
}

//Subject struct
type Subject struct {
	ID        int64                  `json:"subject_id"`
	TutorID   int64                  `json:"tutor_id"`
	CreatedAt time.Time              `json:"created_at"`
	Heading   map[string]interface{} `json:"subject_heading"`
	Body      map[string]interface{} `json:"subject_body"`
}
