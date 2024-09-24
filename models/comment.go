package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserId  uint `json:"user_id"`
	User    *User
	PhotoId uint `json:"photo_id"`
	Photo   *Photo
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
}

type CreateCommentRequest struct {
	PhotoId uint   `json:"photo_id" form:"photo_id"`
	Message string `json:"message" form:"message" binding:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" form:"message" binding:"required"`
}

type CommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserId    uint       `json:"user_id"`
	PhotoId   uint       `json:"photo_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *UserCommentResponse
	Photo     *PhotoCommentResponse
}

type CreateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserId    uint       `json:"user_id"`
	PhotoId   uint       `json:"photo_id"`
	CreatedAt *time.Time `json:"updated_at"`
}

type UpdateCommentResponse struct {
	ID        uint       `json:"id"`
	Message   string     `json:"message"`
	UserId    uint       `json:"user_id"`
	PhotoId   uint       `json:"photo_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// Validasi sebelum create comment
func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
