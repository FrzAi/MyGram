package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string    `json:"title" gorm:"not null;type:varchar(255);" binding:"required"`
	Caption  string    `json:"caption" gorm:"not null;type:varchar(255);"`
	PhotoUrl string    `json:"photo_url" gorm:"not null;type:varchar(255);" binding:"required"`
	UserID   int       `json:"user_id"`
	Comments []Comment `json:"comments" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	User     *User     `json:"user,omitempty"`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
