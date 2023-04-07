package validators

import (
	"errors"
	"net/http"
	"user-jwt-auth/models"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func ValidateSignUpReq(c *gin.Context) (req models.SignUpReq, err error) {
	err = c.ShouldBindJSON(&req)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return req, err
	}

	opts := govalidator.Options{
		Data: &req,
		Rules: govalidator.MapData{
			"name":          []string{"required"},
			"mobile_number": []string{"required", "numeric", "len:10"},
			"email":         []string{"required", "email"},
			"password":      []string{"required"},
		},
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return models.SignUpReq{}, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	return req, nil
}

func ValidateLoginReq(c *gin.Context) (req models.LoginReq, err error) {
	err = c.ShouldBindJSON(&req)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return req, err
	}

	opts := govalidator.Options{
		Data: &req,
		Rules: govalidator.MapData{
			"email":    []string{"required", "email"},
			"password": []string{"required"},
		},
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		for param, message := range e {
			return models.LoginReq{}, errors.New("param: " + param + ", message:" + message[0])
		}
	}

	return req, nil
}

func ValidateUserDetailsReq(c *gin.Context) (string, error) {
	email := c.Request.Header["User-Email"][0]
	if email == "" {
		return "", errors.New("no email passed in the headers")
	}

	return email, nil
}
