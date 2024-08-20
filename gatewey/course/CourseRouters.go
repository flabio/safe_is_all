package course

import "github.com/gofiber/fiber/v2"

func NewCourseRouter(app *fiber.App) {
	app.Get("/course/", func(c *fiber.Ctx) error {
		return MsvcCourse(c)
	})
	app.Get("/course/:id", func(c *fiber.Ctx) error {
		return MsvcCourse(c)
	})
	app.Post("/course", func(c *fiber.Ctx) error {
		return MsvcCourse(c)
	})
	app.Put("/course/:id", func(c *fiber.Ctx) error {
		return MsvcCourse(c)
	})
	app.Delete("/course/:id", func(c *fiber.Ctx) error {
		return MsvcCourse(c)
	})
}
