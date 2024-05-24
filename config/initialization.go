package config

import (
	"go-scrapper/repositories"
	"go-scrapper/services"
)

type Initialization struct {
	BookmarksRepo      repositories.BookmarkRepository
	UsersRepo          repositories.UserRepository
	Worksrepo          repositories.WorksRepository
	BookmarkingService services.BookmarkingService
}

func NewInitialization(
	bookmarkRepo repositories.BookmarkRepository,
	usersRepo repositories.UserRepository,
	worksRepo repositories.WorksRepository,
	bookmarkingService services.BookmarkingService,
) *Initialization {
	return &Initialization{
		BookmarksRepo:      bookmarkRepo,
		UsersRepo:          usersRepo,
		Worksrepo:          worksRepo,
		BookmarkingService: bookmarkingService,
	}
}
