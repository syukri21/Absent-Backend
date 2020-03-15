package middleware

import (
	customHTTP "backend-qrcode/http"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var (
	student = "student"
	teacher = "teacher"
	admin   = "admin"
)

// Middleware ...
func Middleware(next http.Handler, role *string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !checkJWT(w, r) {
			return
		}

		if role != nil {
			if *role == student {
				if !isStudent(w, r) {
					return
				}
			} else if *role == teacher {
				if !isTeacher(w, r) {
					return
				}
			} else if *role == admin {
				if !isAdmin(w, r) {
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

func isAdmin(w http.ResponseWriter, r *http.Request) (ok bool) {
	roleID, err := strconv.Atoi(strings.Join(r.Header["Roleid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return false
	}

	if roleID != 3 {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: not admin")
		return false
	}

	return true

}

func isStudent(w http.ResponseWriter, r *http.Request) (ok bool) {
	roleID, err := strconv.Atoi(strings.Join(r.Header["Roleid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return false
	}

	if roleID != 2 {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: not student")
		return false
	}

	return true

}

func isTeacher(w http.ResponseWriter, r *http.Request) (ok bool) {

	roleID, err := strconv.Atoi(strings.Join(r.Header["Roleid"], ""))

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: "+err.Error())
		return false
	}

	if roleID != 1 {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error: not teacher")
		return false
	}

	return true

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
	roleID := strconv.FormatFloat(claims.(jwt.MapClaims)["role_id"].(float64), 'g', 1, 64)

	r.Header.Set("userId", userID)
	r.Header.Set("roleId", roleID)

	return true
}

type ResultParseJWT struct {
	UserID uint
	RoleID int
}

func ParseJWT(w http.ResponseWriter, r *http.Request) (J *ResultParseJWT, ok bool) {
	tokenString := r.Header.Get("Authorization")
	if len(tokenString) == 0 {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
		return nil, false
	}

	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	claims, err := VerifyToken(tokenString)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusUnauthorized, "Error verifying JWT token: "+err.Error())
		return nil, false
	}

	userID := claims.(jwt.MapClaims)["user_id"].(float64)
	roleID := claims.(jwt.MapClaims)["role_id"].(float64)

	return &ResultParseJWT{
		RoleID: int(roleID),
		UserID: uint(userID),
	}, true
}
