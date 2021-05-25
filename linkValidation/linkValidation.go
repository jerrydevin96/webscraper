package linkValidation

import (
	"log"
	"net/url"
	"strings"

	"github.com/jerrydevin96/webscraper/pageFetcher"
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
			internalLinks = append(internalLinks, link)
		} else if strings.Contains(linkHost, baseHost) {
			internalLinks = append(internalLinks, link)
		} else {
			externalLinks = append(externalLinks, link)
		}
	}
	return internalLinks, externalLinks
}

func checkLinkAccessibility(links []string) []string {
	nonAccessibleLinks := make([]string, 0)
	c := make(chan pageFetcher.PageAccessibilityStatus)
	for _, link := range links {
		go pageFetcher.CheckPageAccessibility(link, c)
	}

	for i := 0; i < len(links); i++ {
		accessibilityStatus := <-c
		url := accessibilityStatus.URL
		statusCode := accessibilityStatus.StatusCode
		err := accessibilityStatus.Error
		if (statusCode != 200 && statusCode != 201) || err != nil {
			log.Println("[WARN] Link not accessible")
			log.Println(accessibilityStatus.URL)
			log.Println(accessibilityStatus.StatusCode)
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
