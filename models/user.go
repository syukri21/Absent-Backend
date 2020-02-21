package models

type User struct {
	NIM         string `gorm:"unique_index" json:"nim"`
	PhoneNumber string `gorm:"unique_index" json:"phoneNumber"`
	Name        string `json:"name"`
	Hash        string `json:"-"`
	RoleID      uint   `json:"roleID"`
	Role        Role   `json:"roles"`
}
