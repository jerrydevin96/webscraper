package scraper

import (
	"log"

	"github.com/gocolly/colly"
)

type Titles struct {
	H1 []string
	H2 []string
	H3 []string
	H4 []string
	H5 []string
	H6 []string
}

func FetchWebPageTitle(URL string) string {
	log.Println("Fetching Web Page Title")
	title := ""
	c := colly.NewCollector()
	c.OnHTML("head title", getTitleCollyHTMLCallBack)
	c.Visit(URL)
	title = Title
	return title
}

func FetchWebPageHeadings(URL string) *Titles {
	log.Println("Fetching headings from page")
	response := &Titles{}
	c := colly.NewCollector()
	c.OnHTML("body h1", getH1CollyHTMLCallBack)
	c.OnHTML("body h2", getH2CollyHTMLCallBack)
	c.OnHTML("body h3", getH3CollyHTMLCallBack)
	c.OnHTML("body h4", getH4CollyHTMLCallBack)
	c.OnHTML("body h5", getH5CollyHTMLCallBack)
	c.OnHTML("body h6", getH6CollyHTMLCallBack)
	c.Visit(URL)
	response.H1 = H1List
	response.H2 = H2List
	response.H3 = H3List
	response.H4 = H4List
	response.H5 = H5List
	response.H6 = H6List
	log.Println("headings are : ")
	log.Println(response)
	return response
}
