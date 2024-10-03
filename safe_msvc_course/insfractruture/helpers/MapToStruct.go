package helpers

import (
	"github.com/safe_msvc_course/insfractruture/utils"
	"github.com/safe_msvc_course/usecase/dto"
)

func MapToStructCourse(courseDto *dto.CourseDTO, dataMap map[string]interface{}) {
	course := dto.CourseDTO{
		Id:     0,
		Name:   dataMap[utils.NAME].(string),
		Active: dataMap[utils.ACTIVE].(bool),
	}
	*courseDto = course
}

func MapToStructTopic(topicDto *dto.TopicDTO, dataMap map[string]interface{}) {
	topic := dto.TopicDTO{
		Id:        0,
		Title:     dataMap[utils.TITLE].(string),
		TimeHours: dataMap[utils.TIME_HOURS].(string),
		CourseId:  uint(dataMap[utils.COURSE_ID].(float64)),
		Active:    dataMap[utils.ACTIVE].(bool),
	}
	*topicDto = topic
}
