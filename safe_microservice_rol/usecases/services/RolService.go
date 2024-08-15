package services

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/msvc_rol/core/interfaces"
	"github.com/msvc_rol/core/repositories"
	"github.com/msvc_rol/globalinterfaces"
	"github.com/msvc_rol/infrastructure/entities"
	"github.com/msvc_rol/infrastructure/utils"
	"github.com/msvc_rol/usecases/dto"
	"github.com/ulule/deepcopier"
)

type rolService struct {
	Irol interfaces.IRol
}

func NewRolService() globalinterfaces.IRolGlobal {
	return &rolService{
		Irol: repositories.GetRolInstance(),
	}
}

func (rolService *rolService) GetFindAll(c *fiber.Ctx) error {
	result, err := rolService.Irol.GetFindAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   result,
	})
}
func (rolService *rolService) GetFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := rolService.Irol.GetFindById(id)

	if result.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}
func (rolService *rolService) Create(c *fiber.Ctx) error {
	var rolCreate entities.Rol
	var rol dto.RolDTO

	rolDtos, msgError := validateRol(0, rol, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	log.Println(rol)
	deepcopier.Copy(rolDtos).To(&rolCreate)
	result, err := rolService.Irol.Create(rolCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    result,
	})
}
func (rolService *rolService) Update(c *fiber.Ctx) error {

	var rolEntity entities.Rol
	var rolDto dto.RolDTO

	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}

	rolDtos, msgError := validateRol(rol.Id, rolDto, rolService, c)
	if msgError != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}

	deepcopier.Copy(rolDtos).To(&rolEntity)
	result, err := rolService.Irol.Update(rol.Id, rolEntity)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    result,
	})
}

func (rolService *rolService) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	rol, _ := rolService.Irol.GetFindById(id)
	if rol.Id == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  http.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := rolService.Irol.Delete(rol.Id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		utils.STATUS:  http.StatusCreated,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}
func validateRol(id uint, rolDto dto.RolDTO, rolService *rolService, c *fiber.Ctx) (dto.RolDTO, string) {
	var msg string = ""
	b := c.Body()
	datass := []byte(b)
	log.Println(datass)
	var dataMap map[string]interface{}
	errJson := json.Unmarshal([]byte(b), &dataMap)

	if errJson != nil {
		msg = errJson.Error()
	}

	MapToStruct(dataMap, &rolDto)

	v := validate.Struct(rolDto)
	if !v.Validate() {
		msg = v.Errors.Error()
	}
	existName, _ := rolService.Irol.GetFindByName(id, rolDto.Name)
	if existName {
		msg = utils.NAME_ALREADY_EXIST
	}
	return rolDto, msg
}
func MapToStruct(m map[string]interface{}, result interface{}) error {
	v := reflect.ValueOf(result).Elem()

	for key, value := range m {
		log.Println(key)
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
