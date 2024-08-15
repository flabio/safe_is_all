package interfaces

import "github.com/all_is_safe/infrastructure/entities"

type IParentesco interface {
	GetParentescoFindAll() ([]entities.Parentesco, error)
	GetParentescoFindById(id uint) (entities.Parentesco, error)
	CreateParentesco(parent entities.Parentesco) (entities.Parentesco, error)
	UpdateParentesco(id uint, parent entities.Parentesco) (entities.Parentesco, error)
	DeleteParentesco(id uint) (bool, error)
	IsDuplicateParentescoName(id uint, name string) (bool, error)
}
