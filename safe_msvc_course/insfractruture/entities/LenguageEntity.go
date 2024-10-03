package entities

import "time"

type Lenguage struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id"`
	TopicId   uint       `gorm:"type:integer" json:"topic_id"`
	Topic     Topic      `gorm:"foreignkey:TopicId;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Active    bool       `gorm:"type:boolean" json:"active"`
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
