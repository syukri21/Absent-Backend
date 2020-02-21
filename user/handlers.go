package user

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var users []User
	//since we're passing a pointer to users, db.Find assigns array to the address
	db.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.DB.First(&user, params["userId"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	var user User
	var params CreateParams

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	user.PhoneNumber = params.PhoneNumber
	user.Name = params.Name
	user.NIM = params.NIM

	//get password hash
	user.Hash = user.hashPassword(params.Password)

	err = db.DB.Create(&user).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	db.DB.Where("email = ?", r.FormValue("email")).Find(&user)
	w.Header().Set("Content-Type", "application/json")
	if user.checkPassword(r.FormValue("password")) {
		token, err := user.generateJWT()
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
			return
		}
		json.NewEncoder(w).Encode(&token)
	} else {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Password incorrect")
		return
	}
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	var users []User

	db.DB.First(&user, params["userId"])
	db.DB.Delete(&user)

	db.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	reqUserId := r.Header.Get("userId")

	w.Header().Set("Content-Type", "application/json")
	if params["userId"] != reqUserId {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Not allowed to edit other users")
		return
	}
	db.DB.First(&user, params["userId"])
	db.DB.Model(&user).Update("name", r.FormValue("name"))
	json.NewEncoder(w).Encode(&user)
}
