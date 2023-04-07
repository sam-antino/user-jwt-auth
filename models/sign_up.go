package models

type SignUpReq struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type SignUpRes struct {
	Status  string `json:"status"`
	Message string `json:"messages"`
}

type UserDetailsRes struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	Email        string `json:"email"`
}
