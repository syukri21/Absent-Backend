package handler

import (
	course "backend-qrcode/course/handler"
	student "backend-qrcode/student/handler"
	teacher "backend-qrcode/teacher/handler"
	"os"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Absent ...
type Absent struct {
	gorm.Model
	AbsentHash string          `json:"-" gorm:"unique_index"`
	StudentID  uint            `json:"studentId"`
	TeacherID  uint            `json:"teacherId"`
	CourseID   uint            `json:"couresId"`
	AbsentTime *time.Time      `json:"absentTime"`
	Student    student.Student `gorm:"foreignkey:StudentID;association_foreignkey:UserID"`
	Teacher    teacher.Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course     course.Course
}

// JWTToken ...
type JWTToken struct {
	Token string `json:"token"`
}

// GenerateJWT ...
func (a Absent) GenerateJWT() (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_ABSENSI_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":         time.Now().Add(time.Hour * 1 * 1).Unix(),
		"teacher_id":  int(a.TeacherID),
		"course_id":   int(a.CourseID),
		"absent_hash": "aa",
	})
	tokenString, err := token.SignedString(signingKey)
	return JWTToken{tokenString}, err
}
