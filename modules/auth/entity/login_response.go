package entity

type LoginResponse struct {
	Id              int    `json:"-"`
	Uid             string `json:"id"`
	FullName        string `json:"fullname"`
	AccessToken     string `json:"access_token"`
	IsEmailVerified bool   `json:"is_email_verified"`
}
