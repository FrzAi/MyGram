package model

import (
	"time"
)

type GormModel struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Created_At time.Time `json:"created_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
	Updated_At time.Time `json:"updated_at" gorm:"type:TIMESTAMP WITHOUT TIME ZONE;default:CURRENT_TIMESTAMP"`
}
