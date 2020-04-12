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
	var count int

	var nom = r.URL.Query().Get("nom")
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	if limit == "" {
		limit = "8"
	}

	if offset == "" {
		offset = "0"
	}

	if err := db.DB.Debug().Model(&model.StudentSchedule{}).Where("schedule_id = ?", uint(scheduleID)).Count(&count).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	chain := db.DB.Preload("Student").Preload("Absent", "number_of_meeting = ?", nom).Limit(limit).Offset(offset)
	if err := chain.Find(&studentSchedules, &model.StudentSchedule{
		ScheduleID: uint(scheduleID),
	}).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	type Result struct {
		Students []model.ShowStudentSchedule `json:"students"`
		Count    int                         `json:"count"`
	}

	json.NewEncoder(w).Encode(&Result{
		Count:    count,
		Students: studentSchedules,
	})
}
