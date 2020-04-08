package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var absents []model.Absent
	params := r.URL.Query()
	scheduleID, _ := strconv.Atoi(params.Get("scheduleId"))
	limit, _ := strconv.Atoi(params.Get("limit"))
	offset, _ := strconv.Atoi(params.Get("offset"))

	if limit == 0 {
		limit = 5
	}

	if err := db.DB.Offset(offset).Limit(limit).Preload("Student").Find(&absents, &model.Absent{
		ScheduleID: uint(scheduleID),
	}).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(absents)

}
