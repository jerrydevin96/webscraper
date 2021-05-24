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
	Title = h.Text
}

func getH1CollyHTMLCallBack(h *colly.HTMLElement) {
	header := h.Text
	H1List = append(H1List, header)
}

func getH2CollyHTMLCallBack(h *colly.HTMLElement) {
	header := h.Text
	H2List = append(H2List, header)
}

func getH3CollyHTMLCallBack(h *colly.HTMLElement) {
	header := h.Text
	H3List = append(H3List, header)
}

func getH4CollyHTMLCallBack(h *colly.HTMLElement) {
	header := h.Text
	H4List = append(H4List, header)
}

func getH5CollyHTMLCallBack(h *colly.HTMLElement) {
	header := h.Text
	H5List = append(H5List, header)
}

func getH6CollyHTMLCallBack(h *colly.HTMLElement) {
	header := h.Text
	H6List = append(H6List, header)
}

func getLinksCollyHTMLCallBack(h *colly.HTMLElement) {
	link := h.Attr("href")
	Links = append(Links, link)
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
