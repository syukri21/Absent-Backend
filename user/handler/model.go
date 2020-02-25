package handler

import "backend-qrcode/user"

// UserWOHash ...
type UserWOHash struct {
	user.User
	Hash string `json:"-"`
}

// TableName ...
func (UserWOHash) TableName() string {
	return "users"
}
