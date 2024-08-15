package core

import (
	"sync"

	"github.com/safe_msvc_user/insfractruture/database"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
)

func NewUserRepository() uicore.UIUserCore {
	var (
		_OPEN *OpenConnection
		_ONCE sync.Once
	)
	_ONCE.Do(func() {
		_OPEN = &OpenConnection{
			connection: database.DatabaseConnection(),
		}
	})
	return _OPEN
}

func (c *OpenConnection) GetUserFindAll() ([]dto.UserResponseDTO, error) {
	var userEntities []dto.UserResponseDTO
	c.mux.Lock()
	result := c.connection.Table("users").Order(utils.DB_ORDER_DESC).Find(&userEntities)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return userEntities, result.Error
}
func (c *OpenConnection) GetUserFindById(id uint) (entities.User, error) {
	var user entities.User
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_ID, id).Find(&user)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, result.Error
}
func (c *OpenConnection) GetUserFindByEmail(email string) (entities.User, error) {
	var user entities.User
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_EMAIL_ID, email).Find(&user)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, result.Error
}

func (c *OpenConnection) CreateUser(user entities.User) (entities.User, error) {
	c.mux.Lock()
	err := c.connection.Create(&user).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, err
}
func (c *OpenConnection) UpdateUser(id uint, user entities.User) (entities.User, error) {
	c.mux.Lock()
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Updates(&user).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return user, err
}

func (c *OpenConnection) DeleteUser(id uint) (bool, error) {
	c.mux.Lock()
	var user entities.User
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Delete(&user).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}
