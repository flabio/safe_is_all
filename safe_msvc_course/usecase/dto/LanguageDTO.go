package dto

type LanguageDTO struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
