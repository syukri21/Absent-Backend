package student

import (
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"strconv"

	"backend-qrcode/db"
	"encoding/json"
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

	var studentSchedule []model.StudentSchedule

	if err := db.DB.Preload("Student").Preload("Course").Find(&studentSchedule, &model.StudentSchedule{
		ScheduleID: uint(scheduleID),
	}).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(studentSchedule)
}
