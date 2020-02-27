package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {

	var courses []Course

	if err := db.DB.Find(&courses).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&courses)

}
