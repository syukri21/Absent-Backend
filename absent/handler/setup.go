package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
)

// Setup ...
func Setup(w http.ResponseWriter, r *http.Request) {

	var params model.AbsentSetupParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	schedule := &model.AbsentSchedule{
		ID: params.ScheduleID,
	}

	if db.DB.First(&schedule).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "Error: no course")
		return
	}

	absent := model.Absent{
		ScheduleID:       schedule.ID,
		CourseID:         schedule.CourseID,
		TeacherID:        schedule.TeacherID,
		NumberOfMeeting: params.NumberOfMeeting,
	}

	token, err := absent.GenerateJWT()

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	json.NewEncoder(w).Encode(&model.AbsentSetupReturn{token.Token})

}
