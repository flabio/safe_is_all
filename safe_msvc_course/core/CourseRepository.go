package core

import (
	"sync"

	"github.com/safe_msvc_course/insfractruture/database"
	"github.com/safe_msvc_course/insfractruture/entities"
	"github.com/safe_msvc_course/insfractruture/ui/uicore"
	"github.com/safe_msvc_course/insfractruture/utils"
	"github.com/safe_msvc_course/usecase/dto"
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

func (c *OpenConnection) GetCourseFindAll() ([]dto.CourseResponseDTO, error) {
	var courseEntities []dto.CourseResponseDTO
	c.mux.Lock()
	result := c.connection.Table("courses").Order(utils.DB_ORDER_DESC).Find(&courseEntities)
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
func (c *OpenConnection) GetCourseFindByEmail(id uint, email string) (entities.Course, error) {
	var course entities.Course
	c.mux.Lock()
	query := c.connection.Where(utils.DB_EQUAL_EMAIL_ID, email)
	if id > 0 {
		query = query.Where(utils.DB_DIFF_ID, id)
	}
	query = query.Find(&course)

	defer database.CloseConnection()
	defer c.mux.Unlock()
	return course, query.Error
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
