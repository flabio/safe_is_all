package entities

import "time"

type States struct {
	Id        uint       `gorm:"primary_key:auto_increment"`
	Name      string     `gorm:"type:varchar(100);not null"`
	CityId    uint       `gorm:"null"`
	City      City       `gorm:"foreignkey:CityId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Active    bool       `gorm:"type:boolean"`
	CreatedAt time.Time  `gorm:"<-:created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" `
}
