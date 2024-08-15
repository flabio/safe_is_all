package core

import (
	"sync"

	"github.com/safe_msvc_user/insfractruture/database"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
)

func NewSchoolRepository() uicore.UISchoolCore {
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

func (c *OpenConnection) GetSchoolFindAll() ([]dto.SchoolResponseDTO, error) {
	var schoolEntities []dto.SchoolResponseDTO
	c.mux.Lock()
	result := c.connection.Table("schools").Order(utils.DB_ORDER_DESC).Find(&schoolEntities)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return schoolEntities, result.Error
}
func (c *OpenConnection) GetSchoolFindById(id uint) (entities.School, error) {
	var school entities.School
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_ID, id).Find(&school)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, result.Error
}
func (c *OpenConnection) GetSchoolFindByEmail(email string) (entities.School, error) {
	var school entities.School
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_EMAIL_ID, email).Find(&school)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, result.Error
}

func (c *OpenConnection) CreateSchool(school entities.School) (entities.School, error) {
	c.mux.Lock()
	err := c.connection.Create(&school).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, err
}
func (c *OpenConnection) UpdateSchool(id uint, school entities.School) (entities.School, error) {
	c.mux.Lock()
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Updates(&school).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return school, err
}

func (c *OpenConnection) DeleteSchool(id uint) (bool, error) {
	c.mux.Lock()
	var school entities.School
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Delete(&school).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}
