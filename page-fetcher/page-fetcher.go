package pageFetcher

import (
	"io/ioutil"
	"log"
	"net/http"
)

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
