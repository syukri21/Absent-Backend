package student

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
)

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var params RegisterParams
	json.NewDecoder(r.Body).Decode(&params)

	var student Student

	student.User.Username = params.Username
	student.User.Hash = student.User.HashPassword(params.Password)
	student.User.RoleID = 2

	if err := db.DB.Debug().Create(&student).Error; err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(student)

}
