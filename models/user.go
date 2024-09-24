package models

import (
	"final-project-golang/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username        string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your Username is required"`
	Email           string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password        string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age             int           `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required"`
	ProfileImageUrl string        `json:"profile_image_url" form:"profile_image_url"`
	Photos          []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Comments        []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	SocialMedias    []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

type CreateUserRequest struct {
	Age      int    `json:"age" form:"age" binding:"required,min=9"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
	Username string `json:"username" form:"username" binding:"required"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Username string `json:"username" form:"username" binding:"required"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=6"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	ID        uint       `json:"id"`
	Age       int        `json:"age"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserPhotoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCommentResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserSocialMediaResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

// Validasi sebelum create User
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
