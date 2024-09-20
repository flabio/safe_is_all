package entities

import "time"

type Course struct {
	Id        uint       `gorm:"primary_key:auto_increment" json="id"`
	Name      string     `gorm:"type:varchar(150);not null" json="name"`
	SchoolId  uint       `gorm:"type:integer" json="school_id"`
	Active    bool       `gorm:"type:boolean" json="active"`
	CreatedAt time.Time  `gorm:"<-:created_at" json="created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json="updated_at"`
}
