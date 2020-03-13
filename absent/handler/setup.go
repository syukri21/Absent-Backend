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

// Setup ...
func Setup(w http.ResponseWriter, r *http.Request) {

	var params model.AbsentSetupParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	if db.DB.First(&model.AbsentSchedule{ID: params.ScheduleID}).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "Error: no course")
		return
	}

	absent := model.Absent{
		ScheduleID:       uint(params.ScheduleID),
		CourseID:         params.CourseID,
		TeacherID:        uint(userID),
		NumberOfMeetings: params.NumberOfMeetings,
	}

	token, err := absent.GenerateJWT()

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	json.NewEncoder(w).Encode(&model.AbsentSetupReturn{token.Token})

}
