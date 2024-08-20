package helpers

import (
	"log"

	"github.com/safe_msvc_course/insfractruture/utils"
	"github.com/safe_msvc_course/usecase/dto"
)

func ValidateRequired(course dto.CourseDTO) string {
	var msg string = utils.EMPTY
	if course.Name == utils.EMPTY {
		msg = utils.NAME_IS_REQUIRED
	}
	log.Println(course.SchoolId)
	if course.SchoolId == 0 {
		msg = utils.SCHOOL_ID_IS_REQUIRED
	}
	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_IS_FIELD_REQUIRED
	}
	if value[utils.SCHOOL_ID] == nil {
		msg = utils.SCHOOL_ID_IS_FIELD_REQUIRED
	}

	return msg
}
