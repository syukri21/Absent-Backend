package handler

import (
	"backend-qrcode/db"
	"encoding/json"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var teachers []TeacherBTUser
	db.DB.Preload("User").Find(&teachers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)

}
