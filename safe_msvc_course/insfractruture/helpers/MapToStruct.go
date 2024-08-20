package helpers

import (
	"reflect"

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
