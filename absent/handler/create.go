package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"backend-qrcode/socket"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	socketIo "backend-qrcode/socket-io"
)

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	var absent model.AbsentReturnCreate
	var params model.AbsentCreateParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var timeNow = time.Now()

	tokenParse, err := absent.VerifyToken(params.TokenAbsent)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	absent.StudentID = uint(userID)
	absent.AbsentTime = &timeNow
	absent.ScheduleID = tokenParse.ScheduleID
	absent.AbsentHash = tokenParse.AbsentHash
	absent.CourseID = tokenParse.CourseID
	absent.TeacherID = tokenParse.TeacherID
	absent.NumberOfMeeting = tokenParse.NumberOfMeeting
	absent.Semester = params.Semester

	go socketGenerateJWT(absent)

	if err := db.DB.Create(&absent).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&absent)

}

func socketGenerateJWT(absent model.AbsentReturnCreate) {

	type SocketReturn struct {
		Type socket.MessageType      `json:"type"`
		Data model.AbsentSetupReturn `json:"data"`
	}

	abs := model.Absent{
		CourseID:         absent.CourseID,
		TeacherID:        absent.TeacherID,
		NumberOfMeeting: absent.NumberOfMeeting,
	}

	token, err := abs.GenerateJWT()

	if err == nil {
		socketReturn := SocketReturn{socket.NewGenerateQrcode, model.AbsentSetupReturn{token.Token}}
		socket := socketIo.GetSocketIO()
		scheduleID := strconv.Itoa(int(absent.ScheduleID))
		socket.Server.BroadcastTo("absent."+scheduleID, "absent", socketReturn)
	}
}
