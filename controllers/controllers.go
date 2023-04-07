package controllers

import (
	"encoding/json"
	"net/http"
	"user-jwt-auth/models"
	"user-jwt-auth/services"
	"user-jwt-auth/validators"

	"github.com/gin-gonic/gin"
)

var userService services.UserServiceInterface = services.UserServiceImplementation{}

func SignUpHandler(c *gin.Context) {
	req, err := validators.ValidateSignUpReq(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	_, err = userService.SignUp(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "something went wrong",
			"message": err.Error(),
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(c.Writer).Encode(models.SignUpRes{
		Status:  "success",
		Message: "user registered successfully",
	})
}

func LoginHandler(c *gin.Context) {
	req, err := validators.ValidateLoginReq(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	token, err := userService.Login(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "something went wrong",
			"message": err.Error(),
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(c.Writer).Encode(models.LoginRes{
		Status:  "success",
		Message: "user logged in successfully",
		Token:   token,
	})
}

func UserDetailsHandler(c *gin.Context) {
	email, err := validators.ValidateUserDetailsReq(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	res, err := userService.UserDetils(c, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "something went wrong",
			"message": err.Error(),
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(c.Writer).Encode(res)
}
