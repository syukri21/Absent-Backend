package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	studentID, err := strconv.Atoi(params["studentId"])

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: StatusBadRequest")
		return
	}

	scheduleID, err := strconv.Atoi(params["scheduleId"])

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: StatusBadRequest")
		return
	}

	var courseID []uint
	var semester []int

	if db.DB.Model(&model.Schedule{}).Where("id = ?", scheduleID).Pluck("course_id", &courseID).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Schedule Not Found")
		return
	}

	if db.DB.Model(&model.Course{}).Where("id = ?", courseID[0]).Pluck("semester", &semester).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Schedule Not Found")
		return
	}

	var bodyParams model.ShowGradeEntity

	err = json.NewDecoder(r.Body).Decode(&bodyParams)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	grade := &model.Grade{
		ScheduleID:  uint(scheduleID),
		CourseID:    uint(courseID[0]),
		StudentID:   uint(studentID),
		Semester:    semester[0],
		Assignment:  bodyParams.Assignment,
		Attendance:  bodyParams.Attendance,
		Uts:         bodyParams.Uts,
		Uas:         bodyParams.Uas,
		WeightValue: bodyParams.WeightValue,
		LetterValue: bodyParams.LetterValue,
	}

	if err := db.DB.Debug().Model(&model.Grade{}).Save(&grade).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
	})

}
