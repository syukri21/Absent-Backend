package handler

import (
	"backend-qrcode/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ShowParams struct {
	ID uint `json:"id"`
}

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var teacher TeacherBTUser

	err := db.DB.Debug().Where("id = ? OR username = ?", params["userId"], params["userId"]).First(&teacher.User).Error

	if err != nil {
		println(err)
	}

	err = db.DB.Debug().First(&teacher, Teacher{
		UserID: teacher.User.ID,
	}).Error

	if err != nil {
		println(err)
	}

	json.NewEncoder(w).Encode(teacher)

}
