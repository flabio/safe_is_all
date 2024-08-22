package helpers

import (
	"github.com/all_is_safe/infrastructure/utils"
	"github.com/all_is_safe/usecases/dto"
)

func ValidateRequired(contact dto.EmergencyContactoDTO) string {
	var msg string = utils.EMPTY
	if contact.FirstName == utils.EMPTY {
		msg = utils.FIRST_NAME_IS_REQUIRED
	}
	if contact.LastName == utils.EMPTY {
		msg = utils.LAST_NAME_IS_REQUIRED
	}
	if contact.Phone == utils.EMPTY {
		msg = utils.PHONE_IS_REQUIRED
	}
	if contact.UserId == 0 {
		msg = utils.USER_ID_IS_REQUIRED
	}
	if contact.ParentescoId == 0 {
		msg = utils.PARENTESCO_ID_IS_REQUIRED
	}

	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.FIRST_NAME] == nil {
		msg = utils.FIRST_NAME_FIELD_IS_REQUIRED
	}
	if value[utils.LAST_NAME] == nil {
		msg = utils.LAST_NAME_FIELD_IS_REQUIRED
	}

	if value[utils.PHONE] == nil {
		msg = utils.PHONE_FIELD_IS_REQUIRED
	}

	if value[utils.USER_ID] == nil {
		msg = utils.USER_ID_FIELD_IS_REQUIRED
	}

	if value[utils.PARENTESCO_ID] == nil {
		msg = utils.PARENTESCO_ID_FIELD_IS_REQUIRED
	}
	if value[utils.ACTIVE] == nil {
		msg = utils.ACTIVE_FIELD_IS_REQUIRED
	}

	return msg
}
