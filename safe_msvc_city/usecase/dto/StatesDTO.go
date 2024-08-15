package dto

type StatesDTO struct {
	Id     uint   `json:"id" `
	Name   string `json:"name" validate:"required|min_len:3"  message:"required:{field} is required|min_len: min 3 caracters" label:"Name"`
	CityId uint   `json:"city_id" validate:"required" message:"required: {field} is required" label:"City"`
	Active bool   `json:"active" validate:"required" message:"required:{field} is required" label:"Active"`
}
