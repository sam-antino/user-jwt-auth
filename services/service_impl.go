package services

import (
	"errors"
	"user-jwt-auth/models"
	"user-jwt-auth/models/entities"
	"user-jwt-auth/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImplementation struct{}

func (u UserServiceImplementation) SignUp(ctx *gin.Context, req models.SignUpReq) (entities.Users, error) {
	// generate hashed password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return entities.Users{}, errors.New("failed to hash password")
	}

	req.Password = string(hashPassword)

	// create user (store the hashed password in DB)
	userEntity, err := repository.CreateUser(req)
	if err != nil {
		if err.Error() == "duplicated key not allowed" {
			return userEntity, errors.New("email already exists")
		}
		return userEntity, err
	}
	return userEntity, nil
}

func (u UserServiceImplementation) Login(ctx *gin.Context, req models.LoginReq) (entities.Users, error) {
	// check the email exists or not
	userEntity, err := repository.GetUser(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userEntity, errors.New("invalid email")
		}
		return userEntity, err
	}

	// check the password is correct or not
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(req.Password))
	if err != nil {
		return userEntity, errors.New("invalid password")
	}

	// generate the jwt token

	return userEntity, nil
}
