package services

import (
	"errors"
	"user-jwt-auth/models"
	"user-jwt-auth/models/entities"
	"user-jwt-auth/repository"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImplementation struct{}

func (u UserServiceImplementation) SignUp(ctx *gin.Context, req models.SignUpReq) (entities.Users, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return entities.Users{}, errors.New("failed to hash password")
	}

	req.Password = string(hashPassword)

	userEntity, err := repository.CreateUser(req)
	if err != nil {
		if err.Error() == "duplicated key not allowed" {
			return userEntity, errors.New("email already exists")
		}
		return userEntity, err
	}
	return userEntity, nil
}
