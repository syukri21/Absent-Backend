package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
)

// Edit ...
func Edit(w http.ResponseWriter, r *http.Request) {

	var params model.Teacher

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	var teacher model.Teacher

	teacher.UserID = params.UserID
	if db.DB.First(&teacher).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Record Not Found")
		return
	}

	if err := db.DB.Debug().Model(&params).Update(&params).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
	})

}
