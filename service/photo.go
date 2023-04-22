package service

import (
	"Final_Project/model"
)

type PhotoService interface {
	PhotoGetAll() (res []model.Photo, err error)
	PhotoGet(photo model.Photo) (res model.Photo, err error)
	PhotoCreate(photo model.Photo) (res model.Photo, err error)
	PhotoUpdate(photo model.Photo) (res model.Photo, err error)
	PhotoDelete(photo model.Photo) (err error)
	PhotoAuthorization(photo model.Photo) (res model.Photo, err error)
}

func (s *Service) PhotoGetAll() (res []model.Photo, err error) {
	res, err = s.repo.PhotoGetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoGet(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoGet(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoCreate(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoCreate(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoUpdate(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoUpdate(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) PhotoDelete(photo model.Photo) (err error) {
	err = s.repo.PhotoDelete(photo)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PhotoAuthorization(photo model.Photo) (res model.Photo, err error) {
	res, err = s.repo.PhotoAuthorization(photo)
	if err != nil {
		return res, err
	}
	return res, nil
}
