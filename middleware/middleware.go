package middleware

import (
	customHTTP "backend-qrcode/http"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Middleware ...
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkJWT(w, r) {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func checkJWT(w http.ResponseWriter, r *http.Request) (ok bool) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
		return false
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := VerifyToken(tokenString)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: "+err.Error())
		return false
	}

	userID := strconv.FormatFloat(claims.(jwt.MapClaims)["user_id"].(float64), 'g', 1, 64)
	r.Header.Set("userId", userID)
	return true
}
