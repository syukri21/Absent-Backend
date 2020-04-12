package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// ShowByScheduleID ...
func ShowByScheduleID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(params["id"])
}
