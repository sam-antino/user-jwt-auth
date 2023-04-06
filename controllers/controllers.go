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

func SignUp(c *gin.Context) {
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

func Login(c *gin.Context) {
	req, err := validators.ValidateLoginReq(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "invalid request error",
			"message": err.Error(),
		})
		return
	}

	_, err = userService.Login(c, req)
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
		Message: "user logged in successfully",
	})
}
