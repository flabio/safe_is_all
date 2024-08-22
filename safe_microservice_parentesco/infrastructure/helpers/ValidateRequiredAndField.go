package helpers

import (
	"github.com/all_is_safe/infrastructure/utils"
	"github.com/all_is_safe/usecases/dto"
)

func ValidateRequired(course dto.ParentescoDTO) string {
	var msg string = utils.EMPTY
	if course.Name == utils.EMPTY {
		msg = utils.NAME_IS_REQUIRED
	}

	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_FIEDL_IS_REQUIRED
	}
	if value[utils.ACTIVE] == nil {
		msg = utils.ACTIVE_FIELD_IS_REQUIRED
	}

	return msg
}
