package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/core"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/helpers"
	"github.com/safe_msvc_course/insfractruture/ui/global"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"

	"github.com/safe_msvc_course/usecase/dto"
	"github.com/ulule/deepcopier"
)

type LanguageService struct {
	Uilenguage uicore.UILanguage
}

func NewLanguageService() global.UILanguageGlobal {
	return &LanguageService{Uilenguage: core.NewLanguageRepository()}
}

func (s *LanguageService) GetLanguageFindAll(c *fiber.Ctx) error {
	results, err := s.Uilenguage.GetLanguageFindAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   results,
	})
}
func (s *LanguageService) GetLanguageFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	resultLenguageById, err := s.Uilenguage.GetLanguageFindById(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if resultLenguageById.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusBadRequest,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   resultLenguageById,
	})
}
func (s *LanguageService) CreateLanguage(c *fiber.Ctx) error {
	var lenguageCreate entities.Language
	dataDto, msgError := ValidateLanguage(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(dataDto).To(&lenguageCreate)
	result, err := s.Uilenguage.CreateLanguage(lenguageCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_CREATE,
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}

// UpdateLenguageById implements global.UILenguageGlobal.
func (s *LanguageService) UpdateLanguageById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteLenguageById implements global.UILenguageGlobal.
func (s *LanguageService) DeleteLanguageById(c *fiber.Ctx) error {
	panic("unimplemented")
}
func ValidateLanguage(id uint, s *LanguageService, c *fiber.Ctx) (dto.LanguageDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var lenguageDto dto.LanguageDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}
	log.Println(dataMap)
	msg = helpers.ValidateField(dataMap)
	if msg != utils.EMPTY {
		return dto.LanguageDTO{}, msg
	}
	helpers.MapToStruct(dataMap, &lenguageDto)
	log.Println(lenguageDto)
	return lenguageDto, msg
}
