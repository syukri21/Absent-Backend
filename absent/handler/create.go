package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/socket"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
	Model
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

// CreateParams ...
type CreateParams struct {
	TokenAbsent string `json:"tokenAbsent"`
	Semester    int    `json:"semester"`
}

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	var absent AbsentReturnCreate
	var params CreateParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var timeNow = time.Now()

	tokenParse, err := absent.VerifyToken(params.TokenAbsent)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	absent.StudentID = uint(userID)
	absent.AbsentTime = &timeNow
	absent.ScheduleID = tokenParse.ScheduleID
	absent.AbsentHash = tokenParse.AbsentHash
	absent.CourseID = tokenParse.CourseID
	absent.TeacherID = tokenParse.TeacherID
	absent.NumberOfMeetings = tokenParse.NumberOfMeetings
	absent.Semester = params.Semester

	go socketGenerateJWT(absent)

	if err := db.DB.Create(&absent).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&absent)

}

func socketGenerateJWT(absent AbsentReturnCreate) {

	type SocketReturn struct {
		Type socket.MessageType `json:"type"`
		Data SetupReturn        `json:"data"`
	}

	abs := Absent{
		CourseID:         absent.CourseID,
		TeacherID:        absent.TeacherID,
		NumberOfMeetings: absent.NumberOfMeetings,
	}

	token, err := abs.GenerateJWT()

	if err == nil {
		socketReturn := SocketReturn{socket.NewGenerateQrcode, SetupReturn{token.Token}}
		socket.SendSocket(socketReturn)
	}
}
