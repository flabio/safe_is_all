package helpers

import (
	utils "github.com/flabio/safe_constants"
	"github.com/safe_msvc_course/usecase/dto"
)

func MapToStructCourse(courseDto *dto.CourseDTO, dataMap map[string]interface{}) {
	course := dto.CourseDTO{
		Id:       0,
		Name:     dataMap[utils.NAME].(string),
		SchoolId: uint(dataMap[utils.SCHOOL_ID].(float64)),
		Active:   dataMap[utils.ACTIVE].(bool),
	}
	*courseDto = course
}

func MapToStructTopic(courseDto *dto.TopicDTO, dataMap map[string]interface{}) {
	course := dto.TopicDTO{
		Id:        0,
		Title:     dataMap[utils.TITLE].(string),
		TimeHours: dataMap[utils.TIME_HOURS].(string),
		CourseId:  uint(dataMap[utils.COURSE_ID].(float64)),
		Active:    dataMap[utils.ACTIVE].(bool),
	}
	*courseDto = course
}
