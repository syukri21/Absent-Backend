package role

import (
	"github.com/jinzhu/gorm"
)

// Role ...
type Role struct {
	gorm.Model
	Name string `gorm:"unique_index" json:"name"`
}
