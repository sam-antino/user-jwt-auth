package models

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Status  string `json:"status"`
	Message string `json:"messages"`
	Token   string `json:"token"`
}
