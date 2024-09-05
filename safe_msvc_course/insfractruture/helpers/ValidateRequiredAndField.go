package helpers

import (
	utils "github.com/flabio/safe_constants"
	"github.com/safe_msvc_course/usecase/dto"
)

func ValidateRequired(course dto.CourseDTO) string {
	var msg string = utils.EMPTY
	if course.Name == utils.EMPTY {
		msg = utils.NAME_IS_REQUIRED
	}
	if course.SchoolId == 0 {
		msg = utils.SCHOOL_ID_IS_REQUIRED
	}
	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_FIELD_IS_REQUIRED
	}
	if value[utils.SCHOOL_ID] == nil {
		msg = utils.SCHOOL_ID_IS_FIELD_REQUIRED
	}

	return msg
}
