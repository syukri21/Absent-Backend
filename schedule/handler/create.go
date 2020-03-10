package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// CreateParams ...
type CreateParams struct {
	CourseID uint   `json:"courseId"`
	Day      int    `json:"day"`
	Week     string `json:"week"`
	Time     int    `json:"time"`
}

// CreateSchedule ...
type CreateSchedule struct {
	ID        uint   `json:"id"`
	CourseID  uint   `json:"courseId"`
	TeacherID uint   `json:"teacherId"`
	Day       int    `json:"day"`
	Week      string `json:"week"`
	Time      int    `json:"time"`
}

// TableName ...
func (CreateSchedule) TableName() string {
	return "schedules"

}

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var params CreateParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var s CreateSchedule

	if err := db.DB.Last(&s).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	schedule := &CreateSchedule{
		ID:        s.ID + 1,
		TeacherID: uint(userID),
		CourseID:  params.CourseID,
		Day:       params.Day,
		Week:      params.Week,
		Time:      params.Time,
	}

	if err := db.DB.Create(&schedule).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&schedule)

}

// CourseID:  1,
// Day:       1,
// Week:      "BOTH",
// Time:      200,
