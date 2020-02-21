package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model         //hides from any json marshalling output
	NIM         string `gorm:"unique_index" json:"nim"`
	PhoneNumber string `gorm:"unique_index" json:"phoneNumber"`
	Name        string `json:"name"`
	Hash        string `json:"-"`
	RoleID      uint   `json:"roleID"`
	Role        Role   `json:"roles"`
}
