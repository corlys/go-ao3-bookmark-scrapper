package services

import (
	"fmt"
	"go-scrapper/domain/dao"
	"go-scrapper/repositories"
	"time"
)

type BookmarkingService interface {
	CheckBookmark(username string, workID int) (dao.Bookmarks, error)
	CreateBookmark(authorUsername string, bookmarkerUsername string, workTitle string, workUpdatedAt time.Time, bookmarkChapter int, bookmarkCreatedAt time.Time)
}

type BookmarkingServiceImpl struct {
	UserRepo      repositories.UserRepository
	WorksRepo     repositories.WorksRepository
	BookmarksRepo repositories.BookmarkRepository
}

func (impl BookmarkingServiceImpl) CheckBookmark(username string, workID int) (dao.Bookmarks, error) {
	bookmark, err := impl.BookmarksRepo.GetBookmarkByUsernameAndWorkId(username, workID)
	if err != nil {
		return dao.Bookmarks{}, err
	}
	return bookmark, nil
}

func (impl BookmarkingServiceImpl) CreateBookmark(authorUsername string, bookmarkerUsername string, workTitle string, workUpdatedAt time.Time, bookmarkChapter int, bookmarkCreatedAt time.Time) {
	// 1. Create author & bookmarker if not exists, if exists get it
	// 2. Get or Create work
	// 3. Get or create bookmark
	// 4. Update Bookmark

	author, err := impl.UserRepo.GetOrInsert(authorUsername)
	if err != nil {
		panic(err)
	}
	bookmarker, err := impl.UserRepo.GetOrInsert(bookmarkerUsername)
	if err != nil {
		panic(err)
	}

	work, err := impl.WorksRepo.GetOrInsertWork(workTitle, workUpdatedAt, author)
	if err != nil {
		panic(err)
	}

	bookmark, err := impl.BookmarksRepo.GetOrInsertBookmark(bookmarkChapter, bookmarkCreatedAt, int(work.ID), bookmarker.Username)
	if err != nil {
		panic(err)
	}

	updatedBookmark, err := impl.BookmarksRepo.UpdateBookmarkByID(int(bookmark.ID), bookmarkChapter, bookmarkCreatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Println("successfully created bookmark with id: ", updatedBookmark.ID)
}

func BookmarkingServiceInit(bookmarksRepo repositories.BookmarkRepository, worksRepo repositories.WorksRepository, usersRepo repositories.UserRepository) *BookmarkingServiceImpl {
	return &BookmarkingServiceImpl{
		WorksRepo:     worksRepo,
		BookmarksRepo: bookmarksRepo,
		UserRepo:      usersRepo,
	}
}
