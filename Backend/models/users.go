package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null;unique"`
	ImageURL string `json:"image_url"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}
