package entities

import "time"

type EmergencyContact struct {
	Id           uint       `gorm:"primary_key:auto_increment" json:"id"`
	FirstName    string     `gorm:"type:varchar(150)" json:"first_name"`
	LastName     string     `gorm:"type:varchar(150)" json:"last_name"`
	Phone        string     `gorm:"type:varchar(100)" json:"phone"`
	UserId       uint       `gorm:"type:integer" json:"user_id"`
	ParentescoId uint       `gorm:"type:integer" json:"parentesco_id"`
	Active       bool       `gorm:"type:boolean" json:"active"`
	CreatedAt    time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"type:TIMESTAMP(6)"  json:"updated_at"`
}
