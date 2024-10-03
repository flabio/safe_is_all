package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_course/handler"
)

var (
	handlerCourse = handler.NewCourseHandler()
)

func NewCourseRouter(app *fiber.App) {
	api := app.Group("/api/course")
	//api.Use(middleware.ValidateToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerCourse.GetCourseFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerCourse.CreateCourse(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerCourse.UpdateCourse(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerCourse.DeleteCourse(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerCourse.GetCourseFindById(c)
	})
}
