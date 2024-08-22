package handler

import (
	"github.com/all_is_safe/globalinterfaces"
	"github.com/all_is_safe/usecases/service"
	"github.com/gofiber/fiber/v2"
)

type emergencyContactHandler struct {
	emergency globalinterfaces.IEmergencyContactGlobal
}

func NewEmergencyContactHandler() globalinterfaces.IEmergencyContactGlobal {
	return &emergencyContactHandler{
		emergency: service.NewEmergencyContactService(),
	}
}

func (h *emergencyContactHandler) GetEmergencyContactFindAll(c *fiber.Ctx) error {
	return h.emergency.GetEmergencyContactFindAll(c)
}
func (h *emergencyContactHandler) GetEmergencyContactFindById(c *fiber.Ctx) error {
	return h.emergency.GetEmergencyContactFindById(c)
}
func (h *emergencyContactHandler) GetEmergencyContactFindByIdUser(c *fiber.Ctx) error {
	return h.emergency.GetEmergencyContactFindByIdUser(c)
}
func (h *emergencyContactHandler) CreateEmergencyContact(c *fiber.Ctx) error {
	return h.emergency.CreateEmergencyContact(c)
}
func (h *emergencyContactHandler) UpdateEmergencyContact(c *fiber.Ctx) error {
	return h.emergency.UpdateEmergencyContact(c)
}
func (h *emergencyContactHandler) DeleteEmergencyContact(c *fiber.Ctx) error {
	return h.emergency.DeleteEmergencyContact(c)
}
