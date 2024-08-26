package helpers

import (
	utils "github.com/flabio/safe_constants"
	"github.com/safe_msvc_user/usecase/dto"
)

func MapToStructSchool(schoolDto *dto.SchoolDTO, dataMap map[string]interface{}) {
	school := dto.SchoolDTO{
		Name:           dataMap[utils.NAME].(string),
		Address:        dataMap[utils.ADDRESS].(string),
		Phone:          dataMap[utils.PHONE].(string),
		Email:          dataMap[utils.EMAIL].(string),
		ZipCode:        dataMap[utils.ZIP_CODE].(string),
		StateId:        uint(dataMap[utils.STATE_ID].(float64)),
		ProviderNumber: dataMap[utils.PROVIDER_NUMBER].(string),
		Active:         dataMap[utils.ACTIVE].(bool),
	}
	*schoolDto = school
}
