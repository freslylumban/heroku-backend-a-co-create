package service

import (
	"heroku-backend-a-cocreate/dto"
	"heroku-backend-a-cocreate/helper/bc"
	"heroku-backend-a-cocreate/model"
	"heroku-backend-a-cocreate/repository"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.RegisterDTO) model.User
	IsDuplicateEmail(email string) bool
}

func VerifyCredential(email, password string) interface{} {
	res := repository.VerifyCredential(email, password)
	if val, ok := res.(model.User); ok {
		comparePass := bc.ComparePass(val.Password, password)
		if val.Email == email && comparePass {
			return res
		}
		return false
	}
	return false
}

func CreateUser(user dto.RegisterDTO) model.Enrollment {
	enrollmentUser := model.Enrollment{
		NamaLengkap:   user.Name,
		Email:         user.Email,
		TopikDiminati: user.TopikDiminati,
	}

	userToCreate := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleID:   uint(user.RoleID),
	}
	res := repository.InsertUser(userToCreate, enrollmentUser)
	return res
}

func IsDuplicateEmail(email string) bool {
	err := repository.IsDuplicateEmail(email).Error
	return err == nil
}
