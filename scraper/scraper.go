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
	c.OnHTML("head title", getTitleCollyHTMLCallBack)
	c.OnHTML("body h1", getH1CollyHTMLCallBack)
	c.OnHTML("body h2", getH2CollyHTMLCallBack)
	c.OnHTML("body h3", getH3CollyHTMLCallBack)
	c.OnHTML("body h4", getH4CollyHTMLCallBack)
	c.OnHTML("body h5", getH5CollyHTMLCallBack)
	c.OnHTML("body h6", getH6CollyHTMLCallBack)
	c.OnHTML("a[href]", getLinksCollyHTMLCallBack)
	go getHTMLVersion(URL, versionChannel)
	c.Visit(URL)
	version := <-versionChannel
	response.Title = Title
	response.H1 = H1List
	response.H2 = H2List
	response.H3 = H3List
	response.H4 = H4List
	response.H5 = H5List
	response.H6 = H6List
	response.Links = Links
	response.Version = version
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
