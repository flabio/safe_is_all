package repository

import (
	"sync"

	"github.com/all_is_safe/core/interfaces"
	"github.com/all_is_safe/infrastructure/database"
	"github.com/all_is_safe/infrastructure/entities"
	"github.com/all_is_safe/infrastructure/utils"
)

func GetRepositoryInstance() interfaces.IParentesco {
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

func (db *OpenConnection) GetParentescoFindAll() ([]entities.Parentesco, error) {
	var parentescos []entities.Parentesco
	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_DESC).Find(&parentescos)
	defer db.mux.Unlock()
	defer database.Closedb()
	return parentescos, result.Error
}

func (db *OpenConnection) GetParentescoFindById(id uint) (entities.Parentesco, error) {
	var parentesco entities.Parentesco
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Find(&parentesco)
	defer db.mux.Unlock()
	defer database.Closedb()
	return parentesco, result.Error
}
func (db *OpenConnection) CreateParentesco(parent entities.Parentesco) (entities.Parentesco, error) {
	db.mux.Lock()
	err := db.connection.Create(&parent).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return parent, err
}

func (db *OpenConnection) UpdateParentesco(id uint, parent entities.Parentesco) (entities.Parentesco, error) {
	db.mux.Lock()
	err := db.connection.Where(utils.DB_EQUAL_ID, id).Save(&parent).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return parent, err
}
func (db *OpenConnection) DeleteParentesco(id uint) (bool, error) {
	db.mux.Lock()
	var parent entities.Parentesco
	query := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&parent)
	defer db.mux.Unlock()
	defer database.Closedb()
	return query.RowsAffected > 0, query.Error
}
func (db *OpenConnection) IsDuplicateParentescoName(id uint, name string) (bool, error) {
	var parent entities.Parentesco
	db.mux.Lock()
	query := db.connection.Where(utils.DB_NAME, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&parent)
	defer db.mux.Unlock()
	defer database.Closedb()
	return query.RowsAffected == 1, query.Error
	//return false, nil
}
