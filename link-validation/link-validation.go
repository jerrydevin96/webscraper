package linkValidation

import (
	"log"
	"net/url"
)

type ValidatedLinks struct {
	InternalLinks     []string
	ExternalLinks     []string
	InAccessibleLinks []string
}

func ClassifyLinks(links []string, URL string) ([]string, []string) {
	internalLinks := make([]string, 0)
	externalLinks := make([]string, 0)
	log.Println("URL is : " + URL)
	u, _ := url.Parse(URL)
	log.Println(u.Host)
	log.Println("printing links")
	for _, link := range links {
		log.Println(link)
		v, _ := url.Parse(link)
		log.Println(v.Host)
	}
	return internalLinks, externalLinks
}
