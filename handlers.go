package main

import (
	"encoding/json"
	"log"
	"strconv"

	linkValidation "github.com/jerrydevin96/webscraper/link-validation"
	"github.com/jerrydevin96/webscraper/scraper"
)

type WebPageDetails struct {
	HTMLVersion             string `json:"htmlVersion"`
	PageTitle               string `json:"pageTitle"`
	H1Length                int    `json:"h1Length"`
	H2Length                int    `json:"h2Length"`
	H3Length                int    `json:"h3Length"`
	H4Length                int    `json:"h4Length"`
	H5Length                int    `json:"h5Length"`
	H6Length                int    `json:"h6Length"`
	InternalLinksLength     int    `json:"internalLinks"`
	ExternalLinksLength     int    `json:"externalLinks"`
	InAccessibleLinksLength int    `json:"inaccessibleLinks"`
}

func fetchPageDetails(URL string) {
	pageDetails := &WebPageDetails{}
	scraperDetails := scraper.FetchWebPageDetails(URL)
	log.Println("HTML version is : " + scraperDetails.Version)
	log.Println("page title is : " + scraperDetails.Title)
	log.Println("h1 length is : " + strconv.Itoa(len(scraperDetails.H1)))
	log.Println("h2 length is : " + strconv.Itoa(len(scraperDetails.H2)))
	log.Println("h3 length is : " + strconv.Itoa(len(scraperDetails.H3)))
	log.Println("h4 length is : " + strconv.Itoa(len(scraperDetails.H4)))
	log.Println("h5 length is : " + strconv.Itoa(len(scraperDetails.H5)))
	log.Println("h6 length is : " + strconv.Itoa(len(scraperDetails.H6)))
	log.Println("links length is : " + strconv.Itoa(len(scraperDetails.Links)))
	validatedLinks := linkValidation.ValidateLinks(scraperDetails.Links, URL)
	log.Println("internal links length : " + strconv.Itoa(len(validatedLinks.InternalLinks)))
	log.Println("external links length : " + strconv.Itoa(len(validatedLinks.ExternalLinks)))
	log.Println("in accessible links length : " + strconv.Itoa(len(validatedLinks.InAccessibleLinks)))
	log.Println("non accessible links : ")
	log.Println(validatedLinks.InAccessibleLinks)
	pageDetails.PageTitle = scraperDetails.Title
	pageDetails.HTMLVersion = scraperDetails.Version
	pageDetails.H1Length = len(scraperDetails.H1)
	pageDetails.H2Length = len(scraperDetails.H2)
	pageDetails.H3Length = len(scraperDetails.H3)
	pageDetails.H4Length = len(scraperDetails.H4)
	pageDetails.H5Length = len(scraperDetails.H5)
	pageDetails.H6Length = len(scraperDetails.H6)
	pageDetails.InternalLinksLength = len(validatedLinks.InternalLinks)
	pageDetails.ExternalLinksLength = len(validatedLinks.ExternalLinks)
	pageDetails.InAccessibleLinksLength = len(validatedLinks.InAccessibleLinks)

	jsonBytes, _ := json.Marshal(pageDetails)
	log.Println(string(jsonBytes))
}
