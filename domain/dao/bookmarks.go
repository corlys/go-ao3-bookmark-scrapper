package dao

import (
	"gorm.io/gorm"
	"time"
)

type (
	Bookmarks struct {
		gorm.Model
		Chapter            int       `json:"chapter"`
		BookmarkCreatedAt  time.Time `json:"bookmark_created_at"`
		WorkID             int       `json:"work_id"`
		BookmarkerUsername string    `json:"bookmarker_username"`
	}
)
