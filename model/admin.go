package model

import (
	"github.com/jinzhu/gorm"
)

// Admin ...
type Admin struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	NIA      string `gorm:"unique_index" json:"nim"`
	Fullname string `json:"fullname"`
	User     User
}
