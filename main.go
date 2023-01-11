package main

import (
	"fmt"
	"time"
	"github.com/gocolly/colly"
)

type DownloadImage struct {
	Name		string
	Image 	string
}

func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	// download_image = make([]DownloadImage, 0)

	c.OnHTML("a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")

}
