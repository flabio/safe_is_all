package interfaces

import "github.com/all_is_safe/infrastructure/entities"

type IEmergencyContact interface {
	GetEmergencyContactFindAll() ([]entities.EmergencyContact, error)
	GetEmergencyContactFindById(id uint) (entities.EmergencyContact, error)
	GetEmergencyContactFindByIdUser(id uint) (entities.EmergencyContact, error)
	CreateEmergencyContact(parent entities.EmergencyContact) (entities.EmergencyContact, error)
	UpdateEmergencyContact(id uint, parent entities.EmergencyContact) (entities.EmergencyContact, error)
	DeleteEmergencyContact(id uint) (bool, error)
}
