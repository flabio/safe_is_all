package interfaces

import "member_church.com/infrastructure/entities"

type IUser interface {
	GetFindUserAll() ([]entities.User, error)
	GetFindUserById(id uint) (entities.User, error)
	IsDuplicateEmail(id uint,email string) (entities.User, error)
	IsDuplicateIdentification(id uint,identification string) (entities.User, error)
	IsDuplicateUserName(id uint,username string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
	UpdateUser(id uint, user entities.User) (entities.User, error)
	DeleteUser(id uint) (bool, error)
}
