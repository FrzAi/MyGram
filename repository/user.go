package repository

import (
	"Final_Project/model"
)

type UserRepo interface {
	UserRegister(model.User) (res model.User, err error)
	UserLogin(model.User) (res model.User, err error)
}

func (r Repo) UserRegister(user model.User) (res model.User, err error) {

	err = r.db.Debug().Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r Repo) UserLogin(user model.User) (res model.User, err error) {

	err = r.db.Debug().Where("username = ?", user.Username).Take(&res).Error
	if err != nil {
		err = r.db.Debug().Where("email = ?", user.Email).Take(&res).Error
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
