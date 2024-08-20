package global

import "github.com/gofiber/fiber/v2"

type UICourseGlobal interface {
	GetCourseFindAll(c *fiber.Ctx) error
	GetCourseFindById(c *fiber.Ctx) error
	CreateCourse(c *fiber.Ctx) error
	UpdateCourse(c *fiber.Ctx) error
	DeleteCourse(c *fiber.Ctx) error
}
