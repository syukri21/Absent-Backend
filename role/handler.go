package role

import (
	"backend-qrcode/db"
	"backend-qrcode/models"

	customHTTP "backend-qrcode/http"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var roles []Role
	//since we're passing a pointer to roles, db.Find assigns array to the address
	db.DB.Find(&roles)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// ShowHandler ...
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var roles []models.Role
	// var users []interface{}
	// var users []models.User

	db.DB.First(&roles, params["roleId"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// CreateHandler ...
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var role Role
	var params CreateParams
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	role.Name = params.Name

	err = db.DB.Create(&role).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&role)
}
