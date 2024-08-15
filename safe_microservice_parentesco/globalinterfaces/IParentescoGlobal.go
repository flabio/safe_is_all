package globalinterfaces

import "github.com/gofiber/fiber/v2"

type IParentescoGlobal interface {
	GetParentescoFindAll(c *fiber.Ctx) error
	GetParentescoFindById(c *fiber.Ctx) error
	CreateParentesco(c *fiber.Ctx) error
	UpdateParentesco(c *fiber.Ctx) error
	DeleteParentesco(c *fiber.Ctx) error
}
