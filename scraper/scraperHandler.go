package scraper

import (
	"errors"
	"log"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/jerrydevin96/webscraper/pageFetcher"
)

func registerCollyHandlers(c *colly.Collector, details *WebPageDetails) {
	c.OnHTML("head title", func(h *colly.HTMLElement) {
		details.Title = h.Text
	})
	c.OnHTML("body h1", func(h *colly.HTMLElement) {
		header := h.Text
		details.H1 = append(details.H1, header)
	})
	c.OnHTML("body h2", func(h *colly.HTMLElement) {
		header := h.Text
		details.H2 = append(details.H2, header)
	})
	c.OnHTML("body h3", func(h *colly.HTMLElement) {
		header := h.Text
		details.H3 = append(details.H3, header)
	})
	c.OnHTML("body h4", func(h *colly.HTMLElement) {
		header := h.Text
		details.H4 = append(details.H4, header)
	})
	c.OnHTML("body h5", func(h *colly.HTMLElement) {
		header := h.Text
		details.H5 = append(details.H5, header)
	})
	c.OnHTML("body h6", func(h *colly.HTMLElement) {
		header := h.Text
		details.H6 = append(details.H6, header)
	})
	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		details.Links = append(details.Links, link)
	})
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
			versions := regexMatcher(htmlHeader, "(HTML|XHTML) ([1-9].([0-9][0-9]|[1-9]))")
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
