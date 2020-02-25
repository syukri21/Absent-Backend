package handler

import (
	user "backend-qrcode/user/handler"
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
	UserID   uint      `json:"userId"`
	Nid      *string   `gorm:"unique_index" json:"nid"`
	Fullname *string   `json:"fullname"`
	User     user.User `json:"-"`
}

type JWTToken struct {
	Token        string `json:"token"`
	TeacherToken string `json:"teacherToken"`
}

type TeacherBTUser struct {
	Teacher
	User user.User `json:"user"`
}

func (TeacherBTUser) TableName() string {
	return "teachers"
}
