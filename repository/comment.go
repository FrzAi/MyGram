package repository

import (
	"Final_Project/model"
	"time"
)

type CommentRepo interface {
	CommentGetAll(comment model.Comment) (res []model.Comment, err error)
	CommentGet(model.Comment) (res model.Comment, err error)
	CommentCreate(model.Comment) (res model.Comment, err error)
	CommentUpdate(model.Comment) (res model.Comment, err error)
	CommentDelete(model.Comment) (err error)
	CommentAuthorization(model.Comment) (res model.Comment, err error)
}

func (r Repo) CommentGetAll(comment model.Comment) (res []model.Comment, err error) {
	err = r.db.Where("photo_id <> ?", comment.PhotoID).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentGet(comment model.Comment) (res model.Comment, err error) {
	err = r.db.Model(&model.Comment{}).First(&res, comment.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentCreate(comment model.Comment) (res model.Comment, err error) {
	err = r.db.Create(&comment).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentUpdate(comment model.Comment) (res model.Comment, err error) {
	err = r.db.First(&res, comment.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", comment.ID).Updates(map[string]interface{}{"message": comment.Message, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CommentDelete(comment model.Comment) (err error) {
	err = r.db.First(&comment, comment.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&comment, comment.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) CommentAuthorization(comment model.Comment) (res model.Comment, err error) {
	err = r.db.Select("user_id").First(&res, comment.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
