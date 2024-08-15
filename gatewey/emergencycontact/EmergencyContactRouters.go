package emergencycontact

import "github.com/gofiber/fiber/v2"

func NewEmergencyContactRouter(app *fiber.App) {
	app.Get("/emergency/", func(c *fiber.Ctx) error {
		return EmergencyContact(c)
	})
	app.Get("/emergency/:id", func(c *fiber.Ctx) error {
		return EmergencyContact(c)
	})
	app.Get("/emergency/user/:id", func(c *fiber.Ctx) error {
		return EmergencyContact(c)
	})
	app.Post("/emergency", func(c *fiber.Ctx) error {
		return EmergencyContact(c)
	})
	app.Put("/emergency/:id", func(c *fiber.Ctx) error {
		return EmergencyContact(c)
	})
	app.Delete("/emergency/:id", func(c *fiber.Ctx) error {
		return EmergencyContact(c)
	})
}
