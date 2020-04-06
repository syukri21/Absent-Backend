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
	schedule.ID = uint(scheduleID)
	schedule.TeacherID = uint(UserID)

	var absents []model.Absent

	query := r.URL.Query()

	isNotFound := db.DB.Preload("Course").First(&schedule, schedule).RecordNotFound()

	if isNotFound {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: not found")
		return
	}

	tx := db.DB.Model(&schedule)

	if query["nom"] != nil {
		numberOfMeeting, err := strconv.Atoi(query["nom"][0])
		if err == nil {
			tx = db.DB.Model(&schedule).Where(&model.Absent{
				NumberOfMeeting: numberOfMeeting,
			})
		}
	}

	err = tx.Preload("Student").Related(&absents, "ScheduleID", "NumberOfMeeting").Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	schedule.Absents = absents

	json.NewEncoder(w).Encode(&schedule)

}
