package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

const ticketSelector = "#post-44447 > div > h2:nth-child(1)"
const URLTemplate = "https://parking-servis.co.rs/edpk/?fine=%s"

func main() {

	col := colly.NewCollector()
	col.OnHTML(ticketSelector, func(el *colly.HTMLElement) {
		fmt.Println(el.Text)
	})

	col.OnError(func(res *colly.Response, err error) {
		if err != nil {
			log.Printf("error scraping: %v\n", err)
			log.Printf("Response status code: %d\n", res.StatusCode)
		}
	})

	err := col.Visit(fmt.Sprintf(URLTemplate, os.Args[1]))
	if err != nil {
		panic(err)
	}
}
