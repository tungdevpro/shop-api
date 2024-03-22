package entity

type SignUpResponse struct {
	Id              int    `json:"-"`
	Uid             string `json:"id"`
	Email           string `json:"email"`
	FullName        string `json:"fullname"`
	IsEmailVerified bool   `json:"is_email_verified"`

	// AccessToken string `json:"access_token"`
}
