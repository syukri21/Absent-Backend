package teacher

import (
	"github.com/jinzhu/gorm"
)

// Teacher ...
type Teacher struct {
	gorm.Model
	UserID   uint   `json:"userId"`
	NID      uint   `gorm:"unique_index" json:"nid"`
	Fullname string `json:"fullname"`
}
