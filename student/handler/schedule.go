package student

import (
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"strconv"

	"backend-qrcode/db"
	"net/http"

	"github.com/gorilla/mux"
)

// Schedule ...
func Schedule(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)

	scheduleID, err := strconv.Atoi(params["id"])

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Bad Request")
		return
	}

	var studentSchedules []model.ShowStudentSchedule

	var nom = r.URL.Query().Get("nom")

	if err := db.DB.Debug().Preload("Student").Preload("Course").Preload("Absent", "number_of_meeting = ?", nom).Find(&studentSchedules, &model.StudentSchedule{
		ScheduleID: uint(scheduleID),
	}).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(studentSchedules)
}
