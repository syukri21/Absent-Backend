package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Register ...
func Register(w http.ResponseWriter, r *http.Request) {

	var params model.RegisterParams
	json.NewDecoder(r.Body).Decode(&params)

	var admin model.Admin

	admin.User.Username = params.Username
	admin.User.Hash = admin.User.HashPassword(params.Password)
	admin.User.RoleID = 3
	admin.Fullname = params.Username
	admin.NIA = strconv.Itoa(int(time.Now().Unix()))

	if err := db.DB.Create(&admin).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&admin)

}
