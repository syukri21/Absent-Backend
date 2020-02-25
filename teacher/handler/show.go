package handler

import (
	"backend-qrcode/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ShowParams struct {
	ID uint `json:"id"`
}

func Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["userId"], 10, 32)

	var teacher TeacherBTUser
	db.DB.Debug().First(&teacher, &Teacher{
		UserID: uint(id),
	}).Related(&teacher.User)
	json.NewEncoder(w).Encode(teacher)

}
