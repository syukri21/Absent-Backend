package user

type CreateParams struct {
	NIM         string `json:"nim"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}
