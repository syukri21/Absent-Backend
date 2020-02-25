package student

import (
	"backend-qrcode/db"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Show ....
func Show(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)

	var student Student

	err := db.DB.Debug().Where("id = ? OR username = ?", params["userId"], params["userId"]).First(&student.User).Error

	if err != nil {
		println(err)
	}

	err = db.DB.Debug().First(&student, Student{
		UserID: student.User.ID,
	}).Error

	if err != nil {
		println(err)
	}

	json.NewEncoder(w).Encode(student)
}
