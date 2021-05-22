package main

import (
	"log"
	"strconv"

	"github.com/jerrydevin96/webscraper/scraper"
)

func main() {
	log.Println("starting scraper")
	title := scraper.FetchWebPageTitle("https://medium.com/")
	log.Println("page title is : " + title)

	headings := scraper.FetchWebPageHeadings("https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery/")
	log.Println("h1 length is : " + strconv.Itoa(len(headings.H1)))
	log.Println("h2 length is : " + strconv.Itoa(len(headings.H2)))
	log.Println("h3 length is : " + strconv.Itoa(len(headings.H3)))
	log.Println("h4 length is : " + strconv.Itoa(len(headings.H4)))
	log.Println("h5 length is : " + strconv.Itoa(len(headings.H5)))
	log.Println("h6 length is : " + strconv.Itoa(len(headings.H6)))
}
