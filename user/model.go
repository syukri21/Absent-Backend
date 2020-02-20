package user

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int    `gorm:"primary_key" json:"id"`
	NIM         string `gorm:"unique_index" json:"nim"`
	PhoneNumber string `gorm:"unique_index" json:"phoneNumber"`
	Name        string `json:"name"`
	Hash        string `json:"-"` //hides from any json marshalling output
}

type JWTToken struct {
	Token string `json:"token"`
}

func (u User) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func (u User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u User) generateJWT() (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 1 * 1).Unix(),
		"user_id": int(u.ID),
		"name":    u.Name,
		"email":   u.PhoneNumber,
		"nim": u.NIM
	})
	tokenString, err := token.SignedString(signingKey)
	return JWTToken{tokenString}, err
}
