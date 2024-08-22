package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/all_is_safe/core/interfaces"
	"github.com/all_is_safe/core/repository"
	"github.com/all_is_safe/globalinterfaces"
	"github.com/all_is_safe/infrastructure/entities"
	"github.com/all_is_safe/infrastructure/helpers"
	"github.com/all_is_safe/infrastructure/utils"
	"github.com/all_is_safe/usecases/dto"
	"github.com/gofiber/fiber/v2"

	"github.com/ulule/deepcopier"
)

type emergencyService struct {
	emergencyRepository interfaces.IEmergencyContact
}

func NewEmergencyContactService() globalinterfaces.IEmergencyContactGlobal {
	return &emergencyService{
		emergencyRepository: repository.GetRepositoryInstance(),
	}
}

func (s *emergencyService) GetEmergencyContactFindAll(c *fiber.Ctx) error {
	result, err := s.emergencyRepository.GetEmergencyContactFindAll()
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_QUERY,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}

func (s *emergencyService) GetEmergencyContactFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.emergencyRepository.GetEmergencyContactFindById(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_QUERY,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}

func (s *emergencyService) GetEmergencyContactFindByIdUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.emergencyRepository.GetEmergencyContactFindByIdUser(uint(id))
	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_QUERY,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (s *emergencyService) CreateEmergencyContact(c *fiber.Ctx) error {
	var emergencyContactCreate entities.EmergencyContact

	emergencyContact, msgError := validateEmergencyContact(c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(emergencyContact).To(&emergencyContactCreate)
	result, err := s.emergencyRepository.CreateEmergencyContact(emergencyContactCreate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.ERROR_CREATE,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}

func (s *emergencyService) UpdateEmergencyContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var emergencyContactUpdate entities.EmergencyContact

	emergencyContactDto, msgError := validateEmergencyContact(c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(emergencyContactDto).To(&emergencyContactUpdate)
	result, err := s.emergencyRepository.UpdateEmergencyContact(uint(id), emergencyContactUpdate)
	if result.Id == 0 {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}

func (s *emergencyService) DeleteEmergencyContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	existemergencyContact, _ := s.emergencyRepository.GetEmergencyContactFindById(uint(id))
	if existemergencyContact.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.emergencyRepository.DeleteEmergencyContact(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			utils.STATUS:  http.StatusInternalServerError,
			utils.MESSAGE: utils.ERROR_DELETE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}

func validateEmergencyContact(c *fiber.Ctx) (dto.EmergencyContactoDTO, string) {
	var emergencyContactDto dto.EmergencyContactoDTO
	var msg string = ""
	body := c.Body()

	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(body), &dataMap)

	if errJson != nil {
		msg = errJson.Error()
	}
	msgValid := helpers.ValidateField(dataMap)
	if msgValid != utils.EMPTY {
		return dto.EmergencyContactoDTO{}, msgValid
	}

	helpers.MapToStructContact(&emergencyContactDto, dataMap)

	msgReq := helpers.ValidateRequired(emergencyContactDto)
	if msgReq != utils.EMPTY {
		return dto.EmergencyContactoDTO{}, msgReq
	}
	return emergencyContactDto, msg
}
