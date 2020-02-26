package handler

import (
	customHTTP "backend-qrcode/http"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(strings.Join(r.Header["Userid"], ""))
	roleID, err := strconv.Atoi(strings.Join(r.Header["Roleid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return
	}

	// var absent Absent

	json.NewEncoder(w).Encode(roleID + userID)

}
