package models

type Role struct {
	ID   int    `json:"-"`
	Name string `gorm:"unique_index" json:"name"`
}
