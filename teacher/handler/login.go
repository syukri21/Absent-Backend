package handler

import (
	"backend-qrcode/db"
	"backend-qrcode/user"

	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
)

// LoginParams ...
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {

	var params LoginParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	var teacher TeacherBTUser

	if db.DB.Debug().First(&teacher.User, &user.User{
		Username: params.Username,
	}).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "Error: NotFound")
		return
	}

	if teacher.User.CheckPassword(params.Password) {
		if token, err := teacher.User.GenerateJWT(); err != nil {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: Password Wrong")
		} else {
			json.NewEncoder(w).Encode(&token)
		}
		return
	}

	customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: Password Wrong")

}
