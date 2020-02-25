package student

import (
	"backend-qrcode/user"

	"github.com/jinzhu/gorm"
)

// Student ...
type Student struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	Nim      string `gorm:"unique_index" json:"nim"`
	Fullname string `json:"fullname"`
	User     user.User
}
