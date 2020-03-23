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

	var params model.CourseEditParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	var course model.Course
	course.ID = params.ID

	if db.DB.Debug().First(&course).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: Record Not Found")
		return
	}

	course.Name = params.Name
	course.Semester = params.Semester
	course.TotalSks = params.TotalSks

	if err := db.DB.Save(&course).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "success",
	})

}
