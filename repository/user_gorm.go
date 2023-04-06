package repository

import (
	"user-jwt-auth/initializers"
	"user-jwt-auth/models"
	"user-jwt-auth/models/entities"
)

func CreateUser(req models.SignUpReq) (user entities.Users, err error) {
	user.Email = req.Email
	user.Name = req.Name
	user.Password = req.Password
	err = initializers.Db.Create(&user).Error
	if err != nil {
		return entities.Users{}, err

	}
	return user, nil
}

func GetUser(email string) (user entities.Users, err error) {
	err = initializers.Db.First(&user, "email=?", email).Error
	if err != nil {
		return entities.Users{}, err
	}
	return user, nil
}
