package core

import (
	"sync"

	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
	"github.com/safe_msvc_course/insfractruture/utils"
)

func NewCourseRepository() uicore.UICourseCore {
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

func (c *OpenConnection) GetCourseFindAll() ([]entities.Course, error) {
	var courseEntities []entities.Course
	c.mux.Lock()
	result := c.connection.Order(utils.DB_ORDER_DESC).Find(&courseEntities)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return courseEntities, result.Error
}
func (c *OpenConnection) GetCourseFindById(id uint) (entities.Course, error) {
	var course entities.Course
	c.mux.Lock()
	result := c.connection.Where(utils.DB_EQUAL_ID, id).Find(&course)
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return course, result.Error
}

func (c *OpenConnection) CreateCourse(course entities.Course) (entities.Course, error) {
	c.mux.Lock()
	err := c.connection.Create(&course).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return course, err
}
func (c *OpenConnection) UpdateCourse(id uint, course entities.Course) (entities.Course, error) {
	c.mux.Lock()
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Updates(&course).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return course, err
}

func (c *OpenConnection) DeleteCourse(id uint) (bool, error) {
	c.mux.Lock()
	var course entities.Course
	err := c.connection.Where(utils.DB_EQUAL_ID, id).Delete(&course).Error
	defer database.CloseConnection()
	defer c.mux.Unlock()
	return err == nil, err
}
func (c *OpenConnection) IsDuplicatedCourseName(id uint, name string) (bool, error) {
	c.mux.Lock()
	var course entities.Course
	query := c.connection.Where(utils.DB_EQUAL_NAME, name)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&course)

	defer database.CloseConnection()
	defer c.mux.Unlock()
	return query.RowsAffected == 1, query.Error
}
