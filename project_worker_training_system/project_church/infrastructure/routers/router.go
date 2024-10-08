package routers

import (
	"github.com/gofiber/fiber/v2"
	"member_church.com/application/handler"
)

var (
	rol handler.RolController = handler.NewRolController()
)
// CreateOrder Creating Order
//
//	@Summary		Creating Order
//	@Description	Creating Order with given request
//	@Tags			Rol
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Failure		400				{string}	string	"Bad Request"
//	@Router			/rols [post]

func NewRouter(app *fiber.App) {

	api := app.Group("/api/rol")
	api.Get("/", func(c *fiber.Ctx) error {
		return rol.GetFindAll(c)
	})
	api.Get("/:id", func(c *fiber.Ctx) error {
		return rol.GetFindById(c)
	})
	api.Post("/", func(c *fiber.Ctx) error {
		return rol.Create(c)
	})
	api.Put("/:id", func(c *fiber.Ctx) error {
		return rol.Update(c)
	})
	api.Delete("/:id", func(c *fiber.Ctx) error {
		return rol.Delete(c)
	})

}
