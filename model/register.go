package model

type RegisterParams struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Fullname *string `json:"fullname"`
}
