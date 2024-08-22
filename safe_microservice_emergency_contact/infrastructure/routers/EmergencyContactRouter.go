package routers

import (
	"github.com/all_is_safe/globalinterfaces"
	"github.com/all_is_safe/handler"
	"github.com/all_is_safe/infrastructure/middleware"
	"github.com/gofiber/fiber/v2"
)

var (
	emergencyHandler globalinterfaces.IEmergencyContactGlobal = handler.NewEmergencyContactHandler()
)

func NewEmergencyRouter(app *fiber.App) {
	// TODO: Implement parentesco router
	//app.Use(middleware.Logger())

	api := app.Group("/api/emergency")
	api.Use(middleware.ValidationToken)
	api.Get("/", emergencyHandler.GetEmergencyContactFindAll)
	api.Get("/:id", emergencyHandler.GetEmergencyContactFindById)
	api.Get("/user/:id", emergencyHandler.GetEmergencyContactFindByIdUser)
	api.Post("/", emergencyHandler.CreateEmergencyContact)
	api.Put("/:id", emergencyHandler.UpdateEmergencyContact)
	api.Delete("/:id", emergencyHandler.DeleteEmergencyContact)
}
