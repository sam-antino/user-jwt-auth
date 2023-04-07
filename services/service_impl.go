package services

import (
	"errors"
	"os"
	"time"
	"user-jwt-auth/models"
	"user-jwt-auth/models/entities"
	"user-jwt-auth/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func (u UserServiceImplementation) Login(ctx *gin.Context, req models.LoginReq) (string, error) {
	// check the email exists or not
	userEntity, err := repository.GetUser(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid email")
		}
		return "", err
	}

	// check the password is correct or not
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// generate the jwt token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userEntity.ID,
		"email": userEntity.Email,
		"exp":   time.Now().Add(time.Minute * 1).Unix(), // expiry of 1 mins
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return tokenString, errors.New("could not create token")
	}

	return tokenString, nil
}

func (u UserServiceImplementation) UserDetils(ctx *gin.Context, email string) (models.UserDetailsRes, error) {
	userEntity, err := repository.GetUser(email)
	if err != nil {
		return models.UserDetailsRes{}, err
	}

	return models.UserDetailsRes{
		Name:         userEntity.Name,
		MobileNumber: userEntity.MobileNumber,
		Email:        userEntity.Email,
	}, nil
}
