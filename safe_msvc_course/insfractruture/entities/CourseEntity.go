package entities

import "time"

type Course struct {
	Id        uint       `gorm:"primary_key:auto_increment"`
	Name      string     `gorm:"type:varchar(150);not null"`
	SchoolId  uint       `gorm:"type:integer"`
	Active    bool       `gorm:"type:boolean"`
	CreatedAt time.Time  `gorm:"<-:created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)"`
}
