package scrapper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"go-scrapper/domain/dto"
)

func QueryCurrentBookmarks(bookmarkerUsername string) []dto.InsertAo3Data {
	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	var datas []dto.InsertAo3Data

	// On every a element which has href attribute call callback
	c.OnHTML("ol.bookmark.index.group", func(e *colly.HTMLElement) {
		e.ForEach("li[role='article']", func(_ int, el *colly.HTMLElement) {
			var data dto.InsertAo3Data
			el.ForEach("div.header.module", func(ind int, ele *colly.HTMLElement) {
				dateString := ele.ChildText("p.datetime")
				parsedTime, err := time.Parse("2 Jan 2006", dateString)
				if err != nil {
					panic(err)
				}
				data.WorkUpdateDate = parsedTime
				fmt.Println("Updated at: ", parsedTime.String())
				ele.ForEach("h4.heading", func(_ int, h *colly.HTMLElement) {
					worksUrl := h.ChildAttr("a[href^='/works/']", "href")
					usersUrl := h.ChildAttr("a[rel='author']", "href")
					workSplit := strings.Split(worksUrl, "/")
					userSplit := strings.Split(usersUrl, "/")
					workID := workSplit[2]
					workIDStr, err := strconv.Atoi(workID)
					if err != nil {
						panic(err)
					}
					var authorUsername string
					if len(strings.Split(usersUrl, "/")) == 1 {
						authorUsername = "anonymous"
					} else {
						authorUsername = userSplit[2]
					}
					// fmt.Println(workSplit, len(workSplit))
					// fmt.Println(userSplit, len(userSplit))
					workTitle := h.ChildText("a[href^='/works/']")
					data.WorkID = workIDStr
					data.AuthorUsername = authorUsername
					data.BookmarkerUsername = bookmarkerUsername
					data.WorkTitle = workTitle
					fmt.Println(workID)
					fmt.Println(authorUsername)
					fmt.Println(workTitle)
				})
			})
			el.ForEach("div.user.module.group", func(_ int, userModuleEl *colly.HTMLElement) {
				dateString := userModuleEl.ChildText("p.datetime")
				parsedTime, err := time.Parse("2 Jan 2006", dateString)
				if err != nil {
					panic(err)
				}
				fmt.Println("Bookmark at: ", parsedTime.String())
				data.BookmarkCreatedAt = parsedTime
				bookmark := userModuleEl.ChildText("blockquote>p")
				bookmarkInt, err := strconv.Atoi(bookmark)
				if err != nil {
					panic(err)
				}
				data.BookmarkChapter = bookmarkInt
			})
			fmt.Println("-------------------")
			datas = append(datas, data)
		})
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://archiveofourown.org/bookmarks?commit=Sort+and+Filter&bookmark_search%5Bsort_column%5D=bookmarkable_date&user_id=" + bookmarkerUsername)

	return datas
}
