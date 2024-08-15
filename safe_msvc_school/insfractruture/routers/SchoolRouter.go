package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/handler"
)

var (
	handlerSchool = handler.NewSchoolHandler()
)

func NewSchoolRouter(app *fiber.App) {
	api := app.Group("/api/school")
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerSchool.GetSchoolFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerSchool.CreateSchool(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerSchool.UpdateSchool(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerSchool.DeleteSchool(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerSchool.GetSchoolFindById(c)
	})
}
