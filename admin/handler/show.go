package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Show ....
func Show(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)

	var admin model.Admin

	err := db.DB.Debug().Where("id = ? OR username = ?", params["userId"], params["userId"]).First(&admin.User).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return

	}

	err = db.DB.Debug().First(&admin, model.Admin{
		UserID: admin.User.ID,
	}).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&admin)
}
