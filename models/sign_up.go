package models

type SignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpRes struct {
	Status  string `json:"status"`
	Message string `json:"messages"`
}

type UserDetailsRes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
