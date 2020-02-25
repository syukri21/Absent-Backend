package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/user"
	"net/http"

	"encoding/json"
)

// RegisterParams ...
type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Register ...
func Register(w http.ResponseWriter, r *http.Request) {

	var params RegisterParams

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	teacher := Teacher{
		User: user.User{
			Username: params.Username,
			RoleID:   1,
		},
	}

	teacher.User.Hash = teacher.User.HashPassword(params.Password)

	if err := db.DB.Debug().Create(&teacher).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(teacher)

}
