package repositories

import (
	"sync"

	"member_church.com/core/interfaces"
	"member_church.com/infrastructure/database"
	"member_church.com/infrastructure/entities"
)

func GetUserInstance() interfaces.IUser {
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
func (db *OpenConnection) GetFindUserAll() ([]entities.User, error) {
	var users []entities.User
	db.mux.Lock()
	err := db.connection.Preload("Rol").Preload("Church").Preload("TeamPesca").Order("id desc").Find(&users).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return users, err
}
func (db *OpenConnection) GetFindUserById(id uint) (entities.User, error) {
	var user entities.User
	db.mux.Lock()
	err := db.connection.Where("id=?", id).First(&user).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) IsDuplicateEmail(id uint,email string) (entities.User, error) {
	var user entities.User
	db.mux.Lock()
	query:=db.connection.Where("email=?", email)
	if id>0 {
		query = query.Where("id<>?",id).First(&user)
	}else{
		query= query.First(&user)
	}
	err:= query.Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) IsDuplicateIdentification(id uint,identification string) (entities.User, error) {
	var user entities.User
	db.mux.Lock()
	query:=db.connection.Where("Identication=?", identification)
	if id>0 {
		query = query.Where("id<>?",id).First(&user)
	}else{
		query= query.First(&user)
	}
	err:= query.Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) IsDuplicateUserName(id uint,username string) (entities.User, error) {
	var user entities.User
	db.mux.Lock()
	query:=db.connection.Where("Username=?", username)
	if id>0 {
		query = query.Where("id<>?",id).First(&user)
	}else{
		query= query.First(&user)
	}
	err:= query.Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) CreateUser(user entities.User) (entities.User, error) {
	db.mux.Lock()
	err := db.connection.Create(&user).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) UpdateUser(id uint, user entities.User) (entities.User, error) {
	db.mux.Lock()
	err := db.connection.Where("id=?", id).Updates(&user).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	return user, err
}
func (db *OpenConnection) DeleteUser(id uint) (bool, error) {
	db.mux.Lock()
	var user entities.User
	err := db.connection.Where("id=?", id).Delete(&user).Error
	defer database.Closedb()
	defer db.mux.Unlock()
	if err == nil {
		return true, err
	}
	return false, err
}

// Add more functions here as needed for your specific use case. For example, functions for login, change password, etc.
// Remember to handle errors appropriately and release the mutex before returning.
// For example:
//    err := db.connection.Where("id=?", id).Delete(&user).Error
//    if err!= nil {
//        db.mux.Unlock()
//        return user, err
//    }
//    db.mux.Unlock()
//    return user, nil
//    // or
//    db.mux.Lock()
//    err := db.connection.Where("id=?", id).Updates(&user).Error
//    if err!= nil {
//        db.mux.Unlock()
//        return user, err
