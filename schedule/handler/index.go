package handler

import (
	course "backend-qrcode/course/handler"
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Student struct {
	UserID   uint   `json:"userId"`
	Nim      string `json:"nim"`
	Fullname string `json:"fullname"`
}

type Absent struct {
	ScheduleID       uint       `json:"scheduleId" gorm:"primary_key;auto_increment:false"`
	AbsentHash       string     `json:"-" gorm:"unique_index"`
	StudentID        uint       `json:"studentId" gorm:"primary_key;auto_increment:false"`
	NumberOfMeetings int        `json:"numberOfMeetings" `
	Semester         int        `json:"semester"`
	AbsentTime       *time.Time `json:"absentTime" `
	Student          Student    `gorm:"foreignkey:StudentID;association_foreignkey:UserID"`
	Model
}

type ScheduleIndex struct {
	gorm.Model
	CourseID  uint   `json:"courseId"`
	TeacherID uint   `json:"teacherId"`
	Day       int    `json:"day"`
	Week      string `json:"week"`
	Time      int    `json:"time"`
	Course    course.Course
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

	var schedules []ScheduleIndex
	// var absents []absent.Absent

	if err := db.DB.Preload("Course").Where(&ScheduleIndex{
		TeacherID: uint(UserID),
	}).Find(&schedules).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&schedules)

}
