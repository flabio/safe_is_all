package service

import (
	"encoding/json"
	"log"

	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/helpers"
	"github.com/safe_msvc_user/usecase/dto"
)

func ValidateSchool(id uint, s *SchoolService, c *fiber.Ctx) (dto.SchoolDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var schoolDto dto.SchoolDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msgValid := helpers.ValidateField(dataMap)
	if msgValid != utils.EMPTY {
		return dto.SchoolDTO{}, msgValid
	}

	helpers.MapToStructSchool(&schoolDto, dataMap)
	msgReq := helpers.ValidateRequired(schoolDto)
	if msgReq != utils.EMPTY {
		return dto.SchoolDTO{}, msgReq
	}
	existEmail, _ := s.UiSchool.GetSchoolFindByEmail(id, schoolDto.Email)
	if existEmail.Email != utils.EMPTY {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	existProviderNumber, _ := s.UiSchool.GetSchoolFindByProviderNumber(id, schoolDto.ProviderNumber)
	if existProviderNumber.ProviderNumber != utils.EMPTY {
		msg = "The Provider number ready exists "
	}
	log.Println("schoolDto")
	log.Println(schoolDto)
	return schoolDto, msg
}
