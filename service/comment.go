package service

import (
	"Final_Project/model"
)

type CommentService interface {
	CommentGetAll(comment model.Comment) (res []model.Comment, err error)
	CommentGet(comment model.Comment) (res model.Comment, err error)
	CommentCreate(comment model.Comment) (res model.Comment, err error)
	CommentUpdate(comment model.Comment) (res model.Comment, err error)
	CommentDelete(comment model.Comment) (err error)
	CommentAuthorization(comment model.Comment) (res model.Comment, err error)
}

func (s *Service) CommentGetAll(comment model.Comment) (res []model.Comment, err error) {
	res, err = s.repo.CommentGetAll(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentGet(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentGet(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentCreate(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentCreate(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentUpdate(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentUpdate(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) CommentDelete(comment model.Comment) (err error) {
	err = s.repo.CommentDelete(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CommentAuthorization(comment model.Comment) (res model.Comment, err error) {
	res, err = s.repo.CommentAuthorization(comment)
	if err != nil {
		return res, err
	}
	return res, nil
}
