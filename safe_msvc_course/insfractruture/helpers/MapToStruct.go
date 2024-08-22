package helpers

import (
	"github.com/safe_msvc_course/insfractruture/utils"
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
