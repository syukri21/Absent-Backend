package handler

import (
	"backend-qrcode/db"

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

	var user User

	if db.DB.Debug().First(&user, &User{
		Username: params.Username,
	}).RecordNotFound() {
		customHTTP.NewErrorResponse(w, http.StatusNotFound, "Error: NotFound")
		return
	}

	if user.CheckPassword(params.Password) {
		if token, err := user.GenerateJWT(); err != nil {
			customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: Password Wrong")
		} else {
			json.NewEncoder(w).Encode(&token)
		}
		return
	}

	customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: Password Wrong")

}
