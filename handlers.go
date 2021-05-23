package main

import (
	"log"
	"strconv"

	linkValidation "github.com/jerrydevin96/webscraper/link-validation"
	"github.com/jerrydevin96/webscraper/scraper"
)

type WebPageDetails struct {
	HTMLVersion       string `json:"htmlVersion"`
	PageTitle         string `json:"pageTitle"`
	H1Length          string `json:"h1Length"`
	H2Length          string `json:"h2Length"`
	H3Length          string `json:"h3Length"`
	H4Length          string `json:"h4Length"`
	H5Length          string `json:"h5Length"`
	H6Length          string `json:"h6Length"`
	InternalLinks     string `json:"internalLinks"`
	ExternalLinks     string `json:"externalLinks"`
	InAccessibleLinks string `json:"inaccessibleLinks"`
}

func fetchPageDetails(URL string) {
	headings := scraper.FetchWebPageDetails(URL)
	log.Println("h1 length is : " + strconv.Itoa(len(headings.H1)))
	log.Println("h2 length is : " + strconv.Itoa(len(headings.H2)))
	log.Println("h3 length is : " + strconv.Itoa(len(headings.H3)))
	log.Println("h4 length is : " + strconv.Itoa(len(headings.H4)))
	log.Println("h5 length is : " + strconv.Itoa(len(headings.H5)))
	log.Println("h6 length is : " + strconv.Itoa(len(headings.H6)))
	log.Println("links length is : " + strconv.Itoa(len(headings.Links)))
	linkValidation.ClassifyLinks(headings.Links, URL)
}
