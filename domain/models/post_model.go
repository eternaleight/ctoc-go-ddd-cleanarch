package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;column:createdAt" json:"createdAt"`
	AuthorID  uint      `gorm:"column:authorId;index" json:"authorId"`
}

func (Post) TableName() string {
	return "Post"
}
