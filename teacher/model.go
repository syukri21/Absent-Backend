package teacher

import (
	"backend-qrcode/user"

	"github.com/jinzhu/gorm"
)

// Teacher ...
type Teacher struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	Nid      string `gorm:"unique_index" json:"nid"`
	Fullname string `json:"fullname"`
	User     user.User
}
