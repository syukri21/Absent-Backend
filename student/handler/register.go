package student

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var params model.RegisterParams
	json.NewDecoder(r.Body).Decode(&params)

	var student model.Student

	student.User.Username = params.Username
	student.User.Hash = student.User.HashPassword(params.Password)
	student.User.RoleID = 2
	student.Fullname = params.Username
	student.Nim = strconv.Itoa(int(time.Now().Unix()))

	if err := db.DB.Debug().Create(&student).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(student)

}
