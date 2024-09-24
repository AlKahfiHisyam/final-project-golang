package models

import "time"

type GormModel struct {
	ID       uint       `grom:"primaryKey" json:"id"`
	CreateAt *time.Time `json:"created_at,omitempty"`
	UpdateAt *time.Time `json:"updated_at,omitempty"`
}
