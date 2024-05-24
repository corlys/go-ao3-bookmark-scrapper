package dto

import "time"

type InsertAo3Data struct {
	WorkID             int
	WorkTitle          string
	AuthorUsername     string
	BookmarkerUsername string
	BookmarkChapter    int
	WorkUpdateDate     time.Time
	BookmarkCreatedAt  time.Time
}
