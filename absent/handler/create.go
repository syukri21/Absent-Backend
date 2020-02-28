package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// AbsentReturnCreate ...
type AbsentReturnCreate struct {
	gorm.Model
	StudentID       uint       `json:"studentId"`
	TeacherID       uint       `json:"teacherId"`
	CourseID        uint       `json:"couresId"`
	NumberOfMeeting int        `json:"numberOfMeeting" `
	Semester        int        `json:"semester" `
	AbsentTime      *time.Time `json:"absentTime"`
	AbsentHash      string     `json:"-" gorm:"unique_index"`
}

// TableName ...
func (AbsentReturnCreate) TableName() string {
	return "absents"
}

// TokenParse ...
type TokenParse struct {
	TeacherID        uint   `json:"teacherId"`
	CourseID         uint   `json:"courseId"`
	AbsentHash       string `json:"absentHash"`
	NumberOfMeetings string `json:"numberOfMeetings"`
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
		CourseID:   uint(token.Claims.(jwt.MapClaims)["courseId"].(float64)),
		TeacherID:  uint(token.Claims.(jwt.MapClaims)["teacherId"].(float64)),
		AbsentHash: token.Claims.(jwt.MapClaims)["absentHash"].(string),
	}

	return &tokenParse, err
}

// CreateParams ...
type CreateParams struct {
	TokenAbsent string `json:"tokenAbsent"`
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
	absent.AbsentHash = tokenParse.AbsentHash
	absent.CourseID = tokenParse.CourseID
	absent.TeacherID = tokenParse.TeacherID

	if err := db.DB.Create(&absent).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&absent)

}
