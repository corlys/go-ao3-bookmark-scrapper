package repositories

import (
	"go-scrapper/domain/dao"
	"time"

	"gorm.io/gorm"
)

type BookmarkRepository interface {
	InsertBookmark(chapter int, bookmarkUpdateDate time.Time, workID int, bookmarkerUsername string) (dao.Bookmarks, error)
	GetOrInsertBookmark(chapter int, bookmarkUpdateDate time.Time, workID int, bookmarkerUsername string) (dao.Bookmarks, error)
	GetBookmarkByUsernameAndWorkId(username string, workID int) (dao.Bookmarks, error)
	UpdateBookmarkByID(bookmarkID int, newBookmarkChapter int, newBookmarkUpdateDate time.Time) (dao.Bookmarks, error)
}

type BookmarkRepositoryImpl struct {
	db *gorm.DB
}

func (repo BookmarkRepositoryImpl) InsertBookmark(chapter int, bookmarkCreatedAt time.Time, workID int, bookmarkerUsername string) (dao.Bookmarks, error) {

	bookmark := dao.Bookmarks{
		Chapter:            chapter,
		BookmarkCreatedAt:  bookmarkCreatedAt,
		BookmarkerUsername: bookmarkerUsername,
		WorkID:             workID,
	}

	err := repo.db.Create(&bookmark).Error
	if err != nil {
		return dao.Bookmarks{}, err
	}
	return bookmark, nil
}

func (repo BookmarkRepositoryImpl) GetOrInsertBookmark(chapter int, bookmarkUpdateDate time.Time, workID int, bookmarkerUsername string) (dao.Bookmarks, error) {
	bookmark, err := repo.GetBookmarkByUsernameAndWorkId(bookmarkerUsername, workID)
	if err == gorm.ErrRecordNotFound {
		newBookmark, errInsert := repo.InsertBookmark(chapter, bookmarkUpdateDate, workID, bookmarkerUsername)
		if errInsert != nil {
			return dao.Bookmarks{}, errInsert
		}
		return newBookmark, nil
	}
	return bookmark, nil
}

func (repo BookmarkRepositoryImpl) GetBookmarkByUsernameAndWorkId(username string, workID int) (dao.Bookmarks, error) {
	var bookmark dao.Bookmarks
	err := repo.db.Where("bookmarker_username = ? AND work_id = ?", username, workID).First(&bookmark).Error
	if err != nil {
		return dao.Bookmarks{}, err
	}
	return bookmark, nil
}

func (repo BookmarkRepositoryImpl) UpdateBookmarkByID(bookmarkID int, newBookmarkChapter int, newBookmarkCreatedAt time.Time) (dao.Bookmarks, error) {

	var bookmark dao.Bookmarks
	err := repo.db.First(&bookmark, bookmarkID).Error
	if err != nil {
		return dao.Bookmarks{}, err
	}
	bookmark.BookmarkCreatedAt = newBookmarkCreatedAt
	bookmark.Chapter = newBookmarkChapter
	errInsert := repo.db.Save(&bookmark).Error
	if errInsert != nil {
		return dao.Bookmarks{}, errInsert
	}
	return bookmark, nil

}

func BookmarkRepositoryInit(db *gorm.DB) *BookmarkRepositoryImpl {
	return &BookmarkRepositoryImpl{
		db: db,
	}
}
