package uicore

import "github.com/safe_msvc_course/insfractruture/entities"

type UILanguage interface {
	GetLanguageFindAll() ([]entities.Language, error)
	GetLanguageFindById(id uint) (entities.Language, error)
	CreateLanguage(data entities.Language) (entities.Language, error)
	UpdateLanguageById(id uint, data entities.Language) (entities.Language, error)
	DeleteLanguageById(id uint) (bool, error)
	DuplicateLanguageName(name string) (bool, error)
}
