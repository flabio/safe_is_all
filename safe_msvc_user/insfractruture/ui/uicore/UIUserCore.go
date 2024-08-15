package uicore

import (
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/usecase/dto"
)

type UIUserCore interface {
	GetUserFindAll() ([]dto.UserResponseDTO, error)
	GetUserFindById(id uint) (entities.User, error)
	GetUserFindByEmail(email string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id uint, user entities.User) (entities.User, error)
	DeleteUser(id uint) (bool, error)
}
