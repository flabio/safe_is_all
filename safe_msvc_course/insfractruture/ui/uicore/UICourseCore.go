package uicore

import (
	"github.com/safe_msvc_course/insfractruture/entities"
)

type UICourseCore interface {
	GetCourseFindAll() ([]entities.Course, error)
	GetCourseFindById(id uint) (entities.Course, error)
	GetCourseFindByEmail(id uint, email string) (entities.Course, error)
	CreateCourse(course entities.Course) (entities.Course, error)
	UpdateCourse(id uint, course entities.Course) (entities.Course, error)
	DeleteCourse(id uint) (bool, error)
}
