package repository

import (
	"sync"

	"github.com/all_is_safe/core/interfaces"
	"github.com/all_is_safe/infrastructure/database"
	"github.com/all_is_safe/infrastructure/entities"
	"github.com/all_is_safe/infrastructure/utils"
)

func GetRepositoryInstance() interfaces.IEmergencyContact {
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

func (db *OpenConnection) GetEmergencyContactFindAll() ([]entities.EmergencyContact, error) {
	var EmergencyContacts []entities.EmergencyContact
	db.mux.Lock()
	result := db.connection.Order(utils.DB_ORDER_DESC).Find(&EmergencyContacts)
	defer db.mux.Unlock()
	defer database.Closedb()
	return EmergencyContacts, result.Error
}

func (db *OpenConnection) GetEmergencyContactFindById(id uint) (entities.EmergencyContact, error) {
	var emergencyContact entities.EmergencyContact
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_ID, id).Find(&emergencyContact)
	defer db.mux.Unlock()
	defer database.Closedb()
	return emergencyContact, result.Error
}
func (db *OpenConnection) GetEmergencyContactFindByIdUser(id uint) (entities.EmergencyContact, error) {
	var emergencyContact entities.EmergencyContact
	db.mux.Lock()
	result := db.connection.Where(utils.DB_EQUAL_USER_ID, id).Find(&emergencyContact)
	defer db.mux.Unlock()
	defer database.Closedb()
	return emergencyContact, result.Error
}
func (db *OpenConnection) CreateEmergencyContact(emergency entities.EmergencyContact) (entities.EmergencyContact, error) {
	db.mux.Lock()
	err := db.connection.Create(&emergency).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return emergency, err
}

func (db *OpenConnection) UpdateEmergencyContact(id uint, emergency entities.EmergencyContact) (entities.EmergencyContact, error) {
	db.mux.Lock()
	err := db.connection.Where(utils.DB_EQUAL_ID, id).Save(&emergency).Error
	defer db.mux.Unlock()
	defer database.Closedb()
	return emergency, err
}
func (db *OpenConnection) DeleteEmergencyContact(id uint) (bool, error) {
	db.mux.Lock()
	var emergency entities.EmergencyContact
	query := db.connection.Where(utils.DB_EQUAL_ID, id).Delete(&emergency)
	defer db.mux.Unlock()
	defer database.Closedb()
	return query.RowsAffected > 0, query.Error
}
