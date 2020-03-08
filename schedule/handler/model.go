package handler

import (
	absent "backend-qrcode/absent/handler"
	course "backend-qrcode/course/handler"
	teacher "backend-qrcode/teacher/handler"

	"time"
)

// Model ...
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Schedule ...
type Schedule struct {
	ID        uint            `json:"id" gorm:"not null;unique_index;AUTO_INCREMENT"`
	CourseID  uint            `json:"courseId" gorm:"primary_key;auto_increment:false"`
	TeacherID uint            `json:"teacherId" gorm:"primary_key;auto_increment:false"`
	Day       int             `json:"day"`
	Week      string          `json:"week"`
	Time      int             `json:"time"`
	Teacher   teacher.Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course    course.Course
	Absents   []absent.Absent `gorm:"foreignkey:ScheduleID"`
}
