package main

import (
	"fmt"
	"go-scrapper/config"
	"go-scrapper/scrapper"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	init := config.Init()
	fmt.Println("Successfully Initializing APP", init)

	datas := scrapper.QueryCurrentBookmarks("liuhei")

	for _, data := range datas {
		init.BookmarkingService.CreateBookmark(
			data.AuthorUsername,
			data.BookmarkerUsername,
			data.WorkTitle,
			data.WorkUpdateDate,
			data.BookmarkChapter,
			data.BookmarkCreatedAt,
		)
	}
}
