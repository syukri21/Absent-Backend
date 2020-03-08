package handler

import (
	absent "backend-qrcode/absent/handler"
	course "backend-qrcode/course/handler"
	teacher "backend-qrcode/teacher/handler"

	"time"

	"github.com/jinzhu/gorm"
)

// Model ...
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Schedule ...
type Schedule struct {
	gorm.Model
	CourseID  uint            `json:"courseId"`
	TeacherID uint            `json:"teacherId"`
	Day       int             `json:"day"`
	Week      string          `json:"week"`
	Time      int             `json:"time"`
	Teacher   teacher.Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course    course.Course
	Absents   []absent.Absent `gorm:"foreignkey:ScheduleID"`
}
