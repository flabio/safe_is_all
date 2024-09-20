package dto

type TopicDTO struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	TimeHours string `json:"time_hours"`
	CourseId  uint   `json:"courseid"`
	Active    bool   `json:"active"`
}
