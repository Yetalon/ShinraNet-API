package models

type Post struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	UserID   int    `json:"user_id" gorm:"not null;index"`
	Title    string `json:"title"`
	PostText string `json:"post_text"`
	Likes    int    `json:"likes" gorm:"default:0"`
}
