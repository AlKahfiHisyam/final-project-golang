package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~SocialMediaUrl is required"`
	UserId         uint   `json:"user_id"`
	User           *User
}

type CreateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" binding:"required"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" form:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url" binding:"required"`
}

type SocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CreatedAt      *time.Time `json:"created_at"`
	User           *UserSocialMediaResponse
}

type CreateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
}

type UpdateSocialMediaResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	SocialMediaUrl string     `json:"social_media_url"`
	UserId         uint       `json:"user_id"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

// Validasi sebelum create social media
func (u *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
