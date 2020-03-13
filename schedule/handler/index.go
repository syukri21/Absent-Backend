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

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	UserID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var schedules []model.ScheduleIndex
	// var absents []absent.Absent

	if err := db.DB.Preload("Course").Where(&model.ScheduleIndex{
		TeacherID: uint(UserID),
	}).Find(&schedules).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&schedules)

}
