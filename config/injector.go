//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"go-scrapper/repositories"
	"go-scrapper/services"
)

var db = wire.NewSet(ConnectToDB)

var bookmarksRepositorySet = wire.NewSet(repositories.BookmarkRepositoryInit,
	wire.Bind(new(repositories.BookmarkRepository), new(*repositories.BookmarkRepositoryImpl)),
)

var usersRepositorySet = wire.NewSet(repositories.UserRepositoryInit,
	wire.Bind(new(repositories.UserRepository), new(*repositories.UserRepositoryImpl)),
)

var worksRepositorySet = wire.NewSet(repositories.WorksRepositoryInit,
	wire.Bind(new(repositories.WorksRepository), new(*repositories.WorksRepositoryImpl)),
)

var bookmarkServiceSet = wire.NewSet(services.BookmarkingServiceInit,
	wire.Bind(new(services.BookmarkingService), new(*services.BookmarkingServiceImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, bookmarksRepositorySet, usersRepositorySet, worksRepositorySet, bookmarkServiceSet)
	return nil
}
