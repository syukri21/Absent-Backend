package teacherHandler

import (
	"backend-qrcode/db"
	"backend-qrcode/user"
	"encoding/json"
	"net/http"
	"time"
)

// Teacher ...

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-"`
}

type Teacher struct {
	Model
	UserID   uint      `json:"userId"`
	Nid      string    `gorm:"unique_index" json:"nid"`
	Fullname string    `json:"fullname"`
	User     user.User `json:"-"`
}

type TeacherBTUser struct {
	Teacher
	User user.User
}

func (TeacherBTUser) TableName() string {
	return "teachers"
}

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	var teachers []TeacherBTUser
	db.DB.Preload("User").Find(&teachers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)

}
