package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/handler"
	"github.com/safe_msvc_user/middleware"
)

var (
	handlerUser = handler.NewUserHandler()
)

func NewUserRouter(app *fiber.App) {

	api := app.Group("/api/user")
	api.Use(middleware.GetToken)
	api.Get("/", func(c *fiber.Ctx) error {
		return handlerUser.GetUserFindAll(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return handlerUser.CreateUser(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return handlerUser.UpdateUser(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return handlerUser.DeleteUser(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return handlerUser.GetUserFindById(c)
	})
}
