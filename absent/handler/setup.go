package handler

import (
	course "backend-qrcode/course/handler"
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Setup ...

// SetupParams ...
type SetupParams struct {
	ScheduleID       uint `json:"scheduleId"`
	CourseID         uint `json:"courseID"`
	NumberOfMeetings int  `json:"numberOfMeetings"`
}

// SetupReturn ...
type SetupReturn struct {
	Token string `json:"token"`
}

// Setup ...
func Setup(w http.ResponseWriter, r *http.Request) {

	var params SetupParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	if db.DB.First(&course.Course{ID: params.ScheduleID}).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "Error: no course")
		return
	}

	absent := Absent{
		ScheduleID:       params.ScheduleID,
		CourseID:         params.CourseID,
		TeacherID:        uint(userID),
		NumberOfMeetings: params.NumberOfMeetings,
	}

	token, err := absent.GenerateJWT()

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Something went wrong")
		return
	}

	json.NewEncoder(w).Encode(&SetupReturn{token.Token})

}
