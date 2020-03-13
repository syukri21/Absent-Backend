package model

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}

type Teacher struct {
	Model
	UserID   uint    `json:"userId"`
	Nid      *string `gorm:"unique_index" json:"nid"`
	Fullname *string `json:"fullname"`
	User     User    `json:"-"`
}

type TeacherBTUser struct {
	Teacher
	User User `json:"user"`
}

func (TeacherBTUser) TableName() string {
	return "teachers"
}

/* -------------------------------------------------------------------------- */
/*                                  REGISTER                                  */
/* -------------------------------------------------------------------------- */

// TeacherRegisterParams ...
type TeacherRegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}