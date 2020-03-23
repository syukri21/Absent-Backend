package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var courses []model.Course

	if err := db.DB.Order("id desc").Find(&courses).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&courses)

}
