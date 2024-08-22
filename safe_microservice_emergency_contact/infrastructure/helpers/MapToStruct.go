package helpers

import (
	"github.com/all_is_safe/infrastructure/utils"
	"github.com/all_is_safe/usecases/dto"
)

func MapToStructContact(contactDto *dto.EmergencyContactoDTO, dataMap map[string]interface{}) {
	contact := dto.EmergencyContactoDTO{
		Id:           0,
		FirstName:    dataMap[utils.FIRST_NAME].(string),
		LastName:     dataMap[utils.LAST_NAME].(string),
		Phone:        dataMap[utils.PHONE].(string),
		UserId:       uint(dataMap[utils.USER_ID].(float64)),
		ParentescoId: uint(dataMap[utils.PARENTESCO_ID].(float64)),
		Active:       dataMap[utils.ACTIVE].(bool),
	}
	*contactDto = contact
}

// func MapToStruct(m map[string]interface{}, result interface{}) error {
// 	v := reflect.ValueOf(result).Elem()

// 	for key, value := range m {
// 		field := v.FieldByName(key)
// 		if !field.IsValid() {
// 			continue
// 		}
// 		fieldType := field.Type()
// 		val := reflect.ValueOf(value)

// 		if val.Type().ConvertibleTo(fieldType) {
// 			field.Set(val.Convert(fieldType))
// 		}
// 	}
// 	return nil
// }
