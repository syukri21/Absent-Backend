package handler

import (
	"backend-qrcode/db"
	customHTTP "backend-qrcode/http"
	"backend-qrcode/middleware"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ShowParams ...
type ShowParams struct {
	ID uint `json:"id"`
}

// Show ...
func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var teacher TeacherBTUser
	var param string
	var err error

	if params["userId"] == "" {
		j, ok := middleware.ParseJWT(w, r)
		if !ok {
			err = errors.New("Something went wrong")
		} else {
			teacher.ID = j.UserID
			param = strconv.Itoa(int(j.UserID))
		}
	} else {
		param = params["userId"]
	}

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	err = db.DB.Debug().Where("id = ? OR username = ?", param, param).First(&teacher.User).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	err = db.DB.Debug().First(&teacher, Teacher{
		UserID: teacher.User.ID,
	}).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	json.NewEncoder(w).Encode(teacher)

}
