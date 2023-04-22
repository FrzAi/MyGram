package model

import (
	"Final_Project/helper"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string       `json:"username" gorm:"not null;unique;type:varchar(255);" binding:"required"`
	Email       string       `json:"email" gorm:"not null;unique;type:varchar(255);" binding:"required"`
	Password    string       `json:"password" gorm:"not null;type:varchar(255);" binding:"required"`
	Age         int          `json:"age" gorm:"not null;"`
	Photos      []Photo      `json:"photos,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments    []Comment    `json:"comments,omitempty" gorm:"OnUpdate:CASCADE,OnDelete:SET NULL"`
	SocialMedia *SocialMedia `json:"social_media,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
