package pageFetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

type PageAccessibilityStatus struct {
	URL        string
	StatusCode int
	Error      error
}

func FetchHTMLPage(URL string) (string, error) {
	pageData := ""
	var err error
	response, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer response.Body.Close()

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading HTTP body. ", err)
		return "", err
	}
	pageData = string(body)
	return pageData, err
}

func CheckPageAccessibility(URL string, c chan PageAccessibilityStatus) {
	statusCode := 0
	var pageStatus PageAccessibilityStatus
	var err error
	pageStatus.URL = URL
	response, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		pageStatus.StatusCode = 0
		pageStatus.Error = err
		c <- pageStatus
		return
	}
	statusCode = response.StatusCode
	pageStatus.StatusCode = statusCode
	pageStatus.Error = err
	c <- pageStatus
}
