package handler

import (
	course "backend-qrcode/course/handler"
	student "backend-qrcode/student/handler"
	teacher "backend-qrcode/teacher/handler"
	"os"

	"github.com/nleeper/goment"

	"time"

	"github.com/dgrijalva/jwt-go"
)

// Model ...
type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Absent ...
type Absent struct {
	ScheduleID       uint            `json:"scheduleId"`
	AbsentHash       string          `json:"-" gorm:"unique_index"`
	StudentID        uint            `json:"studentId" gorm:"primary_key;auto_increment:false"`
	TeacherID        uint            `json:"teacherId"`
	CourseID         uint            `json:"couresId"`
	NumberOfMeetings int             `json:"numberOfMeetings" gorm:"primary_key;auto_increment:false"`
	Semester         int             `json:"semester" gorm:"primary_key;auto_increment:false"`
	AbsentTime       *time.Time      `json:"absentTime" `
	Student          student.Student `gorm:"foreignkey:StudentID;association_foreignkey:UserID"`
	Teacher          teacher.Teacher `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course           course.Course
	Model
}

// JWTToken ...
type JWTToken struct {
	Token string `json:"token"`
}

// GenerateJWT ...
func (a Absent) GenerateJWT() (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_ABSENSI_SECRET"))

	g, err := goment.New()

	if err != nil {
		panic("err")
	}

	absentHash := "U" + g.Format("X")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":              time.Now().Add(time.Hour * 1 * 1).Unix(),
		"teacherId":        int(a.TeacherID),
		"courseId":         int(a.CourseID),
		"absentHash":       absentHash,
		"numberOfMeetings": a.NumberOfMeetings,
	})

	tokenString, err := token.SignedString(signingKey)

	return JWTToken{tokenString}, err
}
