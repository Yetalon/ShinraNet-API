package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   int    `json:"user_id" gorm:"not null;index"`
	Title    string `json:"title"`
	PostText string `json:"post_text"`
	Likes    int    `json:"likes" gorm:"default:0"`
}
