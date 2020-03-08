package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type ScheduleShow struct {
	ScheduleIndex
	Absents []Absent `gorm:"foreignkey:ID;association_foreignkey:ScheduleID"`
}

func (ScheduleShow) TableName() string {
	return "schedules"
}

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

	var schedule ScheduleShow
	schedule.ID = uint(scheduleID)
	schedule.TeacherID = uint(UserID)
	var absents []Absent

	err = db.DB.Debug().Preload("Course").First(&schedule).Error
	err = db.DB.Debug().Model(&schedule).Preload("Student").Related(&absents, "ScheduleID").Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	schedule.Absents = absents

	json.NewEncoder(w).Encode(&schedule)

}
