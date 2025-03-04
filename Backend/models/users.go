package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	ImageURL string `json:"image_url"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}
