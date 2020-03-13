package handler

import (
	"backend-qrcode/db"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var teachers []model.TeacherBTUser
	db.DB.Preload("User").Find(&teachers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)

}
