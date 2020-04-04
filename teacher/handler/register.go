package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"net/http"
	"strconv"
	"time"

	"encoding/json"
)

// Register ...
func Register(w http.ResponseWriter, r *http.Request) {

	var params model.TeacherRegisterParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	teacher := model.Teacher{
		User: model.User{
			Username: params.Username,
			RoleID:   1,
		},
	}

	nid := strconv.Itoa(int(time.Now().Unix()))

	teacher.User.Hash = teacher.User.HashPassword(params.Password)
	teacher.Fullname = &params.Fullname
	teacher.Nid = &nid

	// Check There fullname or not
	println(*teacher.Fullname)
	if &teacher.Fullname == nil || *teacher.Fullname == "" {
		teacher.Fullname = &params.Username
	}

	if err := db.DB.Create(&teacher).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(teacher)

}
