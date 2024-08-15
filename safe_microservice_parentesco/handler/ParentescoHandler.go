package handler

import (
	"github.com/all_is_safe/globalinterfaces"
	"github.com/all_is_safe/usecases/service"
	"github.com/gofiber/fiber/v2"
)

type parentHandler struct {
	parent globalinterfaces.IParentescoGlobal
}

func NewParentHandler() globalinterfaces.IParentescoGlobal {
	return &parentHandler{
		parent: service.NewParentescoService(),
	}
}

func (h *parentHandler) GetParentescoFindAll(c *fiber.Ctx) error {
	return h.parent.GetParentescoFindAll(c)
}
func (h *parentHandler) GetParentescoFindById(c *fiber.Ctx) error {
	return h.parent.GetParentescoFindById(c)
}
func (h *parentHandler) CreateParentesco(c *fiber.Ctx) error {
	return h.parent.CreateParentesco(c)
}
func (h *parentHandler) UpdateParentesco(c *fiber.Ctx) error {
	return h.parent.UpdateParentesco(c)
}
func (h *parentHandler) DeleteParentesco(c *fiber.Ctx) error {
	return h.parent.DeleteParentesco(c)
}
