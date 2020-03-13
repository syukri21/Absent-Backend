package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Show ...
func Show(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	scheduleID, err := strconv.Atoi(params["scheduleId"])

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	UserID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var schedule model.ScheduleShow
	var absents []model.Absent

	isNotFound := db.DB.Debug().Preload("Course").First(&schedule, &model.Schedule{
		ID:        uint(scheduleID),
		TeacherID: uint(UserID),
	}).RecordNotFound()

	if isNotFound {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: not found")
		return
	}

	err = db.DB.Debug().Model(&schedule).Preload("Student").Related(&absents, "ScheduleID").Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	schedule.Absents = absents

	json.NewEncoder(w).Encode(&schedule)

}
