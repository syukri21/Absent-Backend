package model

import (
	"os"

	"github.com/nleeper/goment"

	"time"

	"github.com/dgrijalva/jwt-go"
)

// AbesntModel ...
type AbesntModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Absent ...
type Absent struct {
	ScheduleID       uint       `json:"scheduleId" gorm:"primary_key;auto_increment:false"`
	AbsentHash       string     `json:"-" gorm:"unique_index"`
	StudentID        uint       `json:"studentId" gorm:"primary_key;auto_increment:false"`
	TeacherID        uint       `json:"teacherId"`
	CourseID         uint       `json:"couresId"`
	NumberOfMeetings int        `json:"numberOfMeetings" gorm:"primary_key;auto_increment:false"`
	Semester         int        `json:"semester"`
	AbsentTime       *time.Time `json:"absentTime" `
	Student          Student    `gorm:"foreignkey:StudentID;association_foreignkey:UserID"`
	Teacher          Teacher    `gorm:"foreignkey:TeacherID;association_foreignkey:UserID"`
	Course           Course
	AbesntModel
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
		"scheduleId":       int(a.ScheduleID),
		"absentHash":       absentHash,
		"numberOfMeetings": a.NumberOfMeetings,
	})

	tokenString, err := token.SignedString(signingKey)

	return JWTToken{Token: tokenString}, err
}

// AbsentReturnCreate ...
type AbsentReturnCreate struct {
	StudentID        uint       `json:"studentId"`
	TeacherID        uint       `json:"teacherId"`
	CourseID         uint       `json:"couresId"`
	ScheduleID       uint       `json:"scheduleId"`
	NumberOfMeetings int        `json:"numberOfMeetings"`
	Semester         int        `json:"semester" `
	AbsentTime       *time.Time `json:"absentTime"`
	AbsentHash       string     `json:"-" gorm:"unique_index"`
	AbesntModel
}

// TableName ...
func (AbsentReturnCreate) TableName() string {
	return "absents"
}

// TokenParse ...
type TokenParse struct {
	TeacherID        uint   `json:"teacherId"`
	CourseID         uint   `json:"courseId"`
	ScheduleID       uint   `json:"scheduleId"`
	AbsentHash       string `json:"absentHash"`
	NumberOfMeetings int    `json:"numberOfMeetings"`
}

// VerifyToken ...
func (a AbsentReturnCreate) VerifyToken(tokenString string) (*TokenParse, error) {

	signingKey := []byte(os.Getenv("JWT_ABSENSI_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	tokenParse := TokenParse{
		CourseID:         uint(token.Claims.(jwt.MapClaims)["courseId"].(float64)),
		TeacherID:        uint(token.Claims.(jwt.MapClaims)["teacherId"].(float64)),
		ScheduleID:       uint(token.Claims.(jwt.MapClaims)["scheduleId"].(float64)),
		AbsentHash:       token.Claims.(jwt.MapClaims)["absentHash"].(string),
		NumberOfMeetings: int(token.Claims.(jwt.MapClaims)["numberOfMeetings"].(float64)),
	}

	return &tokenParse, err
}

// AbsentCreateParams ...
type AbsentCreateParams struct {
	TokenAbsent string `json:"tokenAbsent"`
	Semester    int    `json:"semester"`
}

// SetupParams ...
type AbsentSetupParams struct {
	ScheduleID       uint `json:"scheduleId"`
	CourseID         uint `json:"courseID"`
	NumberOfMeetings int  `json:"numberOfMeetings"`
}

// SetupReturn ...
type AbsentSetupReturn struct {
	Token string `json:"token"`
}

// AbsentSchedule ...
type AbsentSchedule struct {
	ID        uint `json:"string"`
	CourseID  uint `json:"courseId"`
	TeacherID uint `json:"teacherId"`
}

func (AbsentSchedule) TableName() string {
	return "schedules"
}
