package scraper

import (
	"errors"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	pageFetcher "github.com/jerrydevin96/webscraper/page-fetcher"
)

var (
	Title  = ""
	H1List = make([]string, 0)
	H2List = make([]string, 0)
	H3List = make([]string, 0)
	H4List = make([]string, 0)
	H5List = make([]string, 0)
	H6List = make([]string, 0)
	Links  = make([]string, 0)
)

func getTitleCollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get title Colly HTML call back")
	Title = h.Text
	log.Println("Title is " + Title)
	log.Println("Get title Colly HTML call back completed")
}

func getH1CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H1 Colly HTML call back")
	header := h.Text
	H1List = append(H1List, header)
	log.Println("Get H1 Colly HTML call back completed")
}

func getH2CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H2 Colly HTML call back")
	header := h.Text
	H2List = append(H2List, header)
	log.Println("Get H2 Colly HTML call back completed")
}

func getH3CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H3 Colly HTML call back")
	header := h.Text
	H3List = append(H3List, header)
	log.Println("Get H3 Colly HTML call back completed")
}

func getH4CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H4 Colly HTML call back")
	header := h.Text
	H4List = append(H4List, header)
	log.Println("Get H4 Colly HTML call back completed")
}

func getH5CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H1 Colly HTML call back")
	header := h.Text
	H5List = append(H5List, header)
	log.Println("Get H5 Colly HTML call back completed")
}

func getH6CollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get H1 Colly HTML call back")
	header := h.Text
	H6List = append(H6List, header)
	log.Println("Get H1 Colly HTML call back completed")
}

func getLinksCollyHTMLCallBack(h *colly.HTMLElement) {
	log.Println("Executing get links Colly HTML call back")
	link := h.Attr("href")
	Links = append(Links, link)
	log.Println("Get Links Colly HTML call back completed")
}

func getHTMLVersionHandler(URL string) (string, error) {
	version := ""
	var err error
	pageData, err := pageFetcher.FetchHTMLPage(URL)
	if err != nil {
		return "", err
	}
	version, err = processHTMLForVersion(pageData)
	return version, err
}

func regexMatcher(compareString string, pattern string) []string {
	var matches []string
	re := regexp.MustCompile(pattern)
	matches = re.FindAllString(compareString, -1)
	return matches
}

func processHTMLForVersion(HTML string) (string, error) {
	var version string
	var err error
	matches := regexMatcher(HTML, `<!(DOCTYPE|doctype) ([A-Z a-z"-/0-9\n:]*)>`)
	if matches == nil {
		log.Println("No matches.")
		return "", errors.New("no matches found for html version")
	} else {
		htmlHeader := matches[0]
		if htmlHeader == `<!DOCTYPE html>` || htmlHeader == `<!doctype html>` {
			version = "HTML 5"
		} else {
			versions := regexMatcher(htmlHeader, "(HTML|XHTML) ([1-9].[0-9][0-9]|[1-9])")
			if versions == nil {
				log.Println("No matches.")
				return "", errors.New("no matches found for html version")
			} else {
				version = versions[0]
			}
		}
	}
	return version, err
}
