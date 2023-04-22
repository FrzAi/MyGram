package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message    string    `json:"message" gorm:"not null;type:varchar(255);" binding:"required"`
	Created_At time.Time `json:"created_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	Updated_At time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	PhotoID    int       `json:"photo_id"`
	UserID     int       `json:"user_id"`
	Photo      *Photo    `json:"photo,omitempty"`
	User       *User     `json:"user,omitempty"`
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
