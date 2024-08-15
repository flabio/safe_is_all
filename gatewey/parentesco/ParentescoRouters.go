package parentesco

import "github.com/gofiber/fiber/v2"

func NewParentescoRouter(app *fiber.App) {
	app.Get("/parentesco/", func(c *fiber.Ctx) error {
		return MsvcParentesco(c)
	})
	app.Get("/parentesco/:id", func(c *fiber.Ctx) error {
		return MsvcParentesco(c)
	})
	app.Post("/parentesco", func(c *fiber.Ctx) error {
		return MsvcParentesco(c)
	})
	app.Put("/parentesco/:id", func(c *fiber.Ctx) error {
		return MsvcParentesco(c)
	})
	app.Delete("/parentesco/:id", func(c *fiber.Ctx) error {
		return MsvcParentesco(c)
	})
}
