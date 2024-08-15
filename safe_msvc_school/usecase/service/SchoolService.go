package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/core"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
	"github.com/ulule/deepcopier"
)

type schoolService struct {
	uiSchool uicore.UISchoolCore
}

func NewSchoolService() global.UISchoolGlobal {
	return &schoolService{uiSchool: core.NewSchoolRepository()}
}

func (s *schoolService) GetSchoolFindAll(c *fiber.Ctx) error {
	results, err := s.uiSchool.GetSchoolFindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: fiber.StatusOK,
		utils.DATA:   results,
	})
}

func (s *schoolService) GetSchoolFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.uiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: fiber.StatusOK,
		utils.DATA:   result,
	})
}
func (s *schoolService) CreateSchool(c *fiber.Ctx) error {
	var schoolCreate entities.School

	dataDto, msgError := validateSchool(s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&schoolCreate)
	result, err := s.uiSchool.CreateSchool(schoolCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_CREATE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}

func (s *schoolService) UpdateSchool(c *fiber.Ctx) error {
	var updatedSchool entities.School
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.uiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	deepcopier.Copy(result).To(&updatedSchool)
	schoolDto, msgError := validateSchool(s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(schoolDto).To(&updatedSchool)
	user, err := s.uiSchool.UpdateSchool(uint(id), updatedSchool)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   user,
	})
}
func (s *schoolService) DeleteSchool(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	schoolFindById, err := s.uiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if schoolFindById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.uiSchool.DeleteSchool(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_DELETE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}

func validateSchool(s *schoolService, c *fiber.Ctx) (dto.SchoolDTO, string) {
	var schoolDto dto.SchoolDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msgValid := validateField(dataMap)
	if msgValid != utils.EMPTY {
		return dto.SchoolDTO{}, msgValid
	}

	MapToStructSchool(&schoolDto, dataMap)
	msg = validateRequired(schoolDto)
	if msg != utils.EMPTY {
		return dto.SchoolDTO{}, msg
	}
	existEmail, _ := s.uiSchool.GetSchoolFindByEmail(schoolDto.Email)
	if existEmail.Email != utils.EMPTY {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	return schoolDto, msg
}
func MapToStructSchool(schoolDto *dto.SchoolDTO, dataMap map[string]interface{}) {
	school := dto.SchoolDTO{
		Name:    dataMap[utils.NAME].(string),
		Address: dataMap[utils.ADDRESS].(string),
		Phone:   dataMap[utils.PHONE].(string),
		Email:   dataMap[utils.EMAIL].(string),
	}
	*schoolDto = school
}
func validateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_IS_FIELD_REQUIRED
	}
	if value[utils.ADDRESS] == utils.EMPTY {
		msg = utils.ADDRESS_IS_FIELD_REQUIRED
	}
	if value[utils.PHONE] == utils.EMPTY {
		msg = utils.PHONE_IS_FIELD_REQUIRED
	}
	return msg
}

func validateRequired(user dto.SchoolDTO) string {
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
