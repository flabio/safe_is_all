package helpers

import (
	"reflect"

	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
)

func MapToStructSchool(schoolDto *dto.SchoolDTO, dataMap map[string]interface{}) {
	school := dto.SchoolDTO{
		Name:    dataMap[utils.NAME].(string),
		Address: dataMap[utils.ADDRESS].(string),
		Phone:   dataMap[utils.PHONE].(string),
		Email:   dataMap[utils.EMAIL].(string),
		Active:  dataMap[utils.ACTIVE].(bool),
	}
	*schoolDto = school
}
func MapToStruct(m map[string]interface{}, result interface{}) error {
	v := reflect.ValueOf(result).Elem()

	for key, value := range m {
		field := v.FieldByName(key)
		if !field.IsValid() {
			continue
		}
		fieldType := field.Type()
		val := reflect.ValueOf(value)

		if val.Type().ConvertibleTo(fieldType) {
			field.Set(val.Convert(fieldType))
		}
	}
	return nil
}
