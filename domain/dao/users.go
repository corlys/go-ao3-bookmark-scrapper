package dao

import "gorm.io/gorm"

type (
	Users struct {
		gorm.Model
		Username  string      `gorm:"uniqueIndex" json:"username"`
		Works     []Works     `gorm:"foreignKey:AuthorUsername;references:Username" json:"works"`
		Bookmarks []Bookmarks `gorm:"foreignKey:BookmarkerUsername;references:Username" json:"bookmarks"`
	}
)
