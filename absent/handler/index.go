package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var absents []Absent

	if err := db.DB.Preload("Student").Preload("Teacher").Preload("Course").Find(&absents).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(absents)

}
