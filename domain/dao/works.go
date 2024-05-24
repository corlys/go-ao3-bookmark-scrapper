package dao

import (
	"gorm.io/gorm"
	"time"
)

type (
	Works struct {
		gorm.Model
		Title           string      `json:"title"`
		WorkUpdatedDate time.Time   `json:"work_updated_date"`
		Bookmarks       []Bookmarks `gorm:"foreignKey:WorkID;references:ID" json:"bookmarks"`
		AuthorUsername  string      `json:"author_user_name"`
	}
)
