package service

import (
	"Final_Project/model"
)

type UserService interface {
	UserRegister(user model.User) (res model.User, err error)
	UserLogin(user model.User) (res model.User, err error)
}

// User Endpoint
func (s *Service) UserRegister(user model.User) (res model.User, err error) {
	res, err = s.repo.UserRegister(user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UserLogin(user model.User) (res model.User, err error) {
	res, err = s.repo.UserLogin(user)
	if err != nil {
		return res, err
	}
	return res, nil
}
