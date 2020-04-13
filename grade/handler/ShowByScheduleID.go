package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ShowByScheduleID ...
func ShowByScheduleID(w http.ResponseWriter, r *http.Request) {

	scheduleID := mux.Vars(r)["id"]

	type Result struct {
		Students []model.ShowGradeByScheduleID `json:"students"`
		Count    int                           `json:"count"`
	}

	var students []model.ShowGradeByScheduleID

	var tx = db.DB.Debug().Preload("Grade").Preload("Student")
	limit := r.URL.Query().Get("limit")
	if limit != "" {
		tx = tx.Limit(limit)
	}
	offset := r.URL.Query().Get("offset")
	if offset != "" {
		tx = tx.Offset(offset)
	}
	if err := tx.Find(&students, "schedule_id = ?", scheduleID).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	var count int
	db.DB.Model(&model.StudentSchedule{}).Where("schedule_id = ?", scheduleID).Count(&count)

	json.NewEncoder(w).Encode(&Result{
		Count:    count,
		Students: students,
	})
}
