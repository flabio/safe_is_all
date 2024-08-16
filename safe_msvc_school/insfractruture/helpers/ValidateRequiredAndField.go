package helpers

import (
	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
)

func ValidateRequired(user dto.SchoolDTO) string {
	var msg string = utils.EMPTY
	if user.Name == utils.EMPTY {
		msg = utils.NAME_IS_REQUIRED
	}
	if user.Address == utils.EMPTY {
		msg = utils.ADDRESS_IS_REQUIRED
	}
	if user.Phone == utils.EMPTY {
		msg = utils.PHONE_IS_REQUIRED
	}
	if user.Email == utils.EMPTY {
		msg = utils.EMAIL_IS_REQUIRED
	}
	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_IS_FIELD_REQUIRED
	}
	if value[utils.ADDRESS] == nil {
		msg = utils.ADDRESS_IS_FIELD_REQUIRED
	}
	if value[utils.PHONE] == nil {
		msg = utils.PHONE_IS_FIELD_REQUIRED
	}
	return msg
}
