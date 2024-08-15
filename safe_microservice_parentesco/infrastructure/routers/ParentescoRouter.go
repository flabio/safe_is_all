package routers

import (
	"github.com/all_is_safe/globalinterfaces"
	"github.com/all_is_safe/handler"
	"github.com/gofiber/fiber/v2"
)

var (
	parentescoHandler globalinterfaces.IParentescoGlobal = handler.NewParentHandler()
)

func NewParentescoRouter(app *fiber.App) {
	// TODO: Implement parentesco router
	//app.Use(middleware.Logger())

	api := app.Group("/api/parentesco")

	api.Get("/", parentescoHandler.GetParentescoFindAll)
	api.Get("/:id", parentescoHandler.GetParentescoFindById)
	api.Post("/", parentescoHandler.CreateParentesco)
	api.Put("/:id", parentescoHandler.UpdateParentesco)
	api.Delete("/:id", parentescoHandler.DeleteParentesco)
}
