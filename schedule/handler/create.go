package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var params model.ScheduleCreateParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var s model.ScheduleCreate

	if err := db.DB.Last(&s).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	schedule := &model.ScheduleCreate{
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
