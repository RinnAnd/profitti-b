package dto

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type UserRes struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Profile  string `json:"profile"`
}
