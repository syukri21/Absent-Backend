package handler

import (
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"

	"backend-qrcode/db"
	"encoding/json"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var admins []model.Admin

	if err := db.DB.Preload("User").Find(&admins).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&admins)
}
