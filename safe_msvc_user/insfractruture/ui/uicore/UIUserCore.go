package uicore

import (
	"github.com/safe_msvc_user/insfractruture/entities"
)

type UIUserCore interface {
	GetUserFindAll() ([]entities.User, error)
	GetUserFindById(id uint) (entities.User, error)
	GetUserFindByEmailAndId(id uint, email string) (bool, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id uint, user entities.User) (entities.User, error)
	DeleteUser(id uint) (bool, error)
}
