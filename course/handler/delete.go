package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	err = db.DB.Delete(&model.Course{ID: uint(id)}).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(&model.CourseDeleted{
		ID:      uint(id),
		Message: "success",
	})

}
