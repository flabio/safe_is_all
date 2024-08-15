package dto

type EmergencyContactoDTO struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"first_name" validate:"required|min_len:3"  message:"required:{field} is required|min_len: min 3 caracters" label:"FirstName"`
	LastName     string `json:"last_name" validate:"required|min_len:3"  message:"required:{field} is required|min_len: min 3 caracters" label:"LastName"`
	Phone        string `json:"phone" validate:"required"  message:"required:{field} is required" label:"Phone"`
	UserId       uint   `json:"user_id" validate:"required"  message:"required:{field} is required" label:"UserId"`
	ParentescoId uint   `json:"parentesco_id" validate:"required"  message:"required:{field} is required" label:"ParentescoId"`
	Active       bool   `json:"active"`
}
