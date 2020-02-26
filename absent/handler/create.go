package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// AbsentReturnCreate ...
type AbsentReturnCreate struct {
	gorm.Model
	StudentID  uint       `json:"studentId"`
	TeacherID  uint       `json:"teacherId"`
	CourseID   uint       `json:"couresId"`
	AbsentTime *time.Time `json:"absentTime"`
}

// TableName ...
func (AbsentReturnCreate) TableName() string {
	return "absents"
}

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var timeNow = time.Now()

	absent := AbsentReturnCreate{
		AbsentTime: &timeNow,
		CourseID:   1,
		StudentID:  uint(userID),
		TeacherID:  1,
	}

	if err := db.DB.Create(&absent).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&absent)

}
