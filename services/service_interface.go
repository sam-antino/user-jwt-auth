package services

import (
	"user-jwt-auth/models"
	"user-jwt-auth/models/entities"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	SignUp(ctx *gin.Context, req models.SignUpReq) (entities.Users, error)
}
