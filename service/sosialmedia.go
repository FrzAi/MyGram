package service

import (
	"Final_Project/model"
)

type SocialMediaService interface {
	SocialMediaGetAll() (res []model.SocialMedia, err error)
	SocialMediaGet(socialmedia model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaCreate(socialmedia model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaUpdate(socialmedia model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaDelete(socialmedia model.SocialMedia) (err error)
	SocialMediaAuthorization(model.SocialMedia) (res model.SocialMedia, err error)
}

func (s *Service) SocialMediaGetAll() (res []model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaGetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaGet(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaGet(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaCreate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaCreate(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaUpdate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaUpdate(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) SocialMediaDelete(socialmedia model.SocialMedia) (err error) {
	err = s.repo.SocialMediaDelete(socialmedia)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SocialMediaAuthorization(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	res, err = s.repo.SocialMediaAuthorization(socialmedia)
	if err != nil {
		return res, err
	}
	return res, nil
}
