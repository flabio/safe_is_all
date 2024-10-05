package core

import (
	"sync"

	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
	"github.com/safe_msvc_course/insfractruture/utils"
)

func NewLanguageRepository() uicore.UILanguage {
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

func (c *OpenConnection) GetLanguageFindAll() ([]entities.Language, error) {
	var languageEntities []entities.Language
	c.mux.Lock()
	result := c.connection.Order(utils.DB_ORDER_DESC).Find(&languageEntities)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return languageEntities, result.Error
}
func (c *OpenConnection) GetLanguageFindById(id uint) (entities.Language, error) {
	var language entities.Language
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_ID, id).Find(&language)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return language, result.Error
}

func (c *OpenConnection) CreateLanguage(language entities.Language) (entities.Language, error) {
	c.mux.Lock()
	err := c.connection.Create(&language).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return language, err
}
func (c *OpenConnection) UpdateLanguageById(id uint, language entities.Language) (entities.Language, error) {
	c.mux.Lock()
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Updates(&language).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return language, err
}

func (c *OpenConnection) DeleteLanguageById(id uint) (bool, error) {
	c.mux.Lock()
	var language entities.Language
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Delete(&language).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}

func (c *OpenConnection) DuplicateLanguageName(name string) (bool, error) {
	c.mux.Lock()
	var language entities.Language
	result := c.connection.Where(utils.DB_EQUAL_NAME, name).Delete(&language)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	if result.RowsAffected == 1 {
		return true, result.Error
	}
	return false, result.Error
}
