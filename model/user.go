package model

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model        //hides from any json marshalling output
	Username   string `gorm:"unique_index" json:"username"`
	RoleID     uint   `json:"roleId"`
	Hash       string `json:"-"`
}

type JWTToken struct {
	Token string `json:"token"`
}

func (u User) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}

func (u User) GenerateJWT() (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 1 * 1).Unix(),
		"user_id":  int(u.ID),
		"username": u.Username,
		"role_id":  u.RoleID,
	})
	tokenString, err := token.SignedString(signingKey)
	return JWTToken{tokenString}, err
}

// LoginReturn ...
type LoginReturn struct {
	Token  string `json:"token"`
	RoleID uint   `josn:"roleId"`
}
