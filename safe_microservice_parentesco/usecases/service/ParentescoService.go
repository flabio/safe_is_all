package service

import (
	"encoding/json"
	"log"
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
	"github.com/gookit/validate"
	"github.com/ulule/deepcopier"
)

type parentescoService struct {
	parentescoRepository interfaces.IParentesco
}

func NewParentescoService() globalinterfaces.IParentescoGlobal {
	return &parentescoService{
		parentescoRepository: repository.GetRepositoryInstance(),
	}
}

// GetParentescoFindAll gets all parentesco records from the database.
func (s *parentescoService) GetParentescoFindAll(c *fiber.Ctx) error {
	result, err := s.parentescoRepository.GetParentescoFindAll()
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

// GetParentescoFindById gets a specific parentesco record from the database by id.
func (s *parentescoService) GetParentescoFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.parentescoRepository.GetParentescoFindById(uint(id))
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

// CreateParentesco creates a new parentesco record in the database.
func (s *parentescoService) CreateParentesco(c *fiber.Ctx) error {
	var parentescoCreate entities.Parentesco

	parentesco, msgError := validateParentesco(0, s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(parentesco).To(&parentescoCreate)
	result, err := s.parentescoRepository.CreateParentesco(parentescoCreate)
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

// UpdateParentesco updates an existing parentesco record in the database.
func (s *parentescoService) UpdateParentesco(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	var parentescoUpdate entities.Parentesco

	parentescoDto, msgError := validateParentesco(uint(id), s, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(parentescoDto).To(&parentescoUpdate)
	result, err := s.parentescoRepository.UpdateParentesco(uint(id), parentescoUpdate)
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

// DeleteParentesco deletes a specific parentesco record from the database by id.
func (s *parentescoService) DeleteParentesco(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	existParent, _ := s.parentescoRepository.GetParentescoFindById(uint(id))
	if existParent.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.parentescoRepository.DeleteParentesco(uint(id))
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

func validateParentesco(id uint, s *parentescoService, c *fiber.Ctx) (dto.ParentescoDTO, string) {
	var parentescoDto dto.ParentescoDTO
	var msg string = ""
	body := c.Body()

	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(body), &dataMap)

	if errJson != nil {
		msg = errJson.Error()
	}

	helpers.MapToStruct(dataMap, &parentescoDto)

	v := validate.Struct(parentescoDto)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := s.parentescoRepository.IsDuplicateParentescoName(id, parentescoDto.Name)
	log.Println(existName)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return parentescoDto, msg
}
