package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/middleware"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Show ....
func Show(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	var param string

	if params["userId"] == "" {
		j, ok := middleware.ParseJWT(w, r)
		if !ok {
			return
		} else {
			param = strconv.Itoa(int(j.UserID))
		}
	} else {
		param = params["userId"]
	}

	var admin model.Admin

	err := db.DB.Where("id = ? OR username = ?", param, param).First(&admin.User).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return

	}

	err = db.DB.First(&admin, model.Admin{
		UserID: admin.User.ID,
	}).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&admin)
}
