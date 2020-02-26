package handler

import (
	course "backend-qrcode/course/handler"
	student "backend-qrcode/student/handler"
	teacher "backend-qrcode/teacher/handler"

	"time"

	"github.com/jinzhu/gorm"
)

// Absent ...
type Absent struct {
	gorm.Model
	StudentID  uint            `json:"studentId"`
	TeacherID  uint            `json:"teacherId"`
	CourseID   uint            `json:"couresId"`
	AbsentTime *time.Time      `json:"absentTime"`
	Student    student.Student `gorm:"foreignkey:StudentID;association_foreignkey:UserID"`
	Teacher    teacher.Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course     course.Course
}
