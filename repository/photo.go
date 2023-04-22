package repository

import (
	"Final_Project/model"
	"time"
)

type PhotoRepo interface {
	PhotoGetAll() (res []model.Photo, err error)
	PhotoGet(model.Photo) (res model.Photo, err error)
	PhotoCreate(model.Photo) (res model.Photo, err error)
	PhotoUpdate(model.Photo) (res model.Photo, err error)
	PhotoDelete(model.Photo) (err error)
	PhotoAuthorization(model.Photo) (res model.Photo, err error)
}

func (r Repo) PhotoGetAll() (res []model.Photo, err error) {
	err = r.db.Model(&model.Photo{}).Preload("Comments").Preload("User").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoGet(photo model.Photo) (res model.Photo, err error) {
	err = r.db.Model(&model.Photo{}).Preload("Comments").Preload("User").First(&res, photo.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoCreate(photo model.Photo) (res model.Photo, err error) {
	err = r.db.Create(&photo).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoUpdate(photo model.Photo) (res model.Photo, err error) {
	err = r.db.First(&res, photo.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", photo.ID).Updates(map[string]interface{}{"title": photo.Title, "caption": photo.Caption, "photo_url": photo.PhotoUrl, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) PhotoDelete(photo model.Photo) (err error) {
	err = r.db.First(&photo, photo.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&photo, photo.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) PhotoAuthorization(photo model.Photo) (res model.Photo, err error) {
	err = r.db.Select("user_id").First(&res, photo.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
