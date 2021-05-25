package scraper

import (
	"log"

	"github.com/gocolly/colly"
)

type WebPageDetails struct {
	Title   string
	Version string
	H1      []string
	H2      []string
	H3      []string
	H4      []string
	H5      []string
	H6      []string
	Links   []string
}

func FetchWebPageDetails(URL string) *WebPageDetails {
	log.Println("Fetching all relevant details from webpage")
	response := &WebPageDetails{}
	versionChannel := make(chan string)
	c := colly.NewCollector()
	registerCollyHandlers(c, response)
	go getHTMLVersion(URL, versionChannel)
	c.Visit(URL)
	version := <-versionChannel
	response.Version = version
	log.Println(response)
	return response
}

func getHTMLVersion(URL string, versionChannel chan string) {
	version, err := getHTMLVersionHandler(URL)
	if err != nil {
		errString := err.Error()
		versionChannel <- errString
		return
	}
	versionChannel <- version
}
