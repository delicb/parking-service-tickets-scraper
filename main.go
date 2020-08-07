package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	col := colly.NewCollector()
	col.OnHTML("#post-44447 > div > h2:nth-child(1)", func(el *colly.HTMLElement) {
		fmt.Println(el.Text)
	})

	col.OnError(func(res *colly.Response, err error) {
		if err != nil {
			log.Printf("error scraping: %v\n", err)
			log.Printf("Response status code: %d\n", res.StatusCode)
		}
	})

	err := col.Visit(fmt.Sprintf("https://parking-servis.co.rs/edpk/?fine=%s", os.Args[1]))
	if err != nil {
		panic(err)
	}
}
