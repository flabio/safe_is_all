package helpers

import (
	"github.com/all_is_safe/infrastructure/utils"
	"github.com/all_is_safe/usecases/dto"
)

func MapToStruct(parentDto *dto.ParentescoDTO, dataMap map[string]interface{}) {
	parent := dto.ParentescoDTO{
		Id:     0,
		Name:   dataMap[utils.NAME].(string),
		Active: dataMap[utils.ACTIVE].(bool),
	}
	*parentDto = parent
}
