package main

import (
	"encoding/json"

	linkValidation "github.com/jerrydevin96/webscraper/link-validation"
	"github.com/jerrydevin96/webscraper/scraper"
)

type WebPageDetails struct {
	HTMLVersion             string         `json:"htmlVersion"`
	PageTitle               string         `json:"pageTitle"`
	H1Length                int            `json:"h1Length"`
	H2Length                int            `json:"h2Length"`
	H3Length                int            `json:"h3Length"`
	H4Length                int            `json:"h4Length"`
	H5Length                int            `json:"h5Length"`
	H6Length                int            `json:"h6Length"`
	InternalLinksLength     int            `json:"internalLinks"`
	ExternalLinksLength     int            `json:"externalLinks"`
	InAccessibleLinksLength int            `json:"inaccessibleLinks"`
	AdditionalInfo          AdditionalData `json:"additionalInfo"`
}

type AdditionalData struct {
	InternalLinks     []string `json:"internalLinks"`
	ExternalLinks     []string `json:"externalLinks"`
	InAccessibleLinks []string `json:"inAccessibleLinks"`
}

func pageDetailsHandler(requestData string) string {
	response := ""
	var additionalInfo AdditionalData
	reqDataJSON := make(map[string]string)
	json.Unmarshal([]byte(requestData), &reqDataJSON)
	pageDetails := &WebPageDetails{}
	scraperDetails := scraper.FetchWebPageDetails(reqDataJSON["url"])
	validatedLinks := linkValidation.ValidateLinks(scraperDetails.Links, reqDataJSON["url"])
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
	additionalInfo.InternalLinks = validatedLinks.InternalLinks
	additionalInfo.ExternalLinks = validatedLinks.ExternalLinks
	additionalInfo.InAccessibleLinks = validatedLinks.InAccessibleLinks
	pageDetails.AdditionalInfo = additionalInfo
	jsonBytes, _ := json.Marshal(pageDetails)
	response = string(jsonBytes)
	return response
}
