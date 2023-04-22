package repository

import (
	"Final_Project/model"
	"time"
)

type SosialMediaRepo interface {
	SocialMediaGetAll() (res []model.SocialMedia, err error)
	SocialMediaGet(model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaCreate(model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaUpdate(model.SocialMedia) (res model.SocialMedia, err error)
	SocialMediaDelete(model.SocialMedia) (err error)
	SocialMediaAuthorization(model.SocialMedia) (res model.SocialMedia, err error)
}

func (r Repo) SocialMediaGetAll() (res []model.SocialMedia, err error) {
	err = r.db.Model(&model.SocialMedia{}).Preload("User").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaGet(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.Model(&model.SocialMedia{}).Preload("User").First(&res, socialmedia.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaCreate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.Create(&socialmedia).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaUpdate(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.First(&res, socialmedia.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", socialmedia.ID).Updates(map[string]interface{}{"name": socialmedia.Name, "social_media_url": socialmedia.SocialMediaUrl, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) SocialMediaDelete(socialmedia model.SocialMedia) (err error) {
	err = r.db.First(&socialmedia, socialmedia.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&socialmedia, socialmedia.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) SocialMediaAuthorization(socialmedia model.SocialMedia) (res model.SocialMedia, err error) {
	err = r.db.Select("user_id").First(&res, socialmedia.ID).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
