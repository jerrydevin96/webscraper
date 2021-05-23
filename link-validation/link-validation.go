package linkValidation

import (
	"log"
	"net/url"
	"strings"

	pageFetcher "github.com/jerrydevin96/webscraper/page-fetcher"
)

type ValidatedLinks struct {
	InternalLinks     []string
	ExternalLinks     []string
	InAccessibleLinks []string
}

func classifyLinks(links []string, URL string) ([]string, []string) {
	internalLinks := make([]string, 0)
	externalLinks := make([]string, 0)
	log.Println("URL is : " + URL)
	u, _ := url.Parse(URL)
	baseHost := u.Host
	baseScheme := u.Scheme
	baseURL := baseScheme + "://" + baseHost
	log.Println(baseHost)
	log.Println("printing links")
	for _, link := range links {
		u, _ = url.Parse(link)
		linkHost := u.Host
		if linkHost == "" {
			link = baseURL + link
			log.Println("appending internal link : " + link)
			internalLinks = append(internalLinks, link)
		} else if strings.Contains(linkHost, baseHost) {
			log.Println("appending internal link : " + link)
			internalLinks = append(internalLinks, link)
		} else {
			log.Println("appending external link : " + link)
			externalLinks = append(externalLinks, link)
		}
	}
	return internalLinks, externalLinks
}

func checkLinkAccessibility(links []string) []string {
	nonAccessibleLinks := make([]string, 0)
	c := make(chan pageFetcher.PageAccessibilityStatus)
	for _, link := range links {
		log.Println(link)
		go pageFetcher.CheckPageAccessibility(link, c)
	}

	for i := 0; i < len(links); i++ {
		log.Println(i + 1)
		accessibilityStatus := <-c
		log.Println(accessibilityStatus.URL)
		log.Println(accessibilityStatus.StatusCode)
		log.Println(accessibilityStatus.Error)
		url := accessibilityStatus.URL
		statusCode := accessibilityStatus.StatusCode
		err := accessibilityStatus.Error
		if (statusCode != 200 && statusCode != 201) || err != nil {
			nonAccessibleLinks = append(nonAccessibleLinks, url)
		}
	}
	return nonAccessibleLinks
}

func ValidateLinks(links []string, URL string) *ValidatedLinks {
	validatedLinks := &ValidatedLinks{}
	combinedLinks := make([]string, 0)
	internalLinks, externalLinks := classifyLinks(links, URL)
	combinedLinks = append(combinedLinks, internalLinks...)
	combinedLinks = append(combinedLinks, externalLinks...)
	nonAccessibleLinks := checkLinkAccessibility(combinedLinks)
	validatedLinks.InternalLinks = internalLinks
	validatedLinks.ExternalLinks = externalLinks
	validatedLinks.InAccessibleLinks = nonAccessibleLinks
	return validatedLinks
}
