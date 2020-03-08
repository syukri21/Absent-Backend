package handler

import (
	course "backend-qrcode/course/handler"
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	teacher "backend-qrcode/teacher/handler"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

type ScheduleIndex struct {
	gorm.Model
	CourseID  uint            `json:"courseId"`
	TeacherID uint            `json:"teacherId"`
	Day       int             `json:"day"`
	Week      string          `json:"week"`
	Time      int             `json:"time"`
	Teacher   teacher.Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course    course.Course
	// Absents   []absent.Absent `gorm:"foreignkey:ScheduleID"`
}

func (ScheduleIndex) TableName() string {
	return "schedules"
}

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	UserID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var schedule []ScheduleIndex

	if err := db.DB.Preload("Course").Preload("Teacher").Where(&ScheduleIndex{
		TeacherID: uint(UserID),
	}).Find(&schedule).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&schedule)

}
