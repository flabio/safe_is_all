package globalinterfaces

import "github.com/gofiber/fiber/v2"

type IEmergencyContactGlobal interface {
	GetEmergencyContactFindAll(c *fiber.Ctx) error
	GetEmergencyContactFindById(c *fiber.Ctx) error
	GetEmergencyContactFindByIdUser(c *fiber.Ctx) error
	CreateEmergencyContact(c *fiber.Ctx) error
	UpdateEmergencyContact(c *fiber.Ctx) error
	DeleteEmergencyContact(c *fiber.Ctx) error
}
