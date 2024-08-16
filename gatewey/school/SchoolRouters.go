package school

import "github.com/gofiber/fiber/v2"

func NewSchoolRouter(app *fiber.App) {
	app.Get("/school/", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Get("/school/:id", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Post("/school", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Put("/school/:id", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Delete("/school/:id", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
}
